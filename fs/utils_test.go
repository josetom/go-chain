package fs_test

import (
	"testing"

	"github.com/josetom/go-chain/fs"
	"github.com/josetom/go-chain/test_helper"
)

func TestDoesExist(t *testing.T) {
	if isExist, _ := fs.DoesExist(test_helper.GetTestFile("valid.db")); !isExist {
		t.Error(test_helper.GetTestFile("valid.db"))
	}
	if isExist, _ := fs.DoesExist(test_helper.GetTestFile("missing.db")); isExist {
		t.Fail()
	}
}
