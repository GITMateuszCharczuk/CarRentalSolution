package repository

import (
	repository "file-storage/Domain/repository_interfaces"

	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

func ProvideFileRepository(collection *mongo.Collection, bucket *gridfs.Bucket) repository.FileRepository {
	return NewFileRepository(collection, bucket)
}

var WireSet = wire.NewSet(ProvideFileRepository)
