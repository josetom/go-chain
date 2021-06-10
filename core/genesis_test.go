package core

import (
	"testing"

	"github.com/josetom/go-chain/fs"
	"github.com/josetom/go-chain/test_helper"
)

func TestLoadGenesis(t *testing.T) {
	fs.Config.DataDir = test_helper.GetTestDataDir()
	genesisContent, err := loadGenesis()
	if err != nil || genesisContent == nil {
		t.Fail()
	}
}
