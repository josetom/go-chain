package core

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/josetom/go-chain/node"
)

func InitFS() error {
	// Create Database directory if it doesn't exist
	if isExist, _ := doesExist(GetDataDir()); !isExist {
		if err := os.MkdirAll(GetDataDir(), os.ModePerm); err != nil && !os.IsExist(err) {
			return err
		}
	}

	// Create Gensis file if it doesn't exist
	if isExist, _ := doesExist(GetGenesisFilePath()); !isExist {
		genesisContent, err := json.Marshal(defaultGenesis)
		if err != nil {
			return err
		}
		if err = ioutil.WriteFile(GetGenesisFilePath(), genesisContent, 0644); err != nil {
			return err
		}
	}

	// Initialize empty block db if it doesn't exist
	if isExist, _ := doesExist(GetBlocksDbPath()); !isExist {
		writeEmptyBlocksDbToDisk(GetBlocksDbPath())
	}

	return nil
}

func GetDataDir() string {
	return filepath.Join(node.Config.DataDir, "database")
}

func GetGenesisFilePath() string {
	return filepath.Join(GetDataDir(), "genesis.json")
}

func GetBlocksDbPath() string {
	return filepath.Join(GetDataDir(), Config.State.DbFile)
}

func doesExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false, err
	}
	return true, nil
}

func writeEmptyBlocksDbToDisk(path string) error {
	return ioutil.WriteFile(path, []byte(""), os.ModePerm)
}
