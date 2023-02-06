package config

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Broker       Broker   `toml:"broker"`
	Rdb          Rdb      `toml:"rdb"`
	Port         int      `toml:"port"`
	Processname  string   `toml:"processname"`
	AllowOrigins []string `toml:"allow_origins"`
}

type Broker struct {
	Host     string   `toml:"host"`
	Port     int      `toml:"port"`
	ClientId string   `toml:"clientId"`
	Username string   `toml:"username"`
	Password string   `toml:"password"`
	Topics   []string `toml:"topics"`
}

type Rdb struct {
	Driver   string `toml:"driver"`
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	DbName   string `toml:"dbname"`
	Username string `toml:"username"`
	Password string `toml:"password"`
}

type Server struct {
	Host         string `toml:"host"`
	Port         int    `toml:"port"`
	PrefixUrl    string `toml:"prefixUrl"`
	ClientId     string `toml:"clientId"`
	ClientSecret string `toml:"clientSecret"`
}

var config Config

func InitConfig() (err error) {
	_, err = toml.DecodeFile("./config/config.toml", &config)
	return
}

func GetConfig() (config Config) {
	return config
}
