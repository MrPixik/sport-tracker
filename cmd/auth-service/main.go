package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mrpixik/sport-tracker/internal/auth/config"
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

	err = db.DeleteUser(context.Background(), "pixik", "123")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = db.CreateUser(context.Background(), "pixik", "123")
	if err != nil {
		fmt.Println(err.Error())
	}

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)

	//TODO: init server
}
