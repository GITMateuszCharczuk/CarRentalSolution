package datafetcher

import (
	"email-service/Domain/fetcher"
	"email-service/Infrastructure/config"

	"github.com/google/wire"
)

func ProvideDataFetcherImpl(cfg *config.Config) fetcher.DataFetcher {
	return NewDataFetcherImpl(cfg.MailhogUrl)
}

var WireSet = wire.NewSet(ProvideDataFetcherImpl)
