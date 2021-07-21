package core_test

import (
	"testing"

	"github.com/josetom/go-chain/common"
	"github.com/josetom/go-chain/test_helper"
	"github.com/josetom/go-chain/test_helper/test_helper_core"
)

func TestGetBlocksAfter(t *testing.T) {
	state, err := test_helper_core.GetTestState()

	if err != nil {
		t.Error(err)
	}

	blocksFrom0, err := state.GetBlocksAfter(common.Hash{})
	if len(blocksFrom0) != 2 || err != nil {
		t.Fail()
	} else {
		hash, err := blocksFrom0[0].Hash()
		if err != nil || hash.String() != test_helper.Hash_Block_0 {
			t.Error("hash", hash.String())
		}
	}

	hash1 := common.Hash{}
	hash1.UnmarshalText([]byte(test_helper.Hash_Block_0))
	blocksFrom1, err := state.GetBlocksAfter(hash1)
	if len(blocksFrom1) != 1 || err != nil {
		t.Fail()
	} else {
		hash, err := blocksFrom1[0].Hash()
		if err != nil || hash.String() != test_helper.Hash_Block_1 {
			t.Error("hash", hash.String())
		}
	}

	cleanup := func() {
		state.Close()
	}
	t.Cleanup(cleanup)
}
