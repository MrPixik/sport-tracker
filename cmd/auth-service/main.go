package main

import (
	"github.com/mrpixik/sport-tracker/internal/auth/config"
	"github.com/mrpixik/sport-tracker/internal/auth/storage/postgres"
	"github.com/mrpixik/sport-tracker/pkg/logger"
)

func main() {
	// $env:auth_config_path="configs/auth/local.yaml"
	cfg := config.MustLoad()

	//fmt.Println(cfg)

	log := logger.MustInit(cfg.Env)

	log.Info("Logger initialized")
	log.Debug("Debug messages are enabled")

	db, err := postgres.New(cfg.StorageDSN)
	if err != nil {
		log.Error("Failed to init storage", "error", err)
	}
	_ = db
	//TODO: init router

	//TODO: init server
}
