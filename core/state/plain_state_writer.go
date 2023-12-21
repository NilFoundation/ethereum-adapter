package state

import (
	"encoding/binary"
	"fmt"
	"time"

	replication_adapter "github.com/NilFoundation/replication-adapter"
	"github.com/NilFoundation/replication-adapter-lib/core"
	"github.com/holiman/uint256"
	libcommon "github.com/ledgerwatch/erigon-lib/common"
	"github.com/ledgerwatch/erigon-lib/kv"
	"github.com/ledgerwatch/erigon-lib/kv/dbutils"
	"github.com/ledgerwatch/erigon/core/types/accounts"
	"github.com/ledgerwatch/erigon/turbo/shards"
)

var _ WriterWithChangeSets = (*PlainStateWriter)(nil)

type putDel interface {
	kv.Putter
	kv.Deleter
	IncrementSequence(bucket string, amount uint64) (uint64, error)
}
type PlainStateWriter struct {
	db          putDel
	csw         *ChangeSetWriter
	accumulator *shards.Accumulator
}

func NewPlainStateWriter(db putDel, changeSetsDB kv.RwTx, blockNumber uint64) *PlainStateWriter {
	return &PlainStateWriter{
		db:  db,
		csw: NewChangeSetWriterPlain(changeSetsDB, blockNumber),
	}
}

func NewPlainStateWriterNoHistory(db putDel) *PlainStateWriter {
	return &PlainStateWriter{
		db: db,
	}
}

func (w *PlainStateWriter) SetAccumulator(accumulator *shards.Accumulator) *PlainStateWriter {
	w.accumulator = accumulator
	return w
}

func (w *PlainStateWriter) UpdateAccountData(address libcommon.Address, original, account *accounts.Account, adapter replication_adapter.Adapter) error {
	if w.csw != nil {
		if err := w.csw.UpdateAccountData(address, original, account, adapter); err != nil {
			return err
		}
	}
	value := make([]byte, account.EncodingLengthForStorage())
	account.EncodeForStorage(value)
	if w.accumulator != nil {
		w.accumulator.ChangeAccount(address, account.Incarnation, value)
	}
	if adapter.IsWritable() {
		op := core.BasicOperation[any]{
			Type:        core.CreateBasicOp,
			BlockNumber: core.BlockNumber(w.csw.blockNumber),
			Params: core.CreateParams{
				AccountType: core.ExternallyOwnedAccount,
				Address:     core.Address(address.String()),
				CodeHash:    core.Hash(account.CodeHash),
				StorageRoot: core.Hash(account.Root),
			},
		}
		resp, err := adapter.SendAccountOperation(op)
		problemConnection := false
		for err != nil {
			problemConnection = true
			fmt.Println("Sleep while write to State Keeper")
			time.Sleep(time.Second * 1)
			resp, err = adapter.SendAccountOperation(op)
		}
		if problemConnection {
			fmt.Println("Connected!")
		}
		if resp.StatusCode == 500 {
			op = core.BasicOperation[any]{
				Type:        core.PutBasicOp,
				BlockNumber: core.BlockNumber(w.csw.blockNumber),
				Params: core.PutAccountParams{
					Address:     core.Address(address.String()),
					Balance:     &account.Balance,
					StorageRoot: core.Hash(account.Root),
				},
			}
			resp, err = adapter.SendAccountOperation(op)
			problemConnection = false
			for err != nil {
				problemConnection = true
				fmt.Println("Sleep while write to State Keeper")
				time.Sleep(time.Second * 1)
				resp, err = adapter.SendAccountOperation(op)
			}
			if problemConnection {
				fmt.Println("Connected!")
			}
			if resp.StatusCode == 500 {
				panic(resp.StatusCode)
			}
		}
	}
	return w.db.Put(kv.PlainState, address[:], value)
}

func (w *PlainStateWriter) UpdateAccountCode(address libcommon.Address, incarnation uint64, codeHash libcommon.Hash, code []byte) error {
	//fmt.Printf("code,%x,%x\n", address, code)
	if w.csw != nil {
		if err := w.csw.UpdateAccountCode(address, incarnation, codeHash, code); err != nil {
			return err
		}
	}
	if w.accumulator != nil {
		w.accumulator.ChangeCode(address, incarnation, code)
	}
	if err := w.db.Put(kv.Code, codeHash[:], code); err != nil {
		return err
	}
	return w.db.Put(kv.PlainContractCode, dbutils.PlainGenerateStoragePrefix(address[:], incarnation), codeHash[:])
}

func (w *PlainStateWriter) DeleteAccount(address libcommon.Address, original *accounts.Account) error {
	//fmt.Printf("delete,%x\n", address)
	if w.csw != nil {
		if err := w.csw.DeleteAccount(address, original); err != nil {
			return err
		}
	}
	if w.accumulator != nil {
		w.accumulator.DeleteAccount(address)
	}
	if err := w.db.Delete(kv.PlainState, address[:]); err != nil {
		return err
	}
	if original.Incarnation > 0 {
		var b [8]byte
		binary.BigEndian.PutUint64(b[:], original.Incarnation)
		if err := w.db.Put(kv.IncarnationMap, address[:], b[:]); err != nil {
			return err
		}
	}
	return nil
}

func (w *PlainStateWriter) WriteAccountStorage(address libcommon.Address, incarnation uint64, key *libcommon.Hash, original, value *uint256.Int, adapter replication_adapter.Adapter) error {
	//fmt.Printf("storage,%x,%x,%x\n", address, *key, value.Bytes())
	if adapter.IsWritable() {
		op := core.BasicOperation[any]{
			Type:        core.PutBasicOp,
			BlockNumber: core.BlockNumber(w.csw.blockNumber),
			Params: core.PutStorageParams{
				Address: core.Address(address.String()),
				Chunks: []core.StorageChunk{
					{core.Address(key.String()), value.Bytes32()},
				},
			},
		}

		resp, err := adapter.SendStorageOperation(op)
		problemConnection := false
		for err != nil {
			problemConnection = true
			fmt.Printf("Sleep while write to State Keeper")
			time.Sleep(time.Second * 1)
			resp, err = adapter.SendStorageOperation(op)
		}
		if problemConnection {
			fmt.Println("Connected!")
		}
		if resp.StatusCode == 500 {
			panic(resp.StatusCode)
		}
		_ = resp
	}

	if w.csw != nil {
		if err := w.csw.WriteAccountStorage(address, incarnation, key, original, value, adapter); err != nil {
			return err
		}
	}
	if *original == *value {
		return nil
	}
	compositeKey := dbutils.PlainGenerateCompositeStorageKey(address.Bytes(), incarnation, key.Bytes())

	v := value.Bytes()
	if w.accumulator != nil {
		w.accumulator.ChangeStorage(address, incarnation, *key, v)
	}
	if len(v) == 0 {
		return w.db.Delete(kv.PlainState, compositeKey)
	}
	return w.db.Put(kv.PlainState, compositeKey, v)
}

func (w *PlainStateWriter) CreateContract(address libcommon.Address) error {
	if w.csw != nil {
		if err := w.csw.CreateContract(address); err != nil {
			return err
		}
	}
	return nil
}

func (w *PlainStateWriter) WriteChangeSets() error {
	if w.csw != nil {

		return w.csw.WriteChangeSets()
	}

	return nil
}

func (w *PlainStateWriter) WriteHistory() error {
	if w.csw != nil {
		return w.csw.WriteHistory()
	}

	return nil
}

func (w *PlainStateWriter) ChangeSetWriter() *ChangeSetWriter {
	return w.csw
}
