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
	block := NewBlock(common.Hash{}, 0, uint64(time.Time{}.UnixNano()), []Transaction{tx})
	blockHash, err := block.Hash()
	if blockHash.String() != "0x8fc3e448220596d4caec6cf6b0767f27a6ff580ef4ae9d80798f14950f39e2a9" || err != nil {
		log.Println(blockHash)
		t.Fail()
	}
}
