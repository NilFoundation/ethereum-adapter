package generic

import (
	"github.com/ledgerwatch/erigon/cmd/devnet/tests"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	runCtx, _ := tests.ContextStart(t, "")
	for {
		time.Sleep(1 * time.Second)
	}
	_ = runCtx
}
