package datafetcher

import (
	interfaces "blog-api/Domain/service_interfaces"
	"blog-api/Infrastructure/config"

	"github.com/google/wire"
)

func ProvideDataFetcherImpl(cfg *config.Config) interfaces.DataFetcher {
	return NewDataFetcherImpl(cfg.IdentityApiUrl)
}

var WireSet = wire.NewSet(ProvideDataFetcherImpl)
