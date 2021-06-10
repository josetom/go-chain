package node

import "github.com/josetom/go-chain/common"

type NodeConfig struct {
	Http           HttpConfig            `yaml:"http,omitempty"`
	BootstrapNodes []BootstrapNodeConfig `yaml:"bootstrapNodes,omitempty"`
	IsBootstrap    bool                  `yaml:"isBootstrap,omitempty"`
	Sync           SyncConfig            `yaml:"sync,omitempty"`
}

type HttpConfig struct {
	Protocol string `yaml:"protocol,omitempty"`
	Host     string `yaml:"host,omitempty"`
	Port     uint64 `yaml:"port,omitempty"`
}

type BootstrapNodeConfig struct {
	Host string `yaml:"host,omitempty"`
}

type SyncConfig struct {
	TickerDuration uint64 `yaml:"tickerDuration,omitempty"`
}

var Config NodeConfig

func init() {
	common.DeepCopy(Defaults, &Config)
}

func SetNodeConfig(nodeConfig NodeConfig) {
	Config = nodeConfig
}
