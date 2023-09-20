package database

import (
	"context"
	"github.com/arilsonsantos/crud-go.git/src/configuration/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

var (
	MONGODB_URL   = "MONGODB_URL"
	DATABASE_NAME = "DATABASE_NAME"
)

func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {

	mongodbUri := os.Getenv(MONGODB_URL)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodbUri))

	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	logger.Info("Banco de dados UP")
	databaseName := os.Getenv(DATABASE_NAME)

	return client.Database(databaseName), nil
}
