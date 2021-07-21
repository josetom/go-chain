package types

type Record struct {
	Key   []byte `json:"hash"`
	Value []byte `json:"block"`
}
