// queries/queries.go

package queries

import (
	"file-storage/Domain/repository"

	"github.com/google/wire"
)

type Queries struct {
	GetFile *GetFileQuery
}

func ProvideQueries(fileRepo repository.FileRepository) *Queries {
	return &Queries{
		GetFile: NewGetFileQuery(fileRepo),
	}
}

var WireSet = wire.NewSet(ProvideQueries)
