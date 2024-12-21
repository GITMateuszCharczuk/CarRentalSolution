package queries

import (
	get_email_query "email-service/Application/query_handlers/get_email"
	get_emails_query "email-service/Application/query_handlers/get_emails"
	fetcher "email-service/Domain/fetcher"
	service_interfaces "email-service/Domain/service_interfaces"
	"log"

	"github.com/mehdihadeli/go-mediatr"
)

func registerGetEmailQueryHandler(fetcher fetcher.DataFetcher, microserviceConnector service_interfaces.MicroserviceConnector) {
	handler := get_email_query.NewGetEmailQueryHandler(fetcher, microserviceConnector)
	err := mediatr.RegisterRequestHandler(handler)
	if err != nil {
		log.Fatal(err)
	}
}

func registerGetEmailsQueryHandler(fetcher fetcher.DataFetcher, microserviceConnector service_interfaces.MicroserviceConnector) {
	handler := get_emails_query.NewGetEmailsQueryHandler(fetcher, microserviceConnector)
	err := mediatr.RegisterRequestHandler(handler)
	if err != nil {
		log.Fatal(err)
	}
}

func RegisterQueryHandlers(fetcher fetcher.DataFetcher, microserviceConnector service_interfaces.MicroserviceConnector) {
	registerGetEmailQueryHandler(fetcher, microserviceConnector)
	registerGetEmailsQueryHandler(fetcher, microserviceConnector)
}
