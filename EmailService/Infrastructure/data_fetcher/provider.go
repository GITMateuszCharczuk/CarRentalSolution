package datafetcher

import (
	"email-service/Infrastructure/config"

	"github.com/google/wire"
)

func ProvideDataFetcherImpl(cfg *config.Config) DataFetcher {
	return NewDataFetcherImpl(cfg.MailhogUrl)
}

var WireSet = wire.NewSet(ProvideDataFetcherImpl)
