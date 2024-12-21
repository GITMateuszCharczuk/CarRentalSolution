package microservice_connector

import (
	interfaces "email-service/Domain/service_interfaces"
	"email-service/Infrastructure/config"

	"github.com/google/wire"
)

func ProvideMicroserviceConnectorImpl(cfg *config.Config) interfaces.MicroserviceConnector {
	return NewMicroserviceConnectorImpl(cfg.IdentityApiUrl)
}

var WireSet = wire.NewSet(ProvideMicroserviceConnectorImpl)
