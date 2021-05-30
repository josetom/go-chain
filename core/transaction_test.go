package core

import (
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
	if tx.TxnHash.String() != "0x2c46899f8df91ce9dcba95cdd63551524f5eba1e825399120a161be19605fe8c" {
		t.Fail()
	}
}
