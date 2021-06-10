package fs

import "github.com/josetom/go-chain/common"

type FsConfig struct {
	DataDir string `yaml:"datadir,omitempty"`
}

var Config FsConfig

func init() {
	common.DeepCopy(Defaults, &Config)
}

func SetFsConfig(fsConfig FsConfig) {
	Config = fsConfig
}
