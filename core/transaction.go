package core

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/josetom/go-chain/common"
)

type TransactionData struct {
	From  common.Address `json:"from"`
	To    common.Address `json:"to"`
	Value uint           `json:"value"`
	Data  string         `json:"data"`
}

type TransactionContent struct {
	TxnData   TransactionData `json:"txnData"`
	Timestamp uint64          `json:"timestamp"`
	IsReward  bool            `json:"isReward"`
}

type Transaction struct {
	TxnContent TransactionContent `json:"txnContent"`
	TxnHash    common.Hash        `json:"txnhHsh"`
	Signature  common.Signature   `json:"signature"`
}

func NewTransaction(from common.Address, to common.Address, value uint, data string) Transaction {
	txnData := TransactionData{from, to, value, data}
	txn := Transaction{
		TxnContent: TransactionContent{
			TxnData:   txnData,
			Timestamp: uint64(time.Now().UnixNano()),
		},
	}
	txn.hash()
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

func (t *Transaction) WithSignature(signature common.Signature) {
	t.Signature = signature
}

func (t *Transaction) IsAuthentic() (bool, error) {
	txnBytes, err := t.Encode()
	if err != nil {
		return false, err
	}
	recoveredAddr, err := common.VerifyAndRecoverAccount(txnBytes, t.Signature)
	if err != nil {
		return false, err
	}
	if !recoveredAddr.Equal(t.From()) {
		return false, fmt.Errorf("singnature doesn't match of sender")
	}
	return true, nil
}

// func (t *Transaction) Sign() error {
// 	txnBytes, err := t.Encode()
// 	if err != nil {
// 		return err
// 	}
// 	sig
// 	t.WithSignature(sig)
// }
