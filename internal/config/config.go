package config

import (
	"flag"
	"fmt"
	"strings"

	"github.com/caarlos0/env/v9"
)

type Config struct {
	ServerAddress string
	BaseURL       string
}

func (c *Config) parseFlags() {
	flag.StringVar(&c.ServerAddress, "a", "localhost:8080", "address and port to run server")
	flag.StringVar(&c.BaseURL, "b", "localhost:8080", "base url for short link")
	flag.Parse()
}

func (c *Config) parseEnv() {
	type EnvVariables struct {
		ServerAddress string `env:"SERVER_ADDRESS"`
		BaseURL       string `env:"BASE_URL"`
	}
	var envVariables EnvVariables
	err := env.Parse(&envVariables)
	if err != nil {
		return
	}
	if len(envVariables.BaseURL) > 0 {
		c.BaseURL = envVariables.BaseURL
	}
	if len(envVariables.ServerAddress) > 0 {
		c.ServerAddress = envVariables.ServerAddress
	}
}

func (c *Config) prepareBaseURL() {
	if len(c.BaseURL) == 0 {
		return
	}
	if !strings.HasPrefix(c.BaseURL, "http") {
		c.BaseURL = fmt.Sprintf(`http://%s`, c.BaseURL)
	}
}

func New() Config {
	c := Config{}
	c.parseFlags()
	c.parseEnv()
	c.prepareBaseURL()
	return c
}
