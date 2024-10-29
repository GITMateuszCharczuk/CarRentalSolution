package handlers

import "github.com/google/wire"

type Handlers struct {
	All []Handler
}

func NewHandlers(all []Handler) *Handlers {
	return &Handlers{All: all}
}

func ProvideHandlers(
	saveFileHandler *SaveFileHandler,
	getFileHandler *GetFileHandler,
	deleteFileHandler *DeleteFileHandler,
) []Handler {
	return []Handler{saveFileHandler, getFileHandler, deleteFileHandler}
}

var WireSet = wire.NewSet(
	NewSaveFileHandler,
	NewGetFileHandler,
	NewDeleteFileHandler,
	ProvideHandlers,
	NewHandlers,
)
