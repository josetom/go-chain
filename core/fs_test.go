package core

import (
	"testing"
)

func TestDoesExist(t *testing.T) {
	if isExist, _ := doesExist("testdata/database/genesis.json"); !isExist {
		t.Fail()
	}
	if isExist, _ := doesExist("testdata/missing.db"); isExist {
		t.Fail()
	}
}
