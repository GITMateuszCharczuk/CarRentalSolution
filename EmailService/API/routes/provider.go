package routes

import (
	"email-service/API/controllers"
	"email-service/Infrastructure/config"

	"github.com/google/wire"
)

func ProvideRouter(Controllers *controllers.Controllers, Config *config.Config) *Router {
	return NewRouter(Controllers, Config)
}

var WireSet = wire.NewSet(ProvideRouter)
