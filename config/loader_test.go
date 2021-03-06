package config

import (
	"testing"

	"github.com/josetom/go-chain/common"
	"github.com/josetom/go-chain/core"
	"github.com/josetom/go-chain/fs"
	"github.com/josetom/go-chain/test_helper"
)

func TestLoadDefaults(t *testing.T) {
	config := Load("")
	if config.FS.DataDir != fs.Defaults().DataDir {
		t.Error("config.fs.DataDir -- ", config.FS.DataDir)
	}
	if config.Core.State.DbFile != core.Defaults().State.DbFile {
		t.Error("config.Core.State.DbFile -- ", config.Core.State.DbFile)
	}
	cleanup := func() {
		common.DeepCopy(Defaults, config)
	}
	t.Cleanup(cleanup)
}

func TestLoadValidFile(t *testing.T) {
	config := Load(test_helper.GetTestFile("valid-config.yaml"))
	if config.FS.DataDir != "dummy_dir" {
		t.Error("config.fs.DataDir -- ", config.FS.DataDir)
	}
	if config.Core.State.DbFile != "dummy_state_file" {
		t.Error("config.Core.State.DbFile -- ", config.Core.State.DbFile)
	}
	cleanup := func() {
		common.DeepCopy(Defaults, config)
	}
	t.Cleanup(cleanup)
}

func TestLoadPartialConfig(t *testing.T) {
	config := Load(test_helper.GetTestFile("partial-config.yaml"))
	if config.FS.DataDir != fs.Defaults().DataDir {
		t.Error("config.fs.DataDir -- ", config.FS.DataDir)
	}
	if config.Core.State.DbFile != "dummy_state_file" {
		t.Error("config.Core.State.DbFile -- ", config.Core.State.DbFile)
	}
	cleanup := func() {
		common.DeepCopy(Defaults, config)
	}
	t.Cleanup(cleanup)
}
