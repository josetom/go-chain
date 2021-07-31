package node

import "github.com/josetom/go-chain/common"

func getDefaultBootstrapNodes() []BootstrapNodeConfig {
	var bootstrapNodes = make([]BootstrapNodeConfig, 0)
	bnc := BootstrapNodeConfig{
		Host: "http://testnet.unigate.network",
	}
	bootstrapNodes = append(bootstrapNodes, bnc)
	return bootstrapNodes
}

func Defaults() NodeConfig {
	return NodeConfig{
		Http: HttpConfig{
			Protocol: "http",
			Host:     "http://127.0.0.1:8600",
			Port:     8600,
		},
		BootstrapNodes: getDefaultBootstrapNodes(),
		IsBootstrap:    false,
		Sync: SyncConfig{
			TickerDuration: 30,
		},
		Miner: MinerConfig{
			TickerDuration: 60,
			Address:        common.NewAddress("0x38c4ce5b96044e7ddd9ad95947178ed3436a4539"),
		},
	}
}
