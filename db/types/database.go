package types

type Database interface {
	Put(key []byte, value []byte) error
	Get(key []byte) ([]byte, error)
	Has(key []byte) (bool, error)
	Delete(key []byte) error
	Close() error
	NewIterator(start []byte, limit []byte) Iterator
}
