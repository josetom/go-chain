package test_helper

import (
	"os"
	"path"
	"path/filepath"
)

func GetTestDataDir() string {
	cwd, _ := os.Getwd()
	p := filepath.Join(cwd, "../test_helper/testdata")
	return path.Clean(os.ExpandEnv(p))
}

func GetTestFile(p string) string {
	return filepath.Join(GetTestDataDir(), p)
}
