package generic

import (
	"context"
	"fmt"
	"github.com/holiman/uint256"
	"github.com/ledgerwatch/erigon-lib/common"
	libcommon "github.com/ledgerwatch/erigon-lib/common"
	"github.com/ledgerwatch/erigon/accounts/abi/bind"
	"github.com/ledgerwatch/erigon/cmd/devnet/accounts"
	"github.com/ledgerwatch/erigon/cmd/devnet/blocks"
	"github.com/ledgerwatch/erigon/cmd/devnet/contracts"
	"github.com/ledgerwatch/erigon/cmd/devnet/devnet"
	"github.com/ledgerwatch/erigon/cmd/devnet/tests"
	"github.com/ledgerwatch/erigon/core/types"
	"github.com/ledgerwatch/erigon/rpc"
	"math/big"
	"testing"
	"time"
)

func waitNextBlock(node devnet.Node) uint64 {
	blockNumber1, _ := node.BlockNumber()
	blockNumber2, _ := node.BlockNumber()
	for blockNumber1 >= blockNumber2 {
		time.Sleep(time.Second * 1)
		blockNumber2, _ = node.BlockNumber()
	}
	return blockNumber2
}

func SendTx(ctx context.Context, to, from string, amount uint64) (libcommon.Hash, error) {
	// get the latest nonce for the next transaction
	const gasPrice uint64 = 912_345_678

	toAddress := libcommon.HexToAddress(to)
	fromAddress := libcommon.HexToAddress(from)

	gasFeeCap, err := blocks.BaseFeeFromBlock(ctx)

	node := devnet.SelectNode(ctx)

	res, err := node.GetTransactionCount(fromAddress, rpc.PendingBlock)

	nonce := res.Uint64()

	signer := *types.LatestSignerForChainID(node.ChainID())
	chainId := *uint256.NewInt(node.ChainID().Uint64())

	if err != nil {
		return libcommon.Hash{}, err
	}

	transaction := types.NewEIP1559Transaction(chainId, nonce, toAddress, uint256.NewInt(amount), uint64(210_000), uint256.NewInt(gasPrice), new(uint256.Int), uint256.NewInt(gasFeeCap), nil)

	signerKey := accounts.SigKey(fromAddress)
	if signerKey == nil {
		return libcommon.Hash{}, fmt.Errorf("devnet.signEIP1559TxsHigherThanBaseFee failed to SignTx: private key not found for address %s", fromAddress)
	}

	signedTransaction, err := types.SignTx(transaction, signer, signerKey)
	if err != nil {
		return libcommon.Hash{}, err
	}
	nonce++
	hash, err := devnet.SelectNode(ctx).SendTransaction(signedTransaction)
	return hash, nil
}

func TestReplication(t *testing.T) {
	runCtx, _ := tests.ContextStart(t, "")
	ctx := runCtx.WithCurrentNetwork(0).WithCurrentNode(0)
	node := devnet.SelectBlockProducer(ctx)
	const recipientAddress = "0xa94f5374Fce5edBC8E2a8697C15331677e6EbF0B"

	const sendValue1 uint64 = 229
	SendTx(ctx, recipientAddress, accounts.DevAddress, sendValue1)

	fmt.Println("Block now:", waitNextBlock(node))

	const sendValue2 uint64 = 114
	SendTx(ctx, recipientAddress, accounts.DevAddress, sendValue2)

	fmt.Println("Block now:", waitNextBlock(node))

	bal, _ := node.GetBalance(common.HexToAddress(recipientAddress), rpc.BlockNumber(-1).AsBlockReference())
	fmt.Println("Balance:", bal)
	fmt.Println("-------------------------------------------------------------------------------------")
	transactor := common.HexToAddress(accounts.DevAddress)
	count, err := node.GetTransactionCount(transactor, rpc.LatestBlock)

	if err != nil {
		panic(err)
	}

	transactOpts, err := bind.NewKeyedTransactorWithChainID(accounts.SigKey(transactor), node.ChainID())

	if err != nil {
		panic(err)
	}

	transactOpts.GasLimit = uint64(200_000)
	transactOpts.GasPrice = big.NewInt(880_000_000)
	transactOpts.Nonce = count

	// deploy the contract and get the contract handler
	address, _, storeContract, err := contracts.DeployPairStore(transactOpts, contracts.NewBackend(node))
	fmt.Println("Contract address:", address)
	if err != nil {
		panic(err)
	}

	transactOpts.Nonce = big.NewInt(0).Add(transactOpts.Nonce, big.NewInt(1))
	_, err = storeContract.SetFirst(transactOpts, big.NewInt(23))

	fmt.Println("Block now:", waitNextBlock(node))

	value, err := storeContract.GetFirst(nil)
	fmt.Println("First value:", value)

	transactOpts.Nonce = big.NewInt(0).Add(transactOpts.Nonce, big.NewInt(1))
	_, err = storeContract.SetSecond(transactOpts, big.NewInt(55))

	fmt.Println("Block now:", waitNextBlock(node))

	value3, err := storeContract.GetSecond(nil)
	fmt.Println("Second value:", value3)
}
