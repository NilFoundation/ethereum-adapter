package jsonrpc

import (
	"context"
	replication_adapter "github.com/NilFoundation/replication-adapter"

	"github.com/ledgerwatch/erigon-lib/common"
	"github.com/ledgerwatch/erigon-lib/common/hexutility"
)

func (api *OtterscanAPIImpl) GetTransactionError(ctx context.Context, hash common.Hash, adapter replication_adapter.Adapter) (hexutility.Bytes, error) {
	tx, err := api.db.BeginRo(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	result, err := api.runTracer(ctx, tx, hash, nil, adapter)
	if err != nil {
		return nil, err
	}

	return result.Revert(), nil
}
