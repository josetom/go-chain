package leveldb

import (
	"github.com/josetom/go-chain/db/errors"
	"github.com/josetom/go-chain/db/types"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

// TODO : Compact

type levelDB struct {
	dbPath string
	db     *leveldb.DB
}

var levelDbNotFoundErrorMessage = "leveldb: not found"

func NewLevelDB(dbPath string) (types.Database, error) {
	db, err := leveldb.OpenFile(dbPath, nil)
	if err != nil {
		return &levelDB{}, err
	}
	return &levelDB{
		dbPath: dbPath,
		db:     db,
	}, nil
}

// Has retrieves if a key is present in the key-value store.
func (db *levelDB) Has(key []byte) (bool, error) {
	return db.db.Has(key, nil)
}

// Get retrieves the given key if it's present in the key-value store.
func (db *levelDB) Get(key []byte) ([]byte, error) {
	dat, err := db.db.Get(key, nil)
	if err != nil {
		if err.Error() == levelDbNotFoundErrorMessage {
			return nil, errors.NotFoundError
		}
		return nil, err
	}
	return dat, nil
}

// Put inserts the given value into the key-value store.
func (db *levelDB) Put(key []byte, value []byte) error {
	return db.db.Put(key, value, nil)
}

// Delete removes the key from the key-value store.
func (db *levelDB) Delete(key []byte) error {
	return db.db.Delete(key, nil)
}

func (db *levelDB) Close() error {
	return db.db.Close()
}

// NewIterator creates a binary-alphabetical iterator over a subset
// of database content with a particular key prefix, starting at a particular
// initial key (or after, if it does not exist).
func (db *levelDB) NewIterator(prefix []byte, start []byte, limit []byte) types.Iterator {
	r := util.BytesPrefix(prefix)
	r.Start = append(r.Start, start...)
	r.Limit = append(r.Limit, limit...)
	return db.db.NewIterator(r, nil)
}

// NewBatch creates a write-only key-value store that buffers changes to its host
// database until a final write is called.
func (db *levelDB) NewBatch() types.Batch {
	return &leveldbBatch{
		db: db.db,
		b:  new(leveldb.Batch),
	}
}
