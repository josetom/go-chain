package core

import (
	"testing"

	"github.com/josetom/go-chain/test_helper"
)

func TestLoadGenesis(t *testing.T) {
	test_helper.SetTestDataDirs()
	genesisContent, err := loadGenesis()
	if err != nil || genesisContent == nil {
		t.Fail()
	}
}
