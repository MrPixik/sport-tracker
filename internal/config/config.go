package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env        string `yaml:"env"`
	StorageDSN string `yaml:"storage"`
	HTTPServer `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address"`
	ReadTimeout time.Duration `yaml:"read_timeout"`
	IdleTimeout time.Duration `yaml:"idle_timeout"`
}

func MustLoad() Config {
	configPath := os.Getenv("AUTH_CONFIG_PATH")

	if configPath == "" {
		log.Fatal("environment variable AUTH_CONFIG_PATH does not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("no config file in %s", configPath)
	}

	var config Config
	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		log.Fatalf("unable to load config: \n%s", err.Error())
	}

	return config
}
