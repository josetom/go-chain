package fs

type FsConfig struct {
	DataDir string `yaml:"datadir,omitempty"`
}

var Config FsConfig = Defaults()

func SetConfig(fsConfig FsConfig) {
	Config = fsConfig
}
