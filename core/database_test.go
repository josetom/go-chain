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
		if err != nil || hash.String() != test_helper.Hash_Block_0 {
			t.Fail()
		}
	}

	hash1 := common.Hash{}
	hash1.UnmarshalText([]byte(test_helper.Hash_Block_0))
	blocksFrom1, err := s.GetBlocksAfter(hash1)
	if len(blocksFrom1) != 1 || err != nil {
		hash, err := blocksFrom1[0].Hash()
		if err != nil || hash.String() != test_helper.Hash_Block_1 {
			t.Fail()
		}
	}
}
