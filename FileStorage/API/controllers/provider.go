package controllers

import "github.com/google/wire"

type Controllers struct {
	All []Controller
}

func NewControllers(all []Controller) *Controllers {
	return &Controllers{All: all}
}

func ProvideControllers(
	saveFileHandler *SaveFileController,
	getFileHandler *GetFileController,
	deleteFileHandler *DeleteFileController,
) []Controller {
	return []Controller{saveFileHandler, getFileHandler, deleteFileHandler}
}

var WireSet = wire.NewSet(
	NewSaveFileController,
	NewGetFileController,
	NewDeleteFileController,
	ProvideControllers,
	NewControllers,
)
