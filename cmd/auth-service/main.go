package main

import (
	"github.com/mrpixik/sport-tracker/internal/auth/config"
	"github.com/mrpixik/sport-tracker/internal/auth/http-server/server"
	"github.com/mrpixik/sport-tracker/internal/auth/storage/postgres"
	"github.com/mrpixik/sport-tracker/pkg/logger"
	"os"
)

func main() {
	cfg := config.MustLoad()

	log := logger.MustInit(cfg.Env)

	log.Info("Logger initialized")
	log.Debug("Debug messages are enabled")

	db, err := postgres.New(cfg.StorageDSN)
	if err != nil {
		log.Error("Failed to init storage", "error", err)
		os.Exit(1)
	}
	log.Debug("Successful initialization of storage")

	router := server.InitRouter(&cfg, log)
	log.Debug("Successful initialization of router")

	_ = db
	_ = router

	//TODO: init server
}
