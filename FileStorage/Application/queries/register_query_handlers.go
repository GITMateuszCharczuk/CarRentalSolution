package queries

import (
	get_query "file-storage/Application/queries/get_file"
	"file-storage/Domain/repository_interfaces"
	"log"

	"github.com/mehdihadeli/go-mediatr"
)

func registerGetFileQueryHandler(fileRepo repository_interfaces.FileRepository) {
	handler := get_query.NewGetFileQueryHandler(fileRepo)
	err := mediatr.RegisterRequestHandler(handler)
	if err != nil {
		log.Fatal(err)
	}
}

func RegisterQueryHandlers(fileRepo repository_interfaces.FileRepository) {
	registerGetFileQueryHandler(fileRepo)
}
