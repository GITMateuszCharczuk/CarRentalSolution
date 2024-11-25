package server

import (
	"file-storage/API/controllers"
	"file-storage/Infrastructure/config"
	"github.com/google/wire"
)

func ProvideRoutes(controllers *controllers.Controllers, config *config.Config) *Server {
	return NewServer(controllers, config)
}

var WireSet = wire.NewSet(ProvideRoutes)
