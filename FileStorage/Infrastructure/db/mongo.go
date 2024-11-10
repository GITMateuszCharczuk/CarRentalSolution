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

		db := clientInstance.Database(dbName)
		collectionInstance = db.Collection(collName)

		collections, err := db.ListCollectionNames(context.TODO(), map[string]interface{}{"name": collName})
		if err != nil {
			log.Fatalf("Failed to list collections: %v", err)
		}

		if len(collections) == 0 {
			if err := db.CreateCollection(context.TODO(), collName); err != nil {
				log.Fatalf("Failed to create collection: %v", err)
			}
			log.Println("Collection created: %a", collName)
		}
		log.Println("Collection already exists:", collName)
	})
	return collectionInstance
}
