package node

import (
	"testing"
)

func TestSetNodeConfig(t *testing.T) {
	testConfig := NodeConfig{
		IsBootstrap: Defaults.IsBootstrap,
	}
	SetNodeConfig(testConfig)
	if Config.IsBootstrap != Defaults.IsBootstrap {
		t.Fail()
	}
}
