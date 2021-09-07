package utils

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"t-rex_wrapper/logger"
)

var Cfg *Config

var CfgGrafana *GrafanaConfig

type Config struct {
	ApiPort    int      `yaml:"api_port"`
	ApiLogFile string   `yaml:"api_log_file"`
	FrontEnd   bool     `yaml:"frontend_enabled"`
	HostTarget []string `yaml:"rig_hosts_list"`
	Username   string   `yaml:"api_username"`
	Password   string   `yaml:"api_password"`
}

type GrafanaConfig struct {
	Username   string `yaml:"api_username"`
	Password   string `yaml:"api_password"`
	EsUser     string `yaml:"es_username"`
	EsPassword string `yaml:"es_password"`
}

func LoadYamlConfig(ConfigFilePath string) {
	t := Config{}
	data, err := ioutil.ReadFile(ConfigFilePath)
	logger.Error(err)
	err = yaml.Unmarshal(data, &t)
	logger.Error(err)
	Cfg = &t
}

func LoadGrafanaYamlConfig(ConfigFilePath string) {
	t := GrafanaConfig{}
	data, err := ioutil.ReadFile(ConfigFilePath)
	logger.Error(err)
	err = yaml.Unmarshal(data, &t)
	logger.Error(err)
	CfgGrafana = &t
}
