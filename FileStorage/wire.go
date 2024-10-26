// wireinject
//go:build wireinject
// +build wireinject

package p // wire.go

import (
	"file-storage/Application/commands"
	"file-storage/Application/queries"
	"file-storage/infrastructure"

	"github.com/google/wire"
)

func InitializeSaveFileCommand() *commands.SaveFileCommand {
	wire.Build(infrastructure.WireSet)
	return &commands.SaveFileCommand{}
}

func InitializeDeleteFileCommand() *commands.DeleteFileCommand {
	wire.Build(infrastructure.WireSet)
	return &commands.DeleteFileCommand{}
}

func InitializeGetFileQuery() *queries.GetFileQuery {
	wire.Build(infrastructure.WireSet)
	return &queries.GetFileQuery{}
}
