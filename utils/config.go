package utils

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

var Cfg *Config

type Config struct {
	HostTarget []string `yaml:"rig_hosts_list"`
}

func LoadYamlConfig(ConfigFilePath string) {
	t := Config{}
	data, err := ioutil.ReadFile(ConfigFilePath)
	Check(err)
	err = yaml.Unmarshal(data, &t)
	Check(err)
	Cfg = &t
}
