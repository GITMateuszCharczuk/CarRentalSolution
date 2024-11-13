package db

import (
	"context"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	clientInstance   *mongo.Client
	dataBaseInstance *mongo.Database
	clientOnce       sync.Once
)

func InitializeDB(mongoURI string, dbName string) *mongo.Database {
	var err error
	clientOnce.Do(func() {
		clientInstance, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
		if err != nil {
			log.Fatalf("Failed to connect to MongoDB: %v", err)
		}
		dataBaseInstance = clientInstance.Database(dbName)
	})
	return dataBaseInstance
}
