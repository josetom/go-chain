package core

import (
	"testing"
)

func TestSetConfig(t *testing.T) {
	testConfig := CoreConfig{
		State: StateConfig{
			DbFile: Defaults().State.DbFile,
		},
	}
	SetConfig(testConfig)
	if Config.State.DbFile != Defaults().State.DbFile {
		t.Fail()
	}
	cleanup := func() {
		Config = Defaults()
	}
	t.Cleanup(cleanup)
}
