package config

import (
	"github.com/josetom/go-chain/core"
	"github.com/josetom/go-chain/node"
)

var Defaults = Config{
	Node: node.Defaults,
	Core: core.Defaults,
}
