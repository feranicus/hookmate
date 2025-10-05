package main

import (
	"log"

	"github.com/feranicus/hookmate/internal/api"
	"github.com/feranicus/hookmate/internal/config"
	"go.uber.org/zap"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	// Initialize logger
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	sugar := logger.Sugar()

	// Setup and run the server
	sugar.Infof("starting server on port %s", cfg.Server.Port)
	if err := api.Run(cfg, sugar); err != nil {
		sugar.Fatalf("could not start server: %v", err)
	}
}