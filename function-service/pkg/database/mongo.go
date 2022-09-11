package database

import (
	"context"
	"function-srv/pkg/messages"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func New(dbURI, dbName string, logger *log.Logger) (*mongo.Database, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbURI))
	if err != nil {
		logger.Fatalf("%s %s", messages.Error(), err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		logger.Fatalf("%s %s", messages.Error(), err)
	}

	return client.Database(dbName), cancel
}
