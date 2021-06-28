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
	block := NewBlock(common.Hash{}, 0, uint64(time.Time{}.UnixNano()), 0, common.Address{}, []Transaction{txn})
	blockHash, err := block.Hash()
	if blockHash.String() != test_helper.Hash_Block_100_Reward || err != nil {
		log.Println(blockHash)
		t.Fail()
	}
}

func TestIsBlockHashValid(t *testing.T) {
	if !IsBlockHashValid(common.Hash{}) {
		t.Fail()
	}

	// Last Config.Block.Complexity bytes are number 0 => FAIL
	h1 := common.Hash{}
	h1.UnmarshalText([]byte("0xf572455bfe4edc8964b3197d07d1f27c6dc16cfaf250fbdc7eaa363030303033"))
	if IsBlockHashValid(h1) {
		log.Println(h1)
		t.Fail()
	}

	// last 5 chars of hex string are 0 => SUCCESS
	h2 := common.Hash{}
	h2.UnmarshalText([]byte("0xf572455bfe4edc8964b3197d07d1f27c6dc16cfaf250fbdc7eaa36abcde00000"))
	if !IsBlockHashValid(h2) {
		log.Println(h2)
		t.Fail()
	}
}
