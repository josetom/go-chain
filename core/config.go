package core

type CoreConfig struct {
	State StateConfig `yaml:"state,omitempty"`
}

type StateConfig struct {
	DbFile string `yaml:"dbfile,omitempty"`
}

var Config *CoreConfig

func SetCoreConfig(config *CoreConfig) {
	Config = config
}
