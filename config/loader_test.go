package config

import (
	"log"
	"testing"

	"github.com/josetom/go-chain/core"
	"github.com/josetom/go-chain/node"
)

func TestLoadDefaults(t *testing.T) {
	config := Load("")
	if config.Node.DataDir != node.Defaults.DataDir {
		log.Println("config.Node.DataDir -- ", config.Node.DataDir)
		t.Fail()
	}
	if config.Core.State.DbFile != core.Defaults.State.DbFile {
		log.Println("config.Core.State.DbFile -- ", config.Core.State.DbFile)
		t.Fail()
	}
	cleanup := func() {
		config = Defaults
	}
	t.Cleanup(cleanup)
}

func TestLoadValidFile(t *testing.T) {
	config := Load("testdata/valid-config.yaml")
	if config.Node.DataDir != "dummy_dir" {
		log.Println("config.Node.DataDir -- ", config.Node.DataDir)
		t.Fail()
	}
	if config.Core.State.DbFile != "dummy_state_file" {
		log.Println("config.Core.State.DbFile -- ", config.Core.State.DbFile)
		t.Fail()
	}
	cleanup := func() {
		config = Defaults
	}
	t.Cleanup(cleanup)
}

func TestLoadPartialConfig(t *testing.T) {
	config := Load("testdata/partial-config.yaml")
	if config.Node.DataDir != node.Defaults.DataDir {
		log.Println("config.Node.DataDir -- ", config.Node.DataDir)
		t.Fail()
	}
	if config.Core.State.DbFile != "dummy_state_file" {
		log.Println("config.Core.State.DbFile -- ", config.Core.State.DbFile)
		t.Fail()
	}
	cleanup := func() {
		config = Defaults
	}
	t.Cleanup(cleanup)
}
