package node

import (
	"context"
	"testing"

	"github.com/josetom/go-chain/core"
	"github.com/josetom/go-chain/db"
	"github.com/josetom/go-chain/test_helper"
	"github.com/josetom/go-chain/test_helper/test_helper_core"
)

func TestMine(t *testing.T) {
	db.Config.Type = db.LEVEL_DB
	test_helper.SetTestDataDirs()
	tempDbPath := test_helper.CreateAndGetTestDbFile()
	core.Config.State.DbFile = tempDbPath

	state, err := core.LoadState()

	if err != nil {
		t.Error(err)
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
		state.Close()
		test_helper.DeleteTestDbFile(tempDbPath)
		core.Config.State.DbFile = core.Defaults().State.DbFile
	}
	t.Cleanup(cleanup)

}
