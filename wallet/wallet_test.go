package wallet

import (
	"testing"

	"github.com/josetom/go-chain/test_helper"
	"github.com/josetom/go-chain/test_helper/test_helper_core"
)

func TestNewKeystoreAccount(t *testing.T) {
	Config.DataDir = test_helper.CreateAndGetTestWalletDir(true)
	_, err := NewKeystoreAccount(test_helper.WALLET_PWD)
	if err != nil {
		t.Fail()
	}
	cleanup := func() {
		test_helper.DeleteTestWalletDir()
		Config.DataDir = Defaults().DataDir
	}
	t.Cleanup(cleanup)
}

func TestSignTxWithKeystoreAccountAndVerify(t *testing.T) {
	Config.DataDir = test_helper.CreateAndGetTestWalletDir(false)
	txn := test_helper_core.GetTestTxn(false)
	signedTx, err := SignTxWithKeystoreAccount(txn, test_helper.WALLET_PWD)
	if err != nil {
		t.Fail()
	}
	isAuthentic, err := signedTx.IsAuthentic()
	if err != nil || !isAuthentic {
		t.Fail()
	}
}
