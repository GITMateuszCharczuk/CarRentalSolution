// main/wire.go

// go:build wireinject
//go:build wireinject
// +build wireinject

package main

import (
	"rental-api/API/controllers"
	"rental-api/API/server"
	validators "rental-api/API/validators"
	car_image_repository "rental-api/Domain/repository_interfaces/car_image_repository"
	car_offer_repository "rental-api/Domain/repository_interfaces/car_offer_repository"
	car_order_repository "rental-api/Domain/repository_interfaces/car_order_repository"
	car_tag_repository "rental-api/Domain/repository_interfaces/car_tag_repository"
	service_interfaces "rental-api/Domain/service_interfaces"
	config "rental-api/Infrastructure/config"
	postgres_db "rental-api/Infrastructure/databases/postgres/config"
	mappers "rental-api/Infrastructure/databases/postgres/mappers"
	unit_of_work "rental-api/Infrastructure/databases/postgres/repository/base/unit_of_work"
	car_image_repository_impl "rental-api/Infrastructure/databases/postgres/repository/car_image_repository"
	car_offer_repository_impl "rental-api/Infrastructure/databases/postgres/repository/car_offer_repository"
	car_order_repository_impl "rental-api/Infrastructure/databases/postgres/repository/car_order_repository"
	car_tag_repository_impl "rental-api/Infrastructure/databases/postgres/repository/car_tag_repository"
	connector "rental-api/Infrastructure/microservice_connector"
	order_management "rental-api/Infrastructure/order_management"

	"github.com/google/wire"
)

type InfrastructureComponents struct {
	Config                *config.Config
	CarOfferQueryRepo     car_offer_repository.CarOfferQueryRepository
	CarOfferCommandRepo   car_offer_repository.CarOfferCommandRepository
	CarOrderQueryRepo     car_order_repository.CarOrderQueryRepository
	CarOrderCommandRepo   car_order_repository.CarOrderCommandRepository
	CarImageQueryRepo     car_image_repository.CarImageQueryRepository
	CarImageCommandRepo   car_image_repository.CarImageCommandRepository
	CarTagQueryRepo       car_tag_repository.CarTagQueryRepository
	CarTagCommandRepo     car_tag_repository.CarTagCommandRepository
	connector             service_interfaces.MicroserviceConnector
	orderManagementSystem service_interfaces.OrderManagementSystem
}

func InitializeInfrastructureComponents() (*InfrastructureComponents, error) {
	wire.Build(
		// Config
		config.WireSet,
		// Database
		unit_of_work.ProvideUnitOfWork,
		postgres_db.WireSet,
		// Repositories
		car_offer_repository_impl.WireSet,
		car_order_repository_impl.WireSet,
		car_image_repository_impl.WireSet,
		car_tag_repository_impl.WireSet,
		// Mappers
		mappers.WireSet,
		// Services
		connector.WireSet,
		order_management.WireSet,
		wire.Struct(new(InfrastructureComponents), "*"),
	)
	return &InfrastructureComponents{}, nil
}

func InitializeApi(
	carOfferQueryRepo car_offer_repository.CarOfferQueryRepository,
	carOfferCommandRepo car_offer_repository.CarOfferCommandRepository,
	carOrderQueryRepo car_order_repository.CarOrderQueryRepository,
	carOrderCommandRepo car_order_repository.CarOrderCommandRepository,
	carImageQueryRepo car_image_repository.CarImageQueryRepository,
	carImageCommandRepo car_image_repository.CarImageCommandRepository,
	carTagQueryRepo car_tag_repository.CarTagQueryRepository,
	carTagCommandRepo car_tag_repository.CarTagCommandRepository,
	connector service_interfaces.MicroserviceConnector,
	orderManagementSystem service_interfaces.OrderManagementSystem,
	config *config.Config,
) (*server.Server, error) {
	wire.Build(
		validators.WireSet,
		controllers.WireSet,
		server.WireSet,
	)
	return &server.Server{}, nil
}
