package db

type DbConfig struct {
	Type string `yaml:"type,omitempty"`
}

var Config DbConfig = Defaults()

func SetConfig(config DbConfig) {
	Config = config
}
