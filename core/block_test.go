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
	block := NewBlock(common.BytesToHash(nil), 0, uint64(time.Time{}.UnixNano()), []Transaction{tx})
	blockHash, err := block.Hash()
	if blockHash.String() != "0xbbed322be92da72f1421e8d5a34cd24b9804e7c558516c0a19eb1494e15743db" || err != nil {
		log.Println(blockHash)
		t.Fail()
	}
}
