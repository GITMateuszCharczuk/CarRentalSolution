package db

import (
    "context"
    "fmt"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "log"
)

var client *mongo.Client
var filesCollection *mongo.Collection

func ConnectMongo() {
    var err error
    client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
    if err != nil {
        log.Fatalf("Failed to connect to MongoDB: %v", err)
    }
    filesCollection = client.Database("filedb").Collection("files")
}

func InsertFile(ctx context.Context, fileID, ownerID, filePath string) error {
    _, err := filesCollection.InsertOne(ctx, map[string]interface{}{
        "fileID":   fileID,
        "ownerID":  ownerID,
        "filePath": filePath,
    })
    return err
}

func GetFileByID(ctx context.Context, fileID string) (string, error) {
    var result map[string]interface{}
    err := filesCollection.FindOne(ctx, map[string]interface{}{"fileID": fileID}).Decode(&result)
    if err != nil {
        return "", fmt.Errorf("file not found")
    }
    return result["filePath"].(string), nil
}

func DeleteFileByID(ctx context.Context, fileID string) error {
    _, err := filesCollection.DeleteOne(ctx, map[string]interface{}{"fileID": fileID})
    return err
}
