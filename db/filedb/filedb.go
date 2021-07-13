package filedb

import (
	"fmt"
	"os"

	"github.com/josetom/go-chain/db/types"
)

// TODO : An extremely piss poor implementation of a file DB. Throw this away and do not use.
// Adding it only for some backward compatibility testing
type FileDB struct {
	dbPath string
	db     *os.File
}

func NewFileDB(dbPath string) (types.Database, error) {
	db, err := os.OpenFile(dbPath, os.O_APPEND|os.O_RDWR, 0600)
	if err != nil {
		return &FileDB{}, err
	}
	return &FileDB{
		dbPath: dbPath,
		db:     db,
	}, nil
}

// Has retrieves if a key is present in the key-value store.
func (db *FileDB) Has(key []byte) (bool, error) {
	return false, fmt.Errorf("not implemented")
}

// Get retrieves the given key if it's present in the key-value store.
func (db *FileDB) Get(key []byte) ([]byte, error) {
	return nil, fmt.Errorf("not implemented")
}

// Put inserts the given value into the key-value store.
func (db *FileDB) Put(key []byte, value []byte) error {
	_, err := db.db.Write(append(value, '\n'))
	return err
}

// Delete removes the key from the key-value store.
func (db *FileDB) Delete(key []byte) error {
	return fmt.Errorf("not implemented")
}

func (db *FileDB) Close() error {
	return db.db.Close()
}

// NewIterator creates a binary-alphabetical iterator over a subset
// of database content with a particular key prefix, starting at a particular
// initial key (or after, if it does not exist).
func (db *FileDB) NewIterator(start []byte, limit []byte) types.Iterator {
	return newFileDbIterator(db.dbPath, start, limit)
}
