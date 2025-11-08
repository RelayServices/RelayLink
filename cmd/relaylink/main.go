package main

import (
	"log/slog"
	"os"

	"github.com/RelayServices/RelayLink/internal/config"
	"github.com/RelayServices/RelayLink/internal/lib/logger/sl"
	"github.com/RelayServices/RelayLink/internal/storage/sqlite"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("starting RelayLink", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	_ = storage

	// TODO: init router: chi, chi render

	// TODO: run server
}

func setupLogger(env string) *slog.Logger {
	var handler slog.Handler

	switch env {
	case envLocal:
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	case envDev:
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	case envProd:
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	default:
		panic("unknown environment: " + env) 
	}

	return slog.New(handler)
}
