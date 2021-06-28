package test_helper

import (
	"log"
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

func CreateAndGetTestDbFile() string {
	os.MkdirAll(GetTestFile("database/temp"), os.ModePerm)
	os.Create(GetTestFile("database/temp/test.db"))
	return "temp/test.db"
}

func DeleteTestDbFile() {
	err := os.Remove(GetTestFile("database/temp/test.db"))
	if err != nil {
		log.Println(err)
	}
}

func CreateAndGetTestWalletDir(isTemp bool) string {
	if isTemp {
		dir := GetTestFile("temp/wallet")
		os.MkdirAll(dir, os.ModePerm)
		return dir
	} else {
		return GetTestDataDir()
	}
}

func DeleteTestWalletDir() {
	err := os.RemoveAll(GetTestFile("temp/wallet"))
	if err != nil {
		log.Println(err)
	}
}
