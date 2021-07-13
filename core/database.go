package core

import (
	"encoding/json"

	"github.com/josetom/go-chain/common"
	"github.com/josetom/go-chain/db/types"
)

// TODO : this needs to be fixed and changed to query from leveldb
func (s *State) GetBlocksAfter(hash common.Hash) ([]Block, error) {

	blocks := make([]Block, 0)

	var iter types.Iterator
	var blockFS BlockFS
	if hash.Equal(common.Hash{}) {
		iter = s.db.NewIterator(nil, nil)
	} else {
		iter = s.db.NewIterator(hash.Bytes(), nil)
		iter.Next() // skip the current one and forward
	}
	for iter.Next() {
		json.Unmarshal(iter.Value(), &blockFS)
		blocks = append(blocks, blockFS.Block)
	}

	return blocks, nil
}
