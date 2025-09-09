package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mrpixik/sport-tracker/internal/auth/config"
	mwLogger "github.com/mrpixik/sport-tracker/internal/auth/http-server/middleware/logger"
	"log/slog"
)

// InitRouter func for creating new chi.Router with all Handlers
func InitRouter(cfg *config.Config, log *slog.Logger) chi.Router {
	router := chi.NewRouter()

	//Global middleware
	router.Use(
		middleware.RequestID,
		mwLogger.New(log),
	)

	//Public endpoints

	//Protected endpoints

	return router
}
