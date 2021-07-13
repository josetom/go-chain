package dbtest

import (
	"os"
	"testing"

	"github.com/josetom/go-chain/db"
	"github.com/josetom/go-chain/db/types"
)

func NewTestDatabase(t *testing.T) types.Database {
	t_db, err := db.NewDatabase("__test")
	if err != nil {
		t.Error(err)
	}

	cleanup := func() {
		t_db.Close()
		os.RemoveAll("__test")
	}
	t.Cleanup(cleanup)

	return t_db
}
