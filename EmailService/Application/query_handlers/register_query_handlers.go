package queries

import (
	get_email_query "email-service/Application/query_handlers/get_email"
	get_emails_query "email-service/Application/query_handlers/get_emails"
	fetcher "email-service/Domain/fetcher"
	"log"

	"github.com/mehdihadeli/go-mediatr"
)

func registerGetEmailQueryHandler(fetcher fetcher.DataFetcher) {
	handler := get_email_query.NewGetEmailQueryHandler(fetcher)
	err := mediatr.RegisterRequestHandler(handler)
	if err != nil {
		log.Fatal(err)
	}
}

func registerGetEmailsQueryHandler(fetcher fetcher.DataFetcher) {
	handler := get_emails_query.NewGetEmailsQueryHandler(fetcher)
	err := mediatr.RegisterRequestHandler(handler)
	if err != nil {
		log.Fatal(err)
	}
}

func RegisterQueryHandlers(fetcher fetcher.DataFetcher) {
	registerGetEmailQueryHandler(fetcher)
	registerGetEmailsQueryHandler(fetcher)
}
