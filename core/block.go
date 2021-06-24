package core

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/josetom/go-chain/common"
)

type Block struct {
	Header       BlockHeader   `json:"header"`
	Transactions []Transaction `json:"transactions"`
}

type BlockHeader struct {
	ParentHash common.Hash    `json:"parentHash"`
	Timestamp  uint64         `json:"timestamp"`
	Number     uint64         `json:"number"`
	Nonce      uint64         `json:"nonce"`
	Miner      common.Address `json:"miner"`
}

type BlockFS struct {
	Hash  common.Hash `json:"hash"`
	Block Block       `json:"block"`
}

func NewBlock(parentHash common.Hash, number uint64, time uint64, nonce uint64, miner common.Address, transactions []Transaction) Block {
	return Block{
		BlockHeader{
			ParentHash: parentHash,
			Timestamp:  time,
			Number:     number,
			Nonce:      nonce,
			Miner:      miner,
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

func (b Block) IsEmpty() bool {
	return len(b.Transactions) == 0
}

func IsBlockHashValid(hash common.Hash) bool {
	fmt_s := "%0" + fmt.Sprint(Config.Block.Complexity) + "d"
	s := fmt.Sprintf(fmt_s, 0) // if complexity = 5; generates "00000"
	return strings.HasSuffix(hash.String(), s)
}
