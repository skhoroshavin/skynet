package http

import (
	"os"
)

type Config struct {
	DomainName string
}

func EnvConfig() *Config {
	r := &Config{
		DomainName: "localhost",
	}

	val, ok := os.LookupEnv("SKYNET_API_DOMAIN")
	if ok {
		r.DomainName = val
	}

	return r
}
