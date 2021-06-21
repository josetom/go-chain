package core

type CoreConfig struct {
	State StateConfig `yaml:"state,omitempty"`
	Block BlockConfig `yaml:"block,omitempty"`
}

type StateConfig struct {
	DbFile string `yaml:"dbfile,omitempty"`
}

type BlockConfig struct {
	Reward     uint64 `yaml:"reward,omitempty"`
	Complexity uint64 `yaml:"complexity,omitempty"`
}

var Config CoreConfig = Defaults()

func SetConfig(config CoreConfig) {
	Config = config
}
