package core_test

import (
	"testing"

	"github.com/josetom/go-chain/common"
	"github.com/josetom/go-chain/core"
	"github.com/josetom/go-chain/db"
	"github.com/josetom/go-chain/test_helper"
	"github.com/josetom/go-chain/test_helper/test_helper_core"
)

func TestLoadStateValid(t *testing.T) {
	db.Config.Type = db.LEVEL_DB
	test_helper.SetTestDataDirs()
	state, err := test_helper_core.GetTestState()
	if err != nil {
		t.Error(err)
	}
	if state.Balances[common.NewAddress(test_helper.Test_Address_2)] != 200 {
		t.Fail()
	}
	cleanup := func() {
		state.Close()
	}
	t.Cleanup(cleanup)
}

func TestAddTransactionSuccess(t *testing.T) {
	db.Config.Type = db.LEVEL_DB
	test_helper.SetTestDataDirs()
	state, err := test_helper_core.GetTestState()
	if err != nil {
		t.Error(err)
	}
	txn := test_helper_core.GetTestTxn()
	err = state.AddTransaction(txn)
	if err != nil {
		t.Error(err)
	}
	cleanup := func() {
		state.Close()
	}
	t.Cleanup(cleanup)
}

func TestAddTransactionInsufficientBalance(t *testing.T) {
	state := &core.State{
		Balances: make(map[common.Address]uint),
	}
	txn := test_helper_core.GetTestTxn()
	err := state.AddTransaction(txn)
	if err == nil || err.Error() != "insufficient_balance" {
		t.Fail()
	}
}

func TestAddBlock(t *testing.T) {
	db.Config.Type = db.LEVEL_DB
	test_helper.SetTestDataDirs()
	tempDbPath := test_helper.CreateAndGetTestDbFile()
	core.Config.State.DbFile = tempDbPath

	state, err := core.LoadState()
	if err != nil {
		t.Error(err)
	}
	state.Balances[common.NewAddress(test_helper.Test_Address_1)] = 100
	txn := test_helper_core.GetTestTxn()

	validBlock, err := test_helper_core.GetTestBlock(true, state, []core.Transaction{txn})
	if err != nil {
		t.Error(err)
	}

	blockHash, err := state.AddBlock(validBlock)

	if err != nil {
		t.Error(err)
	}

	blockFS, err := state.GetBlock(blockHash)
	if err != nil {
		t.Error(err)
	}

	readBlockHash, err := blockFS.Block.Hash()
	if err != nil {
		t.Error(err)
	}

	if blockHash.String() != readBlockHash.String() {
		t.Error(blockHash)
	}

	if blockFS.Block.Transactions[0] != txn {
		t.Fail()
	}

	cleanup := func() {
		state.Close()
		test_helper.DeleteTestDbFile(tempDbPath)
	}
	t.Cleanup(cleanup)
}

func TestNextBlockNumber(t *testing.T) {
	state := &core.State{
		Balances: make(map[common.Address]uint),
	}
	if state.NextBlockNumber() != 1 {
		t.Fail()
	}
}
