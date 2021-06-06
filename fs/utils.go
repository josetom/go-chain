package fs

import (
	"io/ioutil"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/josetom/go-chain/constants"
)

// Expands a file path
// 1. replace tilde with users home dir
// 2. expands embedded environment variables
// 3. cleans the path, e.g. /a/b/../c -> /a/c
// Note, it has limitations, e.g. ~someuser/tmp will not be expanded
func ExpandPath(p string) string {
	if i := strings.Index(p, ":"); i > 0 {
		return p
	}
	if i := strings.Index(p, "@"); i > 0 {
		return p
	}
	if strings.HasPrefix(p, "~/") || strings.HasPrefix(p, "~\\") {
		if home := homeDir(); home != "" {
			p = home + p[1:]
		}
	}
	return path.Clean(os.ExpandEnv(p))
}

func DoesExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false, err
	}
	return true, nil
}

func WriteEmptyBlocksDbToDisk(path string) error {
	return ioutil.WriteFile(path, []byte(""), os.ModePerm)
}

func defaultDataDir() string {
	// Try to place the data folder in the user's home dir
	home := homeDir()
	if home != "" {
		switch runtime.GOOS {
		case "darwin":
			return filepath.Join(home, "Library", constants.BlockChainName)
		case "windows":
			// We used to put everything in %HOME%\AppData\Roaming, but this caused
			// problems with non-typical setups. If this fallback location exists and
			// is non-empty, use it, otherwise DTRT and check %LOCALAPPDATA%.
			fallback := filepath.Join(home, "AppData", "Roaming", constants.BlockChainName)
			appdata := windowsAppData()
			if appdata == "" || isNonEmptyDir(fallback) {
				return fallback
			}
			return filepath.Join(appdata, constants.BlockChainName)
		default:
			return filepath.Join(home, ".", constants.BlockChainName)
		}
	}
	// As we cannot guess a stable location, return empty and handle later
	return ""
}

func windowsAppData() string {
	v := os.Getenv("LOCALAPPDATA")
	if v == "" {
		// Windows XP and below don't have LocalAppData. Crash here because
		// we don't support Windows XP and undefining the variable will cause
		// other issues.
		panic("environment variable LocalAppData is undefined")
	}
	return v
}

func isNonEmptyDir(dir string) bool {
	f, err := os.Open(dir)
	if err != nil {
		return false
	}
	names, _ := f.Readdir(1)
	f.Close()
	return len(names) > 0
}

func homeDir() string {
	if home := os.Getenv("HOME"); home != "" {
		return home
	}
	if usr, err := user.Current(); err == nil {
		return usr.HomeDir
	}
	return ""
}
