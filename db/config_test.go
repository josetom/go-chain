package db

import (
	"testing"
)

func TestSetConfig(t *testing.T) {
	testConfig := DbConfig{
		Type: Defaults().Type,
	}
	SetConfig(testConfig)
	if Config.Type != Defaults().Type {
		t.Fail()
	}
	cleanup := func() {
		Config = Defaults()
	}
	t.Cleanup(cleanup)
}
