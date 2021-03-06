package core

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/josetom/go-chain/common"
	"github.com/josetom/go-chain/wallet"
)

type TransactionData struct {
	From  common.Address `json:"from"`
	To    common.Address `json:"to"`
	Value uint           `json:"value"`
	Data  string         `json:"data"`
}

type TransactionContent struct {
	TxnData   TransactionData `json:"txnData"`
	TxnFee    uint            `json:"txnFee"`
	Timestamp uint64          `json:"timestamp"`
}

type Transaction struct {
	TxnContent TransactionContent `json:"txnContent"`
	TxnHash    common.Hash        `json:"txnHash"`
	Signature  common.Signature   `json:"signature"`
}

func NewTransaction(from common.Address, to common.Address, value uint, data string) Transaction {
	txnData := TransactionData{from, to, value, data}
	txn := Transaction{
		TxnContent: TransactionContent{
			TxnData:   txnData,
			TxnFee:    getTxnFee(value),
			Timestamp: uint64(time.Now().UnixNano()),
		},
	}
	txn.hash()
	// TODO : check when to sign
	txn.Sign()
	return txn
}

func (t *Transaction) Encode() ([]byte, error) {
	return json.Marshal(t.TxnContent)
}

func (t *Transaction) hash() error {
	bytes, err := t.Encode()
	if err == nil {
		t.TxnHash = common.BytesToHash(bytes)
		return nil
	}
	return err
}

func (t *Transaction) Hash() common.Hash {
	t.hash()
	return t.TxnHash
}

func (t *Transaction) From() common.Address {
	return t.TxnContent.TxnData.From
}

func (t *Transaction) To() common.Address {
	return t.TxnContent.TxnData.To
}

func (t *Transaction) Value() uint {
	return t.TxnContent.TxnData.Value
}

func (t *Transaction) Data() string {
	return t.TxnContent.TxnData.Data
}

func (t *Transaction) Fee() uint {
	return t.TxnContent.TxnFee
}

func (t *Transaction) Cost() uint {
	return t.Value() + t.Fee()
}

func (t *Transaction) WithSignature(signature common.Signature) {
	t.Signature = signature
}

func (t *Transaction) IsAuthentic() (bool, error) {
	txnBytes, err := t.Encode()
	if err != nil {
		return false, err
	}
	recoveredAddr, err := wallet.VerifyAndRecoverAccount(txnBytes, t.Signature)
	if err != nil {
		return false, err
	}
	if !recoveredAddr.Equal(t.From()) {
		return false, fmt.Errorf("singnature doesn't match of sender")
	}
	return true, nil
}

func (t *Transaction) Sign() error {
	txnBytes, err := t.Encode()
	if err != nil {
		return err
	}
	// TODO : password needs to be changed
	sig, err := wallet.SignForAddress(txnBytes, t.From(), "wallet_pwd")
	if err != nil {
		return err
	}
	t.WithSignature(sig)
	return nil
}

// TODO : refine this logic
func getTxnFee(value uint) uint {
	if value > 1000 {
		return 10
	} else {
		return 0
	}
}
