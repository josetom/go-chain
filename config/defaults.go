package config

import (
	"github.com/josetom/go-chain/core"
	"github.com/josetom/go-chain/fs"
	"github.com/josetom/go-chain/node"
	"github.com/josetom/go-chain/wallet"
)

var Defaults = Config{
	FS:     fs.Defaults(),
	Node:   node.Defaults(),
	Core:   core.Defaults(),
	Wallet: wallet.Defaults(),
}
