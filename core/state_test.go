package core

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/josetom/go-chain/fs"
	"github.com/josetom/go-chain/test_helper"
)

func TestLoadStateValid(t *testing.T) {
	fs.Config.DataDir = test_helper.GetTestDataDir()
	state, err := LoadState()
	if err != nil {
		t.Fail()
	}
	if state.Balances[NewAddress("0x3030303030303030303030303030303030313030")] != 200 {
		t.Fail()
	}
}

func TestAddTransactionRewardSuccess(t *testing.T) {
	state := &State{
		txMemPool: make([]Transaction, 0),
		Balances:  make(map[Address]uint),
		dbFile:    nil,
	}
	txn := NewTransaction(
		NewAddress("0x0000000000000000000000000000000000000000"),
		NewAddress("0x3030303030303030303030303030303030313030"),
		100,
		"reward",
	)
	err := state.AddTransaction(txn)
	if err != nil {
		t.Fail()
	}
	if state.txMemPool[0] != txn {
		t.Fail()
	}
}

func TestAddTransactionNonRewardSuccess(t *testing.T) {
	fs.Config.DataDir = test_helper.GetTestDataDir()
	state, err := LoadState()
	if err != nil {
		t.Fail()
	}
	txn := NewTransaction(
		NewAddress("0x3030303030303030303030303030303030313030"),
		NewAddress("0x3030303030303030303030303030303030323030"),
		100,
		"something else",
	)
	err = state.AddTransaction(txn)
	if err != nil {
		t.Fail()
	}
	if state.txMemPool[0] != txn {
		t.Fail()
	}
}

func TestAddTransactionInsufficientBalance(t *testing.T) {
	state := &State{
		txMemPool: make([]Transaction, 0),
		Balances:  make(map[Address]uint),
		dbFile:    nil,
	}
	txn := NewTransaction(
		NewAddress("0x0000000000000000000000000000000000000000"),
		NewAddress("0x3030303030303030303030303030303030313030"),
		100,
		"something else",
	)
	err := state.AddTransaction(txn)
	if err == nil || err.Error() != "insufficient_balance" {
		t.Fail()
	}
}

func TestPersistSuccess(t *testing.T) {
	f, _ := os.CreateTemp("", "persist.db") // Temp gives much better performance
	// f, _ := os.Create(test_helper.GetTestFile("database/persist.db")) // Use this to debug if there are any failures
	state := &State{
		txMemPool: make([]Transaction, 0),
		Balances:  make(map[Address]uint),
		dbFile:    f,
	}
	txn := NewTransaction(
		NewAddress("0x0000000000000000000000000000000000000000"),
		NewAddress("0x3030303030303030303030303030303030313030"),
		100,
		"reward",
	)
	txn.Timestamp = uint64(time.Time{}.UnixNano())
	state.AddTransaction(txn)
	blockHash, err := state.Persist()
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

	if len(state.txMemPool) > 0 {
		t.Fail()
	}
}

func TestNextBlockNumber(t *testing.T) {
	state := &State{
		txMemPool: make([]Transaction, 0),
		Balances:  make(map[Address]uint),
		dbFile:    nil,
	}
	if state.NextBlockNumber() != 1 {
		t.Fail()
	}
}
