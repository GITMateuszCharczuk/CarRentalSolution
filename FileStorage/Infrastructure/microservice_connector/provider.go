package microservice_connector

import (
	interfaces "file-storage/Domain/service_interfaces"
	"file-storage/Infrastructure/config"

	"github.com/google/wire"
)

func ProvideMicroserviceConnectorImpl(cfg *config.Config) interfaces.MicroserviceConnector {
	return NewMicroserviceConnectorImpl(cfg.IdentityApiUrl)
}

var WireSet = wire.NewSet(ProvideMicroserviceConnectorImpl)
