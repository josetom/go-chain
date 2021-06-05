package core

import (
	"encoding/json"

	"github.com/josetom/go-chain/common"
)

type Block struct {
	Header       BlockHeader   `json:"header"`
	Transactions []Transaction `json:"transactions"`
}

type BlockHeader struct {
	ParentHash common.Hash `json:"parentHash"`
	Timestamp  uint64      `json:"timestamp"`
	Number     uint64      `json:"number"`
}

type BlockFS struct {
	Hash  common.Hash `json:"hash"`
	Block Block       `json:"block"`
}

func NewBlock(parentHash common.Hash, number uint64, time uint64, transactions []Transaction) Block {
	return Block{
		BlockHeader{
			ParentHash: parentHash,
			Timestamp:  time,
			Number:     number,
		},
		transactions,
	}
}

func (b Block) Hash() (common.Hash, error) {
	blockJson, err := json.Marshal(b)
	if err != nil {
		return common.Hash{}, err
	}
	return common.BytesToHash(blockJson), nil
}
