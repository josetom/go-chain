package test_helper_core

import (
	"time"

	"github.com/josetom/go-chain/common"
	"github.com/josetom/go-chain/core"
	"github.com/josetom/go-chain/db"
	"github.com/josetom/go-chain/fs"
	"github.com/josetom/go-chain/test_helper"
)

// Keep this in sync with getTestTxn in testhelper_test.go in core
func GetTestTxn() core.Transaction {

	data := test_helper.DUMMY_DATA

	txn := core.NewTransaction(
		common.NewAddress(test_helper.Test_Address_1),
		common.NewAddress(test_helper.Test_Address_2),
		100,
		data,
	)
	txn.TxnContent.Timestamp = uint64(time.Time{}.UnixNano())
	txn.Hash()
	txn.Sign()
	return txn
}

var state *core.State

func GetTestState() (*core.State, error) {
	if state == nil {
		db.Config.Type = db.LEVEL_DB
		fs.Config.DataDir = test_helper.GetTestDataDir()
		s, err := core.LoadState()
		if err != nil {
			return nil, err
		}
		state = s
	}
	return state, nil
}

func GetTestBlock(isValid bool, state *core.State, txns []core.Transaction) (core.Block, error) {
	block := core.NewBlock(
		state.LatestBlockHash(),
		state.NextBlockNumber(),
		uint64(time.Time{}.UnixNano()),
		0,
		common.NewAddress(""), // miner.Config.Address
		core.MINING_ALGO_POW,
		uint64(core.Config.Block.Reward),
		txns,
	)

	if isValid {
		return mineBlockHelper(block)
	} else {
		return block, nil
	}
}

func mineBlockHelper(pendingBlock core.Block) (core.Block, error) {
	isBlockValid, err := pendingBlock.IsBlockHashValid()
	if err != nil {
		return core.Block{}, err
	}
	if isBlockValid {
		return pendingBlock, nil
	}
	pendingBlock.Header.Nonce = common.GenNonce()
	return mineBlockHelper(pendingBlock)
}
