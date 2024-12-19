package commands

import (
	"log"
	add_image "rental-api/Application/command_handlers/car_image/add_image"
	delete_image "rental-api/Application/command_handlers/car_image/delete_image"
	create_car_offer "rental-api/Application/command_handlers/car_offer/create_car_offer"
	delete_car_offer "rental-api/Application/command_handlers/car_offer/delete_car_offer"
	update_car_offer "rental-api/Application/command_handlers/car_offer/update_car_offer"
	create_car_order "rental-api/Application/command_handlers/car_order/create_car_order"
	delete_car_order "rental-api/Application/command_handlers/car_order/delete_car_order"
	update_car_order "rental-api/Application/command_handlers/car_order/update_car_order"
	car_image_repository "rental-api/Domain/repository_interfaces/car_image_repository"
	car_offer_repository "rental-api/Domain/repository_interfaces/car_offer_repository"
	car_order_repository "rental-api/Domain/repository_interfaces/car_order_repository"
	connector "rental-api/Domain/service_interfaces"

	"github.com/mehdihadeli/go-mediatr"
)

func registerCreateCarOrderCommandHandler(
	orderCommandRepository car_order_repository.CarOrderCommandRepository,
	offerQueryRepository car_offer_repository.CarOfferQueryRepository,
	connector connector.MicroserviceConnector,
) {
	handler := create_car_order.NewCreateCarOrderCommandHandler(
		orderCommandRepository,
		offerQueryRepository,
		connector,
	)
	err := mediatr.RegisterRequestHandler(handler)
	if err != nil {
		log.Fatal(err)
	}
}

func registerDeleteCarOrderCommandHandler(
	orderCommandRepository car_order_repository.CarOrderCommandRepository,
	orderQueryRepository car_order_repository.CarOrderQueryRepository,
	connector connector.MicroserviceConnector,
) {
	handler := delete_car_order.NewDeleteCarOrderCommandHandler(
		orderCommandRepository,
		orderQueryRepository,
		connector,
	)
	err := mediatr.RegisterRequestHandler(handler)
	if err != nil {
		log.Fatal(err)
	}
}

func registerUpdateCarOrderCommandHandler(
	orderCommandRepository car_order_repository.CarOrderCommandRepository,
	orderQueryRepository car_order_repository.CarOrderQueryRepository,
	offerQueryRepository car_offer_repository.CarOfferQueryRepository,
	connector connector.MicroserviceConnector,
) {
	handler := update_car_order.NewUpdateCarOrderCommandHandler(
		orderCommandRepository,
		orderQueryRepository,
		offerQueryRepository,
		connector,
	)
	err := mediatr.RegisterRequestHandler(handler)
	if err != nil {
		log.Fatal(err)
	}
}

func registerCreateCarOfferCommandHandler(
	offerCommandRepository car_offer_repository.CarOfferCommandRepository,
	connector connector.MicroserviceConnector,
) {
	handler := create_car_offer.NewCreateCarOfferCommandHandler(
		offerCommandRepository,
		connector,
	)
	err := mediatr.RegisterRequestHandler(handler)
	if err != nil {
		log.Fatal(err)
	}
}

func registerUpdateCarOfferCommandHandler(
	offerCommandRepository car_offer_repository.CarOfferCommandRepository,
	offerQueryRepository car_offer_repository.CarOfferQueryRepository,
	connector connector.MicroserviceConnector,
) {
	handler := update_car_offer.NewUpdateCarOfferCommandHandler(
		offerCommandRepository,
		offerQueryRepository,
		connector,
	)
	err := mediatr.RegisterRequestHandler(handler)
	if err != nil {
		log.Fatal(err)
	}
}

func registerDeleteCarOfferCommandHandler(
	offerCommandRepository car_offer_repository.CarOfferCommandRepository,
	offerQueryRepository car_offer_repository.CarOfferQueryRepository,
	connector connector.MicroserviceConnector,
) {
	handler := delete_car_offer.NewDeleteCarOfferCommandHandler(
		offerCommandRepository,
		offerQueryRepository,
		connector,
	)
	err := mediatr.RegisterRequestHandler(handler)
	if err != nil {
		log.Fatal(err)
	}
}

func registerAddImageCommandHandler(
	imageCommandRepository car_image_repository.CarImageCommandRepository,
	offerQueryRepository car_offer_repository.CarOfferQueryRepository,
	connector connector.MicroserviceConnector,
) {
	handler := add_image.NewAddImageCommandHandler(
		imageCommandRepository,
		offerQueryRepository,
		connector,
	)
	err := mediatr.RegisterRequestHandler(handler)
	if err != nil {
		log.Fatal(err)
	}
}

func registerDeleteImageCommandHandler(
	imageCommandRepository car_image_repository.CarImageCommandRepository,
	offerQueryRepository car_offer_repository.CarOfferQueryRepository,
	connector connector.MicroserviceConnector,
) {
	handler := delete_image.NewDeleteImageCommandHandler(
		imageCommandRepository,
		offerQueryRepository,
		connector,
	)
	err := mediatr.RegisterRequestHandler(handler)
	if err != nil {
		log.Fatal(err)
	}
}

func RegisterCommandHandlers(
	carOrderCommandRepository car_order_repository.CarOrderCommandRepository,
	carOrderQueryRepository car_order_repository.CarOrderQueryRepository,
	carOfferCommandRepository car_offer_repository.CarOfferCommandRepository,
	carOfferQueryRepository car_offer_repository.CarOfferQueryRepository,
	carImageCommandRepository car_image_repository.CarImageCommandRepository,
	connector connector.MicroserviceConnector,
) {
	registerCreateCarOrderCommandHandler(carOrderCommandRepository, carOfferQueryRepository, connector)
	registerUpdateCarOrderCommandHandler(carOrderCommandRepository, carOrderQueryRepository, carOfferQueryRepository, connector)
	registerCreateCarOfferCommandHandler(carOfferCommandRepository, connector)
	registerUpdateCarOfferCommandHandler(carOfferCommandRepository, carOfferQueryRepository, connector)
	registerAddImageCommandHandler(carImageCommandRepository, carOfferQueryRepository, connector)
	registerDeleteImageCommandHandler(carImageCommandRepository, carOfferQueryRepository, connector)
	registerDeleteCarOrderCommandHandler(carOrderCommandRepository, carOrderQueryRepository, connector)
	registerDeleteCarOfferCommandHandler(carOfferCommandRepository, carOfferQueryRepository, connector)
}
