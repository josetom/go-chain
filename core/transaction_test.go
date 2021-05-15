package core

import "testing"

func TestIsRewardTrue(t *testing.T) {
	tx := Transaction{"0x0", "0x100", 1000, "reward"}
	isReward := tx.IsReward()
	if !isReward {
		t.Fail()
	}
}

func TestIsRewardFalse(t *testing.T) {
	tx := Transaction{"0x0", "0x100", 1000, "something else"}
	isReward := tx.IsReward()
	if isReward {
		t.Fail()
	}
}
