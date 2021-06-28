package wallet_test

import (
	"testing"

	"github.com/josetom/go-chain/test_helper"
	"github.com/josetom/go-chain/test_helper/test_helper_core"
	"github.com/josetom/go-chain/wallet"
)

func TestNewKeystoreAccount(t *testing.T) {
	wallet.Config.DataDir = test_helper.CreateAndGetTestWalletDir(true)
	_, err := wallet.NewKeystoreAccount(test_helper.WALLET_PWD)
	if err != nil {
		t.Error(err)
	}
	cleanup := func() {
		test_helper.DeleteTestWalletDir()
		wallet.Config.DataDir = wallet.Defaults().DataDir
	}
	t.Cleanup(cleanup)
}

func TestSignForAddressAndVerify(t *testing.T) {
	wallet.Config.DataDir = test_helper.CreateAndGetTestWalletDir(false)
	txn := test_helper_core.GetTestTxn()
	txnBytes, err := txn.Encode()
	if err != nil {
		t.Error(err)
	}
	sig, err := wallet.SignForAddress(txnBytes, txn.From(), test_helper.WALLET_PWD)
	if err != nil {
		t.Error(err)
	}
	txn.WithSignature(sig)
	isAuthentic, err := txn.IsAuthentic()
	if err != nil || !isAuthentic {
		t.Error(err)
	}
}
