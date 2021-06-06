package node

func getDefaultBootstrapNodes() []BootstrapNodeConfig {
	var bootstrapNodes = make([]BootstrapNodeConfig, 0)
	bnc := BootstrapNodeConfig{
		Host: "http://127.0.0.1:8600",
	}
	bootstrapNodes = append(bootstrapNodes, bnc)
	return bootstrapNodes
}

var Defaults = NodeConfig{
	Http: HttpConfig{
		Protocol: "http",
		Host:     "http://127.0.0.1:8600",
		Port:     8600,
	},
	BootstrapNodes: getDefaultBootstrapNodes(),
	IsBootstrap:    false,
	Sync: SyncConfig{
		TickerDuration: 45,
	},
}
