package fs

type FsConfig struct {
	DataDir string `yaml:"datadir,omitempty"`
}

var Config *FsConfig

func SetFsConfig(fsConfig *FsConfig) {
	Config = fsConfig
}
