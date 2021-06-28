package test_helper_core

import (
	"time"

	"github.com/josetom/go-chain/common"
	"github.com/josetom/go-chain/core"
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
