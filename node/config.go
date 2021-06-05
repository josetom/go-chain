package node

type NodeConfig struct {
	DataDir  string `yaml:"datadir,omitempty"`
	HttpPort int    `yaml:"httpPort,omitempty"`
}

var Config *NodeConfig

func SetNodeConfig(nodeConfig *NodeConfig) {
	Config = nodeConfig
}
