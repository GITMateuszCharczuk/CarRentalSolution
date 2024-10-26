package infrastructure

import (
	"file-storage/Application/commands"
	"file-storage/Application/queries"
	"file-storage/Domain/event"
	"file-storage/Domain/repository"
	"file-storage/Infrastructure/db"
	"file-storage/Infrastructure/queue"

	"github.com/google/wire"
)

func ProvideFileRepository() repository.FileRepository {
	files := db.GetFilesCollection()
	return repository.NewFileRepository(files)
}

func ProvideEventPublisher() event.EventPublisher {
	return queue.NewJetStreamPublisher()
}

func ProvideSaveFileCommand(fileRepo repository.FileRepository, eventPublisher event.EventPublisher) *commands.SaveFileCommand {
	return commands.NewSaveFileCommand(fileRepo, eventPublisher)
}

func ProvideDeleteFileCommand(fileRepo repository.FileRepository) *commands.DeleteFileCommand {
	return commands.NewDeleteFileCommand(fileRepo)
}

func ProvideGetFileQuery(fileRepo repository.FileRepository) *queries.GetFileQuery {
	return queries.NewGetFileQuery(fileRepo)
}

var WireSet = wire.NewSet(
	ProvideFileRepository,
	ProvideEventPublisher,
	ProvideSaveFileCommand,
	ProvideDeleteFileCommand,
	ProvideGetFileQuery,
)
