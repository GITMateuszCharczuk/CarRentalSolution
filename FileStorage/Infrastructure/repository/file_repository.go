package repository

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"file-storage/Domain/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

type FileRepositoryImpl struct {
	collection *mongo.Collection
	bucket     *gridfs.Bucket
}

func NewFileRepository(collection *mongo.Collection, bucket *gridfs.Bucket) *FileRepositoryImpl {
	return &FileRepositoryImpl{collection: collection, bucket: bucket}
}

func (r *FileRepositoryImpl) InsertFile(ctx context.Context, file models.File) error {
	contentReader := bytes.NewReader(file.Content)

	gridFSUploadStream, err := r.bucket.OpenUploadStream(file.ID)
	if err != nil {
		return err
	}
	defer gridFSUploadStream.Close()

	if _, err := io.Copy(gridFSUploadStream, contentReader); err != nil {
		return err
	}

	metadata := bson.M{
		"id":        file.ID,
		"owner_id":  file.OwnerID,
		"file_name": file.FileName,
	}

	_, err = r.collection.InsertOne(ctx, metadata)
	return err
}

func (r *FileRepositoryImpl) GetFileByID(ctx context.Context, fileID string) (models.FileStream, error) {
	var metadata struct {
		ID       string `bson:"id"`
		OwnerID  string `bson:"owner_id"`
		FileName string `bson:"file_name"`
	}

	if err := r.collection.FindOne(ctx, bson.M{"id": fileID}).Decode(&metadata); err != nil {
		return models.FileStream{}, fmt.Errorf("file not found: %w", err)
	}

	gridFSDownloadStream, err := r.bucket.OpenDownloadStreamByName(fileID)
	if err != nil {
		return models.FileStream{}, fmt.Errorf("error downloading file content: %w", err)
	}

	fileSize := gridFSDownloadStream.GetFile().Length

	return models.FileStream{
		OwnerID:  metadata.OwnerID,
		FileName: metadata.FileName,
		FileSize: fileSize,
		Stream:   gridFSDownloadStream,
	}, nil
}

func (r *FileRepositoryImpl) DeleteFileByID(ctx context.Context, fileID string) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"id": fileID})
	if err != nil {
		return err
	}

	file, err := r.bucket.Find(bson.M{"filename": fileID})
	if err != nil {
		return err
	}

	if file.Next(ctx) {
		fileID := file.Current.Lookup("_id").ObjectID()
		if err := r.bucket.Delete(fileID); err != nil {
			return fmt.Errorf("failed to delete GridFS file: %w", err)
		}
	}

	return nil
}
