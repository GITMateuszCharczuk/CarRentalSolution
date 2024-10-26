package repository

import (
	"context"
	"fmt"

	"file-storage/Domain/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type FileRepositoryImpl struct {
	collection *mongo.Collection
}

func NewFileRepository(collection *mongo.Collection) *FileRepositoryImpl {
	return &FileRepositoryImpl{collection: collection}
}

func (r *FileRepositoryImpl) InsertFile(ctx context.Context, file models.File) error {
	_, err := r.collection.InsertOne(ctx, file)
	return err
}

func (r *FileRepositoryImpl) GetFileByID(ctx context.Context, fileID string) (models.File, error) {
	var result models.File
	err := r.collection.FindOne(ctx, map[string]interface{}{"id": fileID}).Decode(&result)
	if err != nil {
		return models.File{}, fmt.Errorf("file not found")
	}
	return result, nil
}

func (r *FileRepositoryImpl) DeleteFileByID(ctx context.Context, fileID string) error {
	_, err := r.collection.DeleteOne(ctx, map[string]interface{}{"id": fileID})
	return err
}
