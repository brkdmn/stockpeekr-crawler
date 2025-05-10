package main

import (
	"context"
	"log"
	"os"
	"stockpeekr-crawler/internal/config"
	"stockpeekr-crawler/internal/db"
	"stockpeekr-crawler/internal/logger"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	logg, err := logger.New(cfg.SentryDSN)
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logg.Sync()
	logg.Info("Logger initialized")

	ctx := context.Background()
	client, err := db.Connect(ctx, cfg.MongoURI)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
		os.Exit(1)
	}
	defer client.Disconnect(ctx)
	logg.Info("MongoDB connected")

}
