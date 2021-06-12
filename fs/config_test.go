package fs

import (
	"testing"
)

func TestSetFsConfig(t *testing.T) {
	testConfig := FsConfig{
		DataDir: Defaults().DataDir,
	}
	SetFsConfig(testConfig)
	if Config.DataDir != Defaults().DataDir {
		t.Fail()
	}
	cleanup := func() {
		Config = Defaults()
	}
	t.Cleanup(cleanup)
}
