package commands

type SaveFileCommand struct {
	FileID   string
	OwnerID  string
	FileName string
	Content  []byte
}
