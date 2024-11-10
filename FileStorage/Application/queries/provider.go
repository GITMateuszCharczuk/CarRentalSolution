package queries

import (
	get_query "file-storage/Application/queries/get_file"
	"file-storage/Domain/repository_interfaces"

	"github.com/google/wire"
)

func ProvideGetFileQueryHandler(fileRepo repository_interfaces.FileRepository) *get_query.GetFileQueryHandler {
	return get_query.NewGetFileQueryHandler(fileRepo)
}

type QueryHandlers struct {
	GetFileQuery *get_query.GetFileQuery
}

var WireSet = wire.NewSet(
	ProvideGetFileQueryHandler,
	wire.Struct(new(QueryHandlers), "*"),
)
