package wallet_test

import (
	"testing"

	"github.com/josetom/go-chain/wallet"
)

func TestSetConfig(t *testing.T) {
	testConfig := wallet.WalletConfig{
		DataDir: wallet.Defaults().DataDir,
	}
	wallet.SetConfig(testConfig)
	if wallet.Config.DataDir != wallet.Defaults().DataDir {
		t.Fail()
	}
	cleanup := func() {
		wallet.Config = wallet.Defaults()
	}
	t.Cleanup(cleanup)
}
