package core

import (
	"github.com/josetom/go-chain/common"
	"github.com/josetom/go-chain/db/types"
)

// TODO : this needs to be fixed and changed to query from leveldb
func (s *State) GetBlocksAfter(hash common.Hash) ([]Block, error) {

	blocks := make([]Block, 0)

	var iter types.Iterator

	if hash.Equal(common.Hash{}) {
		iter = s.db.NewIterator([]byte(INDEX_BLOCK_NUMBER), getBlockNumberAsIndexBytes(0), nil)
	} else {

		blockFS, err := s.GetBlock(hash)
		if err != nil {
			return nil, err
		}

		iter = s.db.NewIterator([]byte(INDEX_BLOCK_NUMBER), getBlockNumberAsIndexBytes(blockFS.Block.Header.Number+1), nil)
	}
	for iter.Next() {
		blockFS, err := s.GetBlockWithHashBytes(iter.Value())
		if err != nil {
			return nil, err
		}
		blocks = append(blocks, blockFS.Block)
	}

	return blocks, nil
}
