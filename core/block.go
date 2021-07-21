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
	ParentHash       common.Hash    `json:"parentHash"`
	Timestamp        uint64         `json:"timestamp"`
	Number           uint64         `json:"number"`
	Nonce            uint64         `json:"nonce"`
	Miner            common.Address `json:"miner"`
	MiningAlgo       string         `json:"miningAlgo"`
	MiningComplexity uint64         `json:"miningComplexity"`
	Reward           uint64         `json:"reward"`
}

type BlockFS struct {
	Hash  common.Hash `json:"hash"`
	Block Block       `json:"block"`
}

func NewBlock(
	parentHash common.Hash,
	number uint64,
	time uint64,
	nonce uint64,
	miner common.Address,
	miningAlgo string,
	reward uint64,
	transactions []Transaction,
) Block {
	return Block{
		BlockHeader{
			ParentHash:       parentHash,
			Timestamp:        time,
			Number:           number,
			Nonce:            nonce,
			Miner:            miner,
			MiningAlgo:       miningAlgo,
			MiningComplexity: Config.Block.Complexity,
			Reward:           reward,
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

func (b Block) IsBlockHashValid() (bool, error) {
	hash, err := b.Hash()
	if err != nil {
		return false, err
	}
	fmt_s := "%0" + fmt.Sprint(b.Header.MiningComplexity) + "d"
	s := fmt.Sprintf(fmt_s, 0) // if complexity = 5; generates "00000"
	return strings.HasPrefix(hash.String()[2:], s), nil
}
