package core

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/josetom/go-chain/fs"
)

func InitFS() error {
	// Create Database directory if it doesn't exist
	if isExist, _ := fs.DoesExist(GetDataDir()); !isExist {
		if err := os.MkdirAll(GetDataDir(), os.ModePerm); err != nil && !os.IsExist(err) {
			return err
		}
	}

	// Create Gensis file if it doesn't exist
	if isExist, _ := fs.DoesExist(GetGenesisFilePath()); !isExist {
		genesisContent, err := json.Marshal(defaultGenesis)
		if err != nil {
			return err
		}
		if err = ioutil.WriteFile(GetGenesisFilePath(), genesisContent, 0644); err != nil {
			return err
		}
	}

	// Initialize empty block db if it doesn't exist
	if isExist, _ := fs.DoesExist(GetBlocksDbPath()); !isExist {
		fs.WriteEmptyBlocksDbToDisk(GetBlocksDbPath())
	}

	return nil
}

func GetDataDir() string {
	return filepath.Join(fs.ExpandPath(fs.Config.DataDir), "database")
}

func GetGenesisFilePath() string {
	return filepath.Join(GetDataDir(), "genesis.json")
}

func GetBlocksDbPath() string {
	return filepath.Join(GetDataDir(), Config.State.DbFile)
}
