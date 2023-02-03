package config

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
