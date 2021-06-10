package core

import (
	"testing"
)

func TestSetCoreConfig(t *testing.T) {
	testConfig := CoreConfig{
		State: StateConfig{
			DbFile: Defaults.State.DbFile,
		},
	}
	SetCoreConfig(testConfig)
	if Config.State.DbFile != Defaults.State.DbFile {
		t.Fail()
	}
}
