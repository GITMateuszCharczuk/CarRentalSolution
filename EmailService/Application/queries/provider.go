package queries

import (
	get_email_query "email-service/Application/queries/get_email"
	get_emails_query "email-service/Application/queries/get_emails"
	datafetcher "email-service/Infrastructure/data_fetcher"

	"github.com/google/wire"
)

func ProvideGetEmailQueryHandler(fetcher datafetcher.DataFetcher) *get_email_query.GetEmailQueryHandler {
	return get_email_query.NewGetEmailQueryHandler(fetcher)
}

func ProvideGetEmailsQueryHandler(fetcher datafetcher.DataFetcher) *get_emails_query.GetEmailsQueryHandler {
	return get_emails_query.NewGetEmailsQueryHandler(fetcher)
}

// type QueryHandlers struct {
// 	GetEmailsQuery *get_emails_query.GetEmailsQuery
// 	GetEmailQuery  *get_email_query.GetEmailQuery
// }

var WireSet = wire.NewSet(
	ProvideGetEmailQueryHandler,
	ProvideGetEmailsQueryHandler,
	// wire.Struct(new(QueryHandlers), "*"),
)
