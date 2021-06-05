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
	if blockHash.String() != "0xff312d625b214ae4eba8ed13055285b385d71ebbd1c5e2514c2a444b91de2c15" || err != nil {
		log.Println(blockHash)
		t.Fail()
	}
}
