package config

import (
	"io/ioutil"
	"log"

	"github.com/josetom/go-chain/common"
	"github.com/josetom/go-chain/core"
	"github.com/josetom/go-chain/fs"
	"github.com/josetom/go-chain/node"
	"gopkg.in/yaml.v2"
)

var config Config

func Load(configFile string) Config {
	common.DeepCopy(Defaults, &config)

	content, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Println("Using default configuration")
	}

	if content != nil {
		err = yaml.Unmarshal(content, &config)
		if err != nil {
			log.Fatalln("error while reading config file")
		}
	}

	fs.SetConfig(config.FS)
	node.SetConfig(config.Node)
	core.SetConfig(config.Core)

	return config
}
