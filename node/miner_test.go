package node

import (
	"context"
	"testing"

	"github.com/josetom/go-chain/core"
	"github.com/josetom/go-chain/fs"
	"github.com/josetom/go-chain/test_helper"
	"github.com/josetom/go-chain/test_helper/test_helper_core"
)

func TestMine(t *testing.T) {
	fs.Config.DataDir = test_helper.GetTestDataDir()
	core.Config.State.DbFile = test_helper.CreateAndGetTestDbFile()

	state, err := core.LoadState()

	if err != nil {
		t.Fail()
	}

	txn := test_helper_core.GetTestTxn()

	ctx := context.Background()
	miner := InitMiner(state)
	miner.addPendingTxn(txn)

	block, err := miner.mine(ctx)
	if err != nil {
		t.Fail()
	}

	isBlockValid, err := block.IsBlockHashValid()
	if err != nil {
		t.Error(err)
	}
	if !isBlockValid {
		t.Fail()
	}

	cleanup := func() {
		test_helper.DeleteTestDbFile()
		core.Config.State.DbFile = core.Defaults().State.DbFile
	}
	t.Cleanup(cleanup)

}
