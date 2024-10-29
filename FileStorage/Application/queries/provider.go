package queries

import (
	"file-storage/Domain/repository_interfaces"

	"github.com/google/wire"
)

func ProvideGetFileQuery(fileRepo repository_interfaces.FileRepository) *GetFileQuery {
	return NewGetFileQuery(fileRepo)
}

type Queries struct {
	GetFileQuery *GetFileQuery
}

var WireSet = wire.NewSet(
	ProvideGetFileQuery,
	wire.Struct(new(Queries), "*"),
)
