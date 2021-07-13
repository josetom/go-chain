package db

import (
	"github.com/josetom/go-chain/db/filedb"
	"github.com/josetom/go-chain/db/leveldb"
	"github.com/josetom/go-chain/db/types"
	"github.com/josetom/go-chain/fs"
)

func NewDatabase(dbPath string) (types.Database, error) {
	dbPath = fs.ExpandPath(dbPath)
	if Config.Type == LEVEL_DB {
		return leveldb.NewLevelDB(dbPath)
	}
	return filedb.NewFileDB(dbPath)
}
