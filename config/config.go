package config

import (
	"SiverPineValley/trailer-manager/common"
	"SiverPineValley/trailer-manager/utility"
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
)

const (
	path = "./config/config_%s.toml"
)

type Config struct {
	Broker       Broker   `toml:"broker"`
	Rdb          Rdb      `toml:"rdb"`
	Port         int      `toml:"port"`
	Processname  string   `toml:"process_name"`
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

var config *Config

func InitConfig(mode string) (err error) {
	if !utility.Contains([]string{common.ModeLocal, common.ModeDevelopment, common.ModeStaging, common.ModeProduction}, mode) {

	}

	_, err = toml.DecodeFile(fmt.Sprintf(path, mode), &config)
	if err != nil {
		log.Fatal(err)
		return
	}
	return
}

func GetConfig() (Config) {
	return *config
}
