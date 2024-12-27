package models

import "time"

type FileModel struct {
	ID          string
	Name        string
	Size        int64
	MimeType    string
	Purpose     string
	StoragePath string
	CreatedAt   time.Time
}
