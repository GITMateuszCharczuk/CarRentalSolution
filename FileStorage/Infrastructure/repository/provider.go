package repository

import (
	repository "file-storage/Domain/repository_interfaces"

	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

func ProvideFileRepository(collection *mongo.Collection) repository.FileRepository {
	return NewFileRepository(collection)
}

var WireSet = wire.NewSet(ProvideFileRepository)
