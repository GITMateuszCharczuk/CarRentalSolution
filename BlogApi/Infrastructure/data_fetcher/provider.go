package datafetcher

import (
	interfaces "identity-api/Domain/service_interfaces"
	"identity-api/Infrastructure/config"

	"github.com/google/wire"
)

func ProvideDataFetcherImpl(cfg *config.Config) interfaces.DataFetcher {
	return NewDataFetcherImpl(cfg.IdentityApiUrl)
}

var WireSet = wire.NewSet(ProvideDataFetcherImpl)
