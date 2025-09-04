package main

import (
	"fmt"
	"github.com/mrpixik/sport-tracker/internal/auth/config"
	"github.com/mrpixik/sport-tracker/pkg/logger"
)

func main() {
	// $env:auth_config_path="configs/auth/local.yaml"
	cfg := config.MustLoad()

	fmt.Println(cfg)

	log := logger.MustInit(cfg.Env)

	log.Info("Logger initialized")
	log.Debug("Debug messages are enabled")

	//TODO: init storage

	//TODO: init router

	//TODO: init server
}
