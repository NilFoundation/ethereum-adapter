package state

import (
	"github.com/ledgerwatch/erigon/core/types/accounts"
)

type accountCreate struct {
	accounts.Account
	address string
}

type accountUpdate struct {
	accounts.Account
	address string
}

/////////////////////////////////////////////////////////////////
//
//func (op accountCreate) MakeBasicChain() []core.BasicOperation[any] {
//	return []core.BasicOperation[any]{
//		{
//			Type: core.InsertBasicOp,
//			Params: core.InsertParams{
//				AccountType: core.ExternallyOwnedAccount,
//				Address:     core.Address(op.address),
//				// For key will be also some json
//				Amount: big.NewInt(0),
//				// I wish to make smth like any json here
//			},
//		},
//	}
//}
//
//func (op accountUpdate) MakeBasicChain() []core.BasicOperation[any] {
//	return []core.BasicOperation[any]{
//		{
//			Type: core.PutBasicOp,
//			Params: core.InsertParams{
//				AccountType: core.ExternallyOwnedAccount,
//				Address:     core.Address(op.address),
//				Amount:      op.Balance.ToBig(),
//			},
//		},
//	}
//}

/////////////////////////////////////////////////////////////////

//func NewaccountCreate(account accounts.Account, address string) replication_adapter.Operation {
//	return accountCreate{account, address}
//}
//
//func NewaccountUpdate(account accounts.Account, address string) replication_adapter.Operation {
//	return accountUpdate{account, address}
//}
