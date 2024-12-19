package microservice_connector

import (
	interfaces "rental-api/Domain/service_interfaces"
	"rental-api/Infrastructure/config"

	"github.com/google/wire"
)

func ProvideMicroserviceConnectorImpl(cfg *config.Config) interfaces.MicroserviceConnector {
	return NewMicroserviceConnectorImpl(cfg.IdentityApiUrl, cfg.EmailServiceBaseUrl, cfg.CompanyEmail)
}

var WireSet = wire.NewSet(ProvideMicroserviceConnectorImpl)
