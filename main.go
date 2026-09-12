package main

import (
	"fmt"
	"os"

	"github.com/klpjunki/auth/internal/config"
	"github.com/klpjunki/auth/internal/logger"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "load config: %v\n", err)
		os.Exit(1)
	}

	log, err := logger.New(cfg.Log.Level, cfg.Log.Format)
	if err != nil {
		fmt.Fprintf(os.Stderr, "create logger: %v\n", err)
		os.Exit(1)
	}

	log.Info().
		Str("environment", cfg.App.Environment).
		Msg("auth service initialized")
}
