package core

import (
	"testing"

	"github.com/josetom/go-chain/test_helper"
)

func TestTransactionHash(t *testing.T) {
	txn := getTestTxn()
	if txn.Hash().String() != test_helper.Hash_Txn_100_Reward {
		t.Error(txn.Hash())
	}
}
