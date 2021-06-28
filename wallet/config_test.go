package wallet

import (
	"testing"
)

func TestSetConfig(t *testing.T) {
	testConfig := WalletConfig{
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
