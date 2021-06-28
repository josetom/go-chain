package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/josetom/go-chain/common"
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

func SignForAddress(msg []byte, address common.Address, password string) (common.Signature, error) {
	ks := keystore.NewKeyStore(getWalletDir(), keystore.StandardScryptN, keystore.StandardScryptP)
	acc, err := ks.Find(accounts.Account{Address: ethcommon.Address(address)})
	if err != nil {
		return common.Signature{}, err
	}
	key, err := getKeyForAccount(acc, password)
	if err != nil {
		return common.Signature{}, err
	}
	return Sign(msg, key.PrivateKey)
}

func VerifyAndRecoverAccount(msg []byte, signature common.Signature) (common.Address, error) {
	msgHash := common.BytesToHash(msg)
	recoveredPubKey, err := crypto.SigToPub(msgHash.Bytes(), signature.Bytes())
	if err != nil {
		return common.Address{}, fmt.Errorf("unable to verify signature. %s", err.Error())
	}
	recoveredPubKeyBytes := elliptic.Marshal(crypto.S256(), recoveredPubKey.X, recoveredPubKey.Y)
	recoveredPubKeyBytesHash := crypto.Keccak256(recoveredPubKeyBytes[1:])
	recoveredAccount := common.BytesToAddress(recoveredPubKeyBytesHash[12:])

	return recoveredAccount, nil
}
