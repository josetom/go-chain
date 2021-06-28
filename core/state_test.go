package core

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/josetom/go-chain/common"
	"github.com/josetom/go-chain/fs"
	"github.com/josetom/go-chain/test_helper"
)

func TestLoadStateValid(t *testing.T) {
	fs.Config.DataDir = test_helper.GetTestDataDir()
	state, err := LoadState()
	if err != nil {
		t.Fail()
	}
	if state.Balances[common.NewAddress(test_helper.Test_Address_2)] != 200 {
		t.Fail()
	}
}

func TestAddTransactionSuccess(t *testing.T) {
	fs.Config.DataDir = test_helper.GetTestDataDir()
	state, err := LoadState()
	if err != nil {
		t.Fail()
	}
	txn := getTestTxn()
	err = state.AddTransaction(txn)
	if err != nil {
		t.Fail()
	}
}

func TestAddTransactionInsufficientBalance(t *testing.T) {
	state := &State{
		Balances: make(map[common.Address]uint),
		dbFile:   nil,
	}
	txn := getTestTxn()
	err := state.AddTransaction(txn)
	if err == nil || err.Error() != "insufficient_balance" {
		t.Fail()
	}
}

func TestAddBlock(t *testing.T) {
	f, _ := os.CreateTemp("", "persist.db") // Temp gives much better performance
	// f, _ := os.Create(test_helper.GetTestFile("database/persist.db")) // Use this to debug if there are any failures
	state := &State{
		Balances: make(map[common.Address]uint),
		dbFile:   f,
	}
	state.Balances[common.NewAddress(test_helper.Test_Address_1)] = 100
	txn := getTestTxn()

	block := NewBlock(
		state.latestBlockHash,
		state.NextBlockNumber(),
		uint64(time.Now().UnixNano()),
		0,
		common.NewAddress(""), // miner.Config.Address
		[]Transaction{txn},
	)

	validBlock, err := mineBlockHelper(block)
	if err != nil {
		print(err)
		t.Fail()
	}

	blockHash, err := state.AddBlock(validBlock)

	if err != nil {
		print(err)
		t.Fail()
	}

	content, _ := ioutil.ReadFile(f.Name())

	var blockFS *BlockFS
	err = json.Unmarshal(content, &blockFS)
	if err != nil {
		t.Fail()
	}

	readBlockHash, err := blockFS.Block.Hash()
	if err != nil {
		t.Fail()
	}

	if blockHash.String() != readBlockHash.String() {
		t.Fail()
	}

	if blockFS.Block.Transactions[0] != txn {
		t.Fail()
	}
}

func TestNextBlockNumber(t *testing.T) {
	state := &State{
		Balances: make(map[common.Address]uint),
		dbFile:   nil,
	}
	if state.NextBlockNumber() != 1 {
		t.Fail()
	}
}
