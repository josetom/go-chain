package core

import "github.com/josetom/go-chain/common"

type CoreConfig struct {
	State StateConfig `yaml:"state,omitempty"`
}

type StateConfig struct {
	DbFile string `yaml:"dbfile,omitempty"`
}

var Config CoreConfig

func init() {
	common.DeepCopy(Defaults, &Config)
}

func SetCoreConfig(config CoreConfig) {
	Config = config
}
