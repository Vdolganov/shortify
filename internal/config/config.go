package config

import (
	"flag"

	"github.com/caarlos0/env/v9"
)

type Config struct {
	ServerAddress string
	BaseUrl       string
}

func (c *Config) parseFlags() {
	flag.StringVar(&c.ServerAddress, "a", "localhost:8080", "address and port to run server")
	flag.StringVar(&c.BaseUrl, "b", "localhost:8080", "base url for short link")
	flag.Parse()
}

func (c *Config) parseEnv() {
	type EnvVariables struct {
		ServerAddress string `env:"SERVER_ADDRESS"`
		BaseUrl       string `env:"BASE_URL"`
	}
	var envVariables EnvVariables
	err := env.Parse(&envVariables)
	if err != nil {
		return
	}
	if len(envVariables.BaseUrl) > 0 {
		c.BaseUrl = envVariables.BaseUrl
	}
	if len(envVariables.ServerAddress) > 0 {
		c.ServerAddress = envVariables.ServerAddress
	}
}

func InitConfig() Config {
	c := Config{}
	c.parseFlags()
	c.parseEnv()
	return c
}
