package config

import (
	"log"
	"testing"

	"github.com/josetom/go-chain/common"
	"github.com/josetom/go-chain/core"
	"github.com/josetom/go-chain/fs"
	"github.com/josetom/go-chain/test_helper"
)

func TestLoadDefaults(t *testing.T) {
	config := Load("")
	if config.FS.DataDir != fs.Defaults().DataDir {
		log.Println("config.fs.DataDir -- ", config.FS.DataDir)
		t.Fail()
	}
	if config.Core.State.DbFile != core.Defaults().State.DbFile {
		log.Println("config.Core.State.DbFile -- ", config.Core.State.DbFile)
		t.Fail()
	}
	cleanup := func() {
		common.DeepCopy(Defaults, config)
	}
	t.Cleanup(cleanup)
}

func TestLoadValidFile(t *testing.T) {
	config := Load(test_helper.GetTestFile("valid-config.yaml"))
	if config.FS.DataDir != "dummy_dir" {
		log.Println("config.fs.DataDir -- ", config.FS.DataDir)
		t.Fail()
	}
	if config.Core.State.DbFile != "dummy_state_file" {
		log.Println("config.Core.State.DbFile -- ", config.Core.State.DbFile)
		t.Fail()
	}
	cleanup := func() {
		common.DeepCopy(Defaults, config)
	}
	t.Cleanup(cleanup)
}

func TestLoadPartialConfig(t *testing.T) {
	config := Load(test_helper.GetTestFile("partial-config.yaml"))
	if config.FS.DataDir != fs.Defaults().DataDir {
		log.Println("config.fs.DataDir -- ", config.FS.DataDir)
		t.Fail()
	}
	if config.Core.State.DbFile != "dummy_state_file" {
		log.Println("config.Core.State.DbFile -- ", config.Core.State.DbFile)
		t.Fail()
	}
	cleanup := func() {
		common.DeepCopy(Defaults, config)
	}
	t.Cleanup(cleanup)
}
