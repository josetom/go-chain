package fs_test

import (
	"testing"

	"github.com/josetom/go-chain/fs"
)

func TestSetConfig(t *testing.T) {
	testConfig := fs.FsConfig{
		DataDir: fs.Defaults().DataDir,
	}
	fs.SetConfig(testConfig)
	if fs.Config.DataDir != fs.Defaults().DataDir {
		t.Fail()
	}
	cleanup := func() {
		fs.Config = fs.Defaults()
	}
	t.Cleanup(cleanup)
}
