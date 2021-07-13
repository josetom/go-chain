package db_test

import (
	"testing"

	"github.com/josetom/go-chain/db/dbtest"
)

func TestDatabase(t *testing.T) {

	// new db
	database := dbtest.NewTestDatabase(t)

	// put
	err := database.Put([]byte("k1"), []byte("v1"))
	if err != nil {
		t.Error(err)
	}

	// has
	has, err := database.Has([]byte("k1"))
	if err != nil {
		t.Error(err)
	}
	if !has {
		t.Fail()
	}
	// has not
	has, err = database.Has([]byte("k2"))
	if err != nil {
		t.Error(err)
	}
	if has {
		t.Fail()
	}

	// get
	value, err := database.Get([]byte("k1"))
	if err != nil {
		t.Error(err)
	}
	if string(value) != "v1" {
		t.Fail()
	}

	// delete
	err = database.Delete([]byte("k1"))
	if err != nil {
		t.Error(err)
	}

	err = database.Close()
	if err != nil {
		t.Error(err)
	}

}
