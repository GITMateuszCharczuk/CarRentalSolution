package models

type File struct {
	ID       string
	OwnerID  string
	FileName string
	Content  []byte
}
