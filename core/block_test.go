package core

import (
	"log"
	"testing"
	"time"

	"github.com/josetom/go-chain/common"
	"github.com/josetom/go-chain/test_helper"
)

func TestBlockHash(t *testing.T) {
	txn := getTestTxn()
	block := NewBlock(
		common.Hash{},
		0,
		uint64(time.Time{}.UnixNano()),
		0,
		common.Address{},
		MINING_ALGO_POW,
		uint64(Config.Block.Reward),
		[]Transaction{txn},
	)
	blockHash, err := block.Hash()
	if blockHash.String() != test_helper.Hash_Block_100_Reward || err != nil {
		log.Println(blockHash)
		t.Fail()
	}
}

func TestIsBlockHashValid(t *testing.T) {

	// invalid block
	block1, err := getTestBlock(false, &State{}, []Transaction{})
	if err != nil {
		t.Error(err)
	}
	isBlockValid1, err := block1.IsBlockHashValid()
	if err != nil {
		t.Error(err)
	}
	if isBlockValid1 {
		t.Fail()
	}

	// valid block
	block2, err := getTestBlock(true, &State{}, []Transaction{})
	if err != nil {
		t.Error(err)
	}
	isBlockValid2, err := block2.IsBlockHashValid()
	if err != nil {
		t.Error(err)
	}
	if !isBlockValid2 {
		t.Fail()
	}
}
