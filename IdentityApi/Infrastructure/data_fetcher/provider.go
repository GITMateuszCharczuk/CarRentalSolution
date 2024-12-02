package datafetcher

import (
	"identity-api/Domain/fetcher"
	"identity-api/Infrastructure/config"

	"github.com/google/wire"
)

func ProvideDataFetcherImpl(cfg *config.Config) fetcher.DataFetcher {
	return NewDataFetcherImpl(cfg.MailhogUrl)
}

var WireSet = wire.NewSet(ProvideDataFetcherImpl)
