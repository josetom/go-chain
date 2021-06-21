package node

import (
	"testing"
)

func TestSetConfig(t *testing.T) {
	testConfig := NodeConfig{
		IsBootstrap: Defaults().IsBootstrap,
	}
	SetConfig(testConfig)
	if Config.IsBootstrap != Defaults().IsBootstrap {
		t.Fail()
	}
	cleanup := func() {
		Config = Defaults()
	}
	t.Cleanup(cleanup)
}
