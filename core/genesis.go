package core

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/josetom/go-chain/common"
)

type GenesisConfig struct {
	ChainId string `json:"chainId"`
}

type Genesis struct {
	Timestamp uint64                  `json:"timestamp"`
	Config    GenesisConfig           `json:"config"`
	Balances  map[common.Address]uint `json:"balances"`
}

func InitGenesis() error {
	// Create Gensis file if it doesn't exist
	err := initializeGenesisFile()
	if err != nil {
		return err
	}

	genesisData, err := loadGenesisDataFromFile()
	if err != nil {
		return err
	}

	state, err := NewState()
	if err != nil {
		return err
	}

	err = state.applyGenesis(genesisData)
	if err != nil {
		return err
	}

	block := NewBlock(
		common.Hash{},
		state.NextBlockNumber(),
		uint64(time.Now().UnixNano()),
		0,
		common.Address{},
		"genesis",
		0,
		[]Transaction{},
		genesisData,
	)

	_, err = state.AddBlock(block)

	return err
}

func (g Genesis) Hash() (common.Hash, error) {
	blockJson, err := json.Marshal(g)
	if err != nil {
		return common.Hash{}, err
	}
	return common.BytesToHash(blockJson), nil
}

func loadGenesisDataFromFile() ([]byte, error) {
	genesisFilePath := GetGenesisFilePath()
	content, err := ioutil.ReadFile(genesisFilePath)
	if err != nil {
		return nil, err
	}

	return content, nil
}
