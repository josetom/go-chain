package core

import (
	"encoding/json"
	"io/ioutil"
	"log"

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

var genesisContent *Genesis

func loadGenesis() (*Genesis, error) {
	genesisFilePath := GetGenesisFilePath()
	content, err := ioutil.ReadFile(genesisFilePath)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(content, &genesisContent)
	if err != nil {
		return nil, err
	}
	log.Println("Genesis file loaded")

	return genesisContent, nil
}

func (g Genesis) Hash() (common.Hash, error) {
	blockJson, err := json.Marshal(g)
	if err != nil {
		return common.Hash{}, err
	}
	return common.BytesToHash(blockJson), nil
}
