package config

import "flag"

type Config struct {
	AppAddress       string
	ShortUrlBaseAddr string
}

func (c *Config) parseFlags() {
	flag.StringVar(&c.AppAddress, "a", "localhost:8080", "address and port to run server")
	flag.StringVar(&c.ShortUrlBaseAddr, "b", "localhost:8080", "base url for short link")
	flag.Parse()
}

func InitConfig() Config {
	c := Config{}
	c.parseFlags()
	return c
}
