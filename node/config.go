package node

import "github.com/josetom/go-chain/common"

type NodeConfig struct {
	Http           HttpConfig            `yaml:"http,omitempty"`
	BootstrapNodes []BootstrapNodeConfig `yaml:"bootstrapNodes,omitempty"`
	IsBootstrap    bool                  `yaml:"isBootstrap,omitempty"`
	Sync           SyncConfig            `yaml:"sync,omitempty"`
	Miner          MinerConfig           `yaml:"miner,omitempty"`
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

type MinerConfig struct {
	Address        common.Address `yaml:"address,omitempty"`
	TickerDuration uint64         `yaml:"tickerDuration,omitempty"`
}

var Config NodeConfig = Defaults()

func SetConfig(nodeConfig NodeConfig) {
	Config = nodeConfig
}
