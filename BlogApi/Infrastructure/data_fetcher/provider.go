package datafetcher

import (
	"identity-api/Infrastructure/config"
	interfaces "identity-api/Infrastructure/data_fetcher/interfaces"

	"github.com/google/wire"
)

func ProvideDataFetcherImpl(cfg *config.Config) interfaces.DataFetcher {
	return NewDataFetcherImpl(cfg.IdentityApiUrl)
}

var WireSet = wire.NewSet(ProvideDataFetcherImpl)
