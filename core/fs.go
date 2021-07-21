package core

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/josetom/go-chain/fs"
)

type FsProvider interface {
	InitFS() error
	GetDataDir() string
	GetGenesisFilePath() string
	GetBlocksDbPath() string
}

func InitFS() error {
	// Create Database directory if it doesn't exist
	err := initializeDbDirectory()
	if err != nil {
		return err
	}

	// Create Gensis file if it doesn't exist
	err = initializeGenesisFile()
	if err != nil {
		return err
	}

	// Initialize empty block db if it doesn't exist
	err = initializeBlockDb()
	if err != nil {
		return err
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

func initializeDbDirectory() error {
	if isExist, _ := fs.DoesExist(GetDataDir()); !isExist {
		if err := os.MkdirAll(GetDataDir(), os.ModePerm); err != nil && !os.IsExist(err) {
			return err
		}
	}
	return nil
}

func initializeGenesisFile() error {
	if isExist, _ := fs.DoesExist(GetGenesisFilePath()); !isExist {
		genesisContent, err := json.Marshal(defaultGenesis)
		if err != nil {
			return err
		}
		if err = ioutil.WriteFile(GetGenesisFilePath(), genesisContent, 0644); err != nil {
			return err
		}
	}
	return nil
}

func initializeBlockDb() error {
	// if isExist, _ := fs.DoesExist(GetBlocksDbPath()); !isExist {
	// err := fs.WriteEmptyBlocksDbToDisk(GetBlocksDbPath())
	// return err
	// }
	return nil
}
