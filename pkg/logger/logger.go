package logger

import (
	"github.com/mrpixik/sport-tracker/pkg/logger/sl/handlers/slogpretty"
	"log/slog"
	"os"
)

func MustInit(env string) *slog.Logger {
	switch env {
	case "local":
		return slogpretty.SetupPrettySlog()
	case "dev":
		return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case "prod":
		return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	default:
		return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
}
