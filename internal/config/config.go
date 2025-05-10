package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	TwelveDataAPIKey string
	MongoURI         string
	SentryDSN        string
	AdminUIPort      int
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	apiKey := os.Getenv("TWELVEDATA_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("environment variable TWELVEDATA_API_KEY is required")
	}

	mongo := os.Getenv("MONGODB_URI")
	if mongo == "" {
		return nil, fmt.Errorf("environment variable MONGODB_URI is required")
	}

	sentry := os.Getenv("SENTRY_DSN")

	portStr := os.Getenv("ADMIN_UI_PORT")
	if portStr == "" {
		portStr = "8080"
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, fmt.Errorf("invalid ADMIN_UI_PORT: %w", err)
	}

	return &Config{
		TwelveDataAPIKey: apiKey,
		MongoURI:         mongo,
		SentryDSN:        sentry,
		AdminUIPort:      port,
	}, nil
}
