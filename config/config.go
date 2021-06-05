package config

import (
	"github.com/josetom/go-chain/core"
	"github.com/josetom/go-chain/node"
)

type Config struct {
	Node node.NodeConfig `yaml:"node"`
	Core core.CoreConfig `yaml:"core"`
}
