package fs

import "testing"

func TestDoesExist(t *testing.T) {
	if isExist, _ := DoesExist("testdata/valid.db"); !isExist {
		t.Fail()
	}
	if isExist, _ := DoesExist("testdata/missing.db"); isExist {
		t.Fail()
	}
}
