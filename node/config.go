package node

type NodeConfig struct {
	DataDir string `yaml :"datadir,omitempty"`
}

var Config *NodeConfig

func SetNodeConfig(nodeConfig *NodeConfig) {
	Config = nodeConfig
}
