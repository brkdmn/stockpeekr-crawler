package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DB_NAME = "stockpeekr"
)

// Connect establishes a connection to MongoDB using the provided URI.
func Connect(ctx context.Context, uri string) (*mongo.Client, error) {
	clientOpts := options.Client().ApplyURI(uri)
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, fmt.Errorf("mongo.NewClient: %w", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("client.Ping: %w", err)
	}
	return client, nil
}

func ParityCollection(client *mongo.Client) *mongo.Collection {
	return client.Database(DB_NAME).Collection("parity")
}

func ParityTickerCollection(client *mongo.Client) *mongo.Collection {
	return client.Database(DB_NAME).Collection("parityticker")
}
