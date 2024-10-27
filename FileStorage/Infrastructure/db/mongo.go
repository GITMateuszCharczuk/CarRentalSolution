package db

import (
	"context"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	clientInstance     *mongo.Client
	collectionInstance *mongo.Collection
	clientOnce         sync.Once
	collectionOnce     sync.Once
)

func ConnectMongo(mongoURI string) error {
	var err error
	clientOnce.Do(func() {
		clientInstance, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
		if err != nil {
			log.Fatalf("Failed to connect to MongoDB: %v", err)
		}
	})
	return err
}

func GetFilesCollection(mongoURI string, dbName string, collName string) *mongo.Collection {
	collectionOnce.Do(func() {
		if clientInstance == nil {
			err := ConnectMongo(mongoURI)
			if err != nil {
				log.Fatalf("Failed to initialize MongoDB client: %v", err)
			}
		}
		collectionInstance = clientInstance.Database(dbName).Collection(collName)
	})
	return collectionInstance
}
