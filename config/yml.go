package config

import (
	"io/ioutil"

	"gitlab.com/gaming0skar123/go/cdn/common"
	"gopkg.in/yaml.v2"
)

type YmlConfig struct {
	Discord YmlConfigDiscord
}

type YmlConfigDiscord struct {
	API_Channel string
	Channels    []string
}

var API_Channel string
var Channels []string

func init() {
	f, err := ioutil.ReadFile("config.yml")
	if common.CheckErr(err, "read config file") {
		return
	}

	var data YmlConfig
	err = yaml.Unmarshal(f, &data)
	if common.CheckErr(err, "unmarshal config file") {
		return
	}

	API_Channel = data.Discord.API_Channel
	Channels = data.Discord.Channels
}
