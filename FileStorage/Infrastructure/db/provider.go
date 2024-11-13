package db

import (
	"file-storage/Infrastructure/config"

	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

func ProvideMongoDB(cfg *config.Config) *mongo.Database {
	return InitializeDB(cfg.MongoDBUrl, cfg.MongoDBName)
}

func ProvideMongoCollection(db *mongo.Database, cfg *config.Config) *mongo.Collection {
	return InitializeFilesCollection(db, cfg.MongoDBCollName)
}

func ProvideBucket(db *mongo.Database) *gridfs.Bucket {
	return InitializeBucket(db)
}

var WireSet = wire.NewSet(ProvideMongoDB, ProvideMongoCollection, ProvideBucket)
