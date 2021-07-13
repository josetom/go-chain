package leveldb

import (
	"log"

	"github.com/josetom/go-chain/db/types"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

// TODO : Compact

type LevelDB struct {
	dbPath string
	db     *leveldb.DB
}

func NewLevelDB(dbPath string) (types.Database, error) {
	log.Println(dbPath) // TODO : clear
	db, err := leveldb.OpenFile(dbPath, nil)
	if err != nil {
		return &LevelDB{}, err
	}
	return &LevelDB{
		dbPath: dbPath,
		db:     db,
	}, nil
}

// Has retrieves if a key is present in the key-value store.
func (db *LevelDB) Has(key []byte) (bool, error) {
	return db.db.Has(key, nil)
}

// Get retrieves the given key if it's present in the key-value store.
func (db *LevelDB) Get(key []byte) ([]byte, error) {
	dat, err := db.db.Get(key, nil)
	if err != nil {
		return nil, err
	}
	return dat, nil
}

// Put inserts the given value into the key-value store.
func (db *LevelDB) Put(key []byte, value []byte) error {
	return db.db.Put(key, value, nil)
}

// Delete removes the key from the key-value store.
func (db *LevelDB) Delete(key []byte) error {
	return db.db.Delete(key, nil)
}

func (db *LevelDB) Close() error {
	return db.db.Close()
}

// NewIterator creates a binary-alphabetical iterator over a subset
// of database content with a particular key prefix, starting at a particular
// initial key (or after, if it does not exist).
func (db *LevelDB) NewIterator(start []byte, limit []byte) types.Iterator {
	r := util.Range{
		Start: start,
		Limit: limit,
	}
	return db.db.NewIterator(&r, nil)
}
