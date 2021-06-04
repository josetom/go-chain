package core

import (
	"testing"
)

func TestLoadGenesis(t *testing.T) {
	setDataDirWithLocalTestPath()
	genesisContent, err := loadGenesis()
	if err != nil || genesisContent == nil {
		t.Fail()
	}
}
