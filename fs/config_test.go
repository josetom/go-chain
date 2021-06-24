package fs

import (
	"testing"
)

func TestSetConfig(t *testing.T) {
	testConfig := FsConfig{
		DataDir: Defaults().DataDir,
	}
	SetConfig(testConfig)
	if Config.DataDir != Defaults().DataDir {
		t.Fail()
	}
	cleanup := func() {
		Config = Defaults()
	}
	t.Cleanup(cleanup)
}
