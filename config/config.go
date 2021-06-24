package config

import (
	"github.com/josetom/go-chain/core"
	"github.com/josetom/go-chain/fs"
	"github.com/josetom/go-chain/node"
)

type Config struct {
	FS   fs.FsConfig     `yaml:"fs,omitempty"`
	Node node.NodeConfig `yaml:"node,omitempty"`
	Core core.CoreConfig `yaml:"core,omitempty"`
}
