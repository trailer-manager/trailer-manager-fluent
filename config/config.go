package config

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Broker Broker `toml:"broker"`
}

type Broker struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	ClientId string `toml:"clientId"`
	Username string `toml:"username"`
	Password string `toml:"password"`
}

var config Config

func InitConfig() (err error) {
	_, err = toml.DecodeFile("./config/config.toml", &config)
	return
}

func GetConfig() (config Config) {
	return config
}