package config

import (
	"github.com/caarlos0/env/v11"
	"log"
	"time"
)

type Config struct {
	Env        string `env:"ENVIRONMENT" envDefault:"local"`
	StorageDSN string `env:"STORAGE_DSN,required"`
	HTTPServer
}

type HTTPServer struct {
	Address     string        `env:"SERVER_ADDR" envDefault:"localhost:8080"`
	ReadTimeout time.Duration `env:"READ_TIMEOUT" envDefault:"4s"`
	IdleTimeout time.Duration `env:"IDLE_TIMEOUT" envDefault:"120s"`
}

func MustLoad() Config {
	var config Config

	err := env.Parse(&config)
	if err != nil {
		log.Fatalf("unable to load config: \n%s", err.Error())
	}

	return config
}
