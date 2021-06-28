package wallet

import (
	"github.com/josetom/go-chain/fs"
)

func defaultDataDir() string {
	return fs.Defaults().DataDir
}

func Defaults() WalletConfig {
	return WalletConfig{
		DataDir: defaultDataDir(),
	}
}
