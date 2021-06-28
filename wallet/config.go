package wallet

type WalletConfig struct {
	DataDir string `yaml:"datadir,omitempty"`
}

var Config WalletConfig = Defaults()

func SetConfig(walletConfig WalletConfig) {
	Config = walletConfig
}
