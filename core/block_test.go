package core

import (
	"log"
	"testing"
	"time"

	"github.com/josetom/go-chain/common"
)

func TestBlockHash(t *testing.T) {
	tx := NewTransaction(
		NewAddress("0x0000000000000000000000000000000000000000"),
		NewAddress("0x3030303030303030303030303030303030313030"),
		100,
		"reward",
	)
	tx.Timestamp = uint64(time.Time{}.UnixNano())
	block := NewBlock(common.BytesToHash(nil), uint64(time.Time{}.UnixNano()), []Transaction{tx})
	blockHash, err := block.Hash()
	if blockHash.String() != "0xd64e3ead57afb35ab068167fe8849e4828c4e62bc727b9846e3ce004fb9eea84" || err != nil {
		log.Println(blockHash)
		t.Fail()
	}
}
