package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DB_HOST string = "mongodb://localhost:27017"
const DB_NAME string = "admin"
const DB_COLLECTION string = "my-unsplash"

func Connect() (*mongo.Client, context.CancelFunc, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(DB_HOST))
	if err != nil {
		return nil, nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	if err := client.Connect(ctx); err != nil {
		return nil, cancel, err
	}
	return client, cancel, nil
}
