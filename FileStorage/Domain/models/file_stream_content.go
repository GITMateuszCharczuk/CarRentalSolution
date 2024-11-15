package models

import "go.mongodb.org/mongo-driver/mongo/gridfs"

type FileStream struct {
	OwnerID  string
	FileName string
	FileSize int64
	Stream   *gridfs.DownloadStream
}
