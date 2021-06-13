package fs

import (
	"log"
	"testing"

	"github.com/josetom/go-chain/test_helper"
)

func TestDoesExist(t *testing.T) {
	if isExist, _ := DoesExist(test_helper.GetTestFile("valid.db")); !isExist {
		log.Println(test_helper.GetTestFile("valid.db"))
		t.Fail()
	}
	if isExist, _ := DoesExist(test_helper.GetTestFile("missing.db")); isExist {
		t.Fail()
	}
}
