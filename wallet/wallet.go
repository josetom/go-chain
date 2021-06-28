package wallet

import (
	"crypto/ecdsa"
	"io/ioutil"
	"path/filepath"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/josetom/go-chain/common"
	"github.com/josetom/go-chain/core"
	"github.com/josetom/go-chain/fs"
)

func getWalletDir() string {
	return filepath.Join(fs.ExpandPath(Config.DataDir), "keystore")
}

func NewKeystoreAccount(password string) (common.Address, error) {
	ks := keystore.NewKeyStore(getWalletDir(), keystore.StandardScryptN, keystore.StandardScryptP)
	acc, err := ks.NewAccount(password)
	if err != nil {
		return common.Address{}, err
	}
	return common.Address(acc.Address), nil
}

func SignTxWithKeystoreAccount(txn core.Transaction, password string) (core.Transaction, error) {
	ks := keystore.NewKeyStore(getWalletDir(), keystore.StandardScryptN, keystore.StandardScryptP)
	acc, err := ks.Find(accounts.Account{Address: ethcommon.Address(txn.From())})
	if err != nil {
		return core.Transaction{}, err
	}
	key, err := getKeyForAccount(acc, password)
	if err != nil {
		return core.Transaction{}, err
	}
	signedTx, err := SignTxn(txn, key.PrivateKey)
	if err != nil {
		return core.Transaction{}, err
	}
	return signedTx, nil
}

func getKeyForAccount(acc accounts.Account, password string) (*keystore.Key, error) {
	ksAccJson, err := ioutil.ReadFile(acc.URL.Path)
	if err != nil {
		return &keystore.Key{}, err
	}
	key, err := keystore.DecryptKey(ksAccJson, password)
	if err != nil {
		return &keystore.Key{}, err
	}
	return key, nil
}

func SignTxn(txn core.Transaction, pk *ecdsa.PrivateKey) (core.Transaction, error) {
	txnBytes, err := txn.Encode()
	if err != nil {
		return core.Transaction{}, err
	}
	signature, err := Sign(txnBytes, pk)
	if err != nil {
		return core.Transaction{}, nil
	}
	txn.WithSignature(signature)
	return txn, nil
}

func Sign(msg []byte, pk *ecdsa.PrivateKey) (common.Signature, error) {
	msgHash := common.BytesToHash(msg)
	sigBytes, err := crypto.Sign(msgHash.Bytes(), pk)
	if err != nil {
		return common.Signature{}, nil
	}
	var s common.Signature
	s.SetBytes(sigBytes)
	return s, nil
}
