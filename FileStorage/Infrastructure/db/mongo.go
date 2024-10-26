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

func ConnectMongo() error {
	var err error
	clientOnce.Do(func() {
		clientInstance, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
		if err != nil {
			log.Fatalf("Failed to connect to MongoDB: %v", err)
		}
	})
	return err
}

func GetMongoClient() *mongo.Client {
	return clientInstance
}

func GetFilesCollection() *mongo.Collection {
	collectionOnce.Do(func() {
		if clientInstance == nil {
			err := ConnectMongo()
			if err != nil {
				log.Fatalf("Failed to initialize MongoDB client: %v", err)
			}
		}
		collectionInstance = clientInstance.Database("filedb").Collection("files")
	})
	return collectionInstance
}
