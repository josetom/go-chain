package core

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/josetom/go-chain/node"
)

type GenesisConfig struct {
	ChainId string `json:"chainId"`
}

type Genesis struct {
	Timestamp uint64          `json:"timestamp"`
	Config    GenesisConfig   `json:"config"`
	Balances  map[string]uint `json:"balances"`
}

var genesisContent *Genesis

func loadGenesis() (*Genesis, error) {
	genesisFilePath := filepath.Join(node.Config.DataDir, "database/genesis.json")
	content, err := ioutil.ReadFile(genesisFilePath)
	if err != nil {
		log.Println("Load file failed", genesisFilePath)
		return nil, err
	}

	err = json.Unmarshal(content, &genesisContent)
	if err != nil {
		log.Println("Unmarshall failed", genesisFilePath)
		return nil, err
	}
	log.Println("Loaded genesis")

	return genesisContent, nil
}
