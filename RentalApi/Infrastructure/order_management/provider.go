package order_management

import (
	offer_repository "rental-api/Domain/repository_interfaces/car_offer_repository"
	order_repository "rental-api/Domain/repository_interfaces/car_order_repository"
	services "rental-api/Domain/service_interfaces"

	"github.com/google/wire"
)

func ProvideOrderStatusChecker(
	orderQueryRepo order_repository.CarOrderQueryRepository,
	orderCommandRepo order_repository.CarOrderCommandRepository,
	offerQueryRepo offer_repository.CarOfferQueryRepository,
	microserviceConnector services.MicroserviceConnector,
) services.OrderManagementSystem {
	return NewOrderStatusChecker(orderQueryRepo, orderCommandRepo, offerQueryRepo, microserviceConnector)
}

var WireSet = wire.NewSet(
	ProvideOrderStatusChecker,
)
