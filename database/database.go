package database

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB_NAME string = os.Getenv("DB_NAME")
var DB_COLLECTION string = os.Getenv("DB_COLLECTION")
var DB_CONNECT_STRING string = os.Getenv("DB_CONNECT_STRING")

func Connect() (*mongo.Client, context.CancelFunc, error) {
	// Connect
	client, err := mongo.NewClient(options.Client().ApplyURI(DB_CONNECT_STRING))
	if err != nil {
		return nil, nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	if err := client.Connect(ctx); err != nil {
		return nil, cancel, err
	}
	return client, cancel, nil
}
