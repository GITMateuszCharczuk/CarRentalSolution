package repository

import (
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

func ProvideFileRepository(collection *mongo.Collection) FileRepository {
	return NewFileRepository(collection)
}

var WireSet = wire.NewSet(ProvideFileRepository)
