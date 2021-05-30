package core

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/josetom/go-chain/node"
)

func setDataDirWithLocalTestPath() {
	node.Config = &node.Defaults
	Config = &Defaults
	cwd, _ := os.Getwd()
	node.Config.DataDir = filepath.Join(cwd, "testdata")
}

// func TestLoadStateValid(t *testing.T) {
// 	setDataDirWithLocalTestPath()
// 	state, err := LoadState()
// 	if err != nil {
// 		t.Fail()
// 	}
// 	if state.Balances[NewAddress("0x3030303030303030303030303030303030313030")] != 600 {
// 		t.Fail()
// 	}
// 	if state.Balances[NewAddress("0x3030303030303030303030303030303030323030")] != 100 {
// 		t.Fail()
// 	}
// }

func TestLoadStateMissingFile(t *testing.T) {
	setDataDirWithLocalTestPath()
	Config.State.DbFile = "database/missing.db"
	_, err := LoadState()
	if err == nil {
		t.Fail()
	}
}

func TestAddSuccess(t *testing.T) {
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
	state.Add(txn)
	if state.txMemPool[0] != txn {
		t.Fail()
	}
}

func TestAddInsufficientBalance(t *testing.T) {
	state := &State{
		txMemPool: make([]Transaction, 0),
		Balances:  make(map[Address]uint),
		dbFile:    nil,
	}
	txn := NewTransaction(
		NewAddress("0x0000000000000000000000000000000000000000"),
		NewAddress("0x3030303030303030303030303030303030313030"),
		100,
		"random data",
	)
	err := state.Add(txn)
	if err == nil || err.Error() != "insufficient_balance" {
		t.Fail()
	}
}

func TestPersistSuccess(t *testing.T) {
	f, _ := os.CreateTemp("", "persist.db") // Temp gives much better performance
	// f, _ := os.Create("testdata/database/persist.db") // Use this to debug if there are any failures
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
	txn.Timestamp = time.Time{}
	state.Add(txn)
	err := state.Persist()
	if err != nil {
		print(err)
		t.Fail()
	}

	content, _ := ioutil.ReadFile(f.Name())
	var t2 *Transaction
	err = json.Unmarshal(content, &t2)

	if err != nil {
		t.Fail()
	}

	if *t2 != txn {
		t.Fail()
	}

	if len(state.txMemPool) > 0 {
		t.Fail()
	}
}
