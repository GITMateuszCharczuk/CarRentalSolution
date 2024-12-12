package server

import (
	"identity-api/API/controllers"
	"identity-api/Infrastructure/config"

	"github.com/google/wire"
)

func ProvideServer(Controllers *controllers.Controllers, Config *config.Config) *Server {
	return NewServer(Controllers, Config)
}

var WireSet = wire.NewSet(ProvideServer)
