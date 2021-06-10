package core

import (
	"testing"

	"github.com/josetom/go-chain/common"
	"github.com/josetom/go-chain/fs"
	"github.com/josetom/go-chain/test_helper"
)

func TestGetBlocksAfter(t *testing.T) {
	fs.Config.DataDir = test_helper.GetTestDataDir()
	s := &State{}

	blocksFrom0, err := s.GetBlocksAfter(common.Hash{})
	if len(blocksFrom0) != 2 || err != nil {
		hash, err := blocksFrom0[0].Hash()
		if err != nil || hash.String() != "0xbfa63a77a70876ac1b5ebaba6d9113b181259aae5afa11207aeb5143a6ed9990" {
			t.Fail()
		}
	}

	hash1 := common.Hash{}
	hash1.UnmarshalText([]byte("0xbfa63a77a70876ac1b5ebaba6d9113b181259aae5afa11207aeb5143a6ed9990"))
	blocksFrom1, err := s.GetBlocksAfter(hash1)
	if len(blocksFrom1) != 1 || err != nil {
		hash, err := blocksFrom1[0].Hash()
		if err != nil || hash.String() != "0x39714f635bda97ef70bf48ecae1a8ea27a42cc5e35dd40895db35d44107bf1bd" {
			t.Fail()
		}
	}
}
