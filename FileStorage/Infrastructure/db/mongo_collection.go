package db

import (
	"context"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	collectionInstance *mongo.Collection
	collectionOnce     sync.Once
)

func InitializeFilesCollection(db *mongo.Database, collName string) *mongo.Collection {
	collectionOnce.Do(func() {
		collectionInstance = db.Collection(collName)

		collections, err := db.ListCollectionNames(context.TODO(), map[string]interface{}{"name": collName})
		if err != nil {
			log.Fatalf("Failed to list collections: %v", err)
		}

		if len(collections) == 0 {
			if err := db.CreateCollection(context.TODO(), collName); err != nil {
				log.Fatalf("Failed to create collection: %v", err)
			}
			log.Println("Collection created:", collName)
		} else {
			log.Println("Collection already exists:", collName)
		}
	})
	return collectionInstance
}
