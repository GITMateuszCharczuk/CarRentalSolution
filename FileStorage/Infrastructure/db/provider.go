package db

import (
	"file-storage/Infrastructure/config"

	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

func ProvideMongoCollection(cfg *config.Config) (*mongo.Collection, error) {
	col := GetFilesCollection(cfg.MongoDBUrl, cfg.MongoDBName, cfg.MongoDBCollName)
	return col, nil
}

var WireSet = wire.NewSet(ProvideMongoCollection)
