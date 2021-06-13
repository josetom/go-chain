package fs

func Defaults() FsConfig {
	return FsConfig{
		DataDir: defaultDataDir(),
	}
}
