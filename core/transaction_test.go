package core

import (
	"log"
	"testing"
)

func TestIsRewardTrue(t *testing.T) {
	tx := NewTransaction(
		NewAddress("0x0000000000000000000000000000000000000000"),
		NewAddress("0x3030303030303030303030303030303030313030"),
		100,
		"reward",
	)
	isReward := tx.IsReward()
	if !isReward {
		t.Fail()
	}
}

func TestIsRewardFalse(t *testing.T) {
	tx := NewTransaction(
		NewAddress("0x0000000000000000000000000000000000000000"),
		NewAddress("0x3030303030303030303030303030303030313030"),
		100,
		"something else",
	)
	isReward := tx.IsReward()
	if isReward {
		t.Fail()
	}
}

func TestTransactionHash(t *testing.T) {
	tx := NewTransaction(
		NewAddress("0x0000000000000000000000000000000000000000"),
		NewAddress("0x3030303030303030303030303030303030313030"),
		100,
		"reward",
	)
	if tx.TxnHash.String() != "0x7f9a283e2cccb8396e2cd5af4dbf357217edaf0f9b6be50225b8637db00cd2a3" {
		log.Println(tx.TxnHash)
		t.Fail()
	}
}
