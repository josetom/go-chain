package config

import (
	"io/ioutil"
	"log"

	"github.com/josetom/go-chain/core"
	"github.com/josetom/go-chain/node"
	"gopkg.in/yaml.v2"
)

var config Config

func Load(configFile string) Config {
	config = Defaults

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

	node.SetNodeConfig(&config.Node)
	core.SetCoreConfig(&config.Core)

	return config
}
