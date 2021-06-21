package core

import (
	"log"
	"testing"

	"github.com/josetom/go-chain/test_helper"
)

func TestIsRewardTrue(t *testing.T) {
	txn := getTestTxn(true)
	if !txn.TxnContent.IsReward {
		t.Fail()
	}
}

func TestIsRewardFalse(t *testing.T) {
	txn := getTestTxn(false)
	if txn.TxnContent.IsReward {
		t.Fail()
	}
}

func TestTransactionHash(t *testing.T) {
	txn := getTestTxn(true)
	if txn.Hash().String() != test_helper.Hash_Txn_100_Reward {
		log.Println(txn.Hash())
		t.Fail()
	}
}
