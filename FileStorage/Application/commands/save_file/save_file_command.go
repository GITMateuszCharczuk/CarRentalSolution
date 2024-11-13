package commands

type SaveFileCommand struct {
	OwnerID  string
	FileName string
	Content  []byte
}
