package db

import (
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

var (
	bucketInstance *gridfs.Bucket
	bucketOnce     sync.Once
)

func InitializeBucket(db *mongo.Database) *gridfs.Bucket {
	var err error
	bucketOnce.Do(func() {
		bucketInstance, err = gridfs.NewBucket(db)
		if err != nil {
			log.Fatalf("Failed to create gridfs bucket: %v", err)
		}
	})
	return bucketInstance
}
