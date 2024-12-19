package queries

import (
	"log"
	// Existing imports
	get_car_images "rental-api/Application/query_handlers/car_image/get_images"
	get_car_offer "rental-api/Application/query_handlers/car_offer/get_car_offer"
	get_car_offers "rental-api/Application/query_handlers/car_offer/get_car_offers"
	get_car_order "rental-api/Application/query_handlers/car_order/get_car_order"
	get_car_orders "rental-api/Application/query_handlers/car_order/get_car_orders"
	get_car_tags "rental-api/Application/query_handlers/car_tag/get_tags"
	car_image_repository "rental-api/Domain/repository_interfaces/car_image_repository"
	car_offer_repository "rental-api/Domain/repository_interfaces/car_offer_repository"
	car_order_repository "rental-api/Domain/repository_interfaces/car_order_repository"
	car_tag_repository "rental-api/Domain/repository_interfaces/car_tag_repository"
	connector "rental-api/Domain/service_interfaces"

	"github.com/mehdihadeli/go-mediatr"
)

func registerGetCarOrdersQueryHandler(
	orderQueryRepository car_order_repository.CarOrderQueryRepository,
	connector connector.MicroserviceConnector,
) {
	handler := get_car_orders.NewGetCarOrdersQueryHandler(orderQueryRepository, connector)
	err := mediatr.RegisterRequestHandler(handler)
	if err != nil {
		log.Fatal(err)
	}
}

func registerGetCarOrderQueryHandler(
	orderQueryRepository car_order_repository.CarOrderQueryRepository,
	connector connector.MicroserviceConnector,
) {
	handler := get_car_order.NewGetCarOrderQueryHandler(orderQueryRepository, connector)
	err := mediatr.RegisterRequestHandler(handler)
	if err != nil {
		log.Fatal(err)
	}
}

func registerGetCarOffersQueryHandler(
	offerQueryRepository car_offer_repository.CarOfferQueryRepository,
	connector connector.MicroserviceConnector,
) {
	handler := get_car_offers.NewGetCarOffersQueryHandler(offerQueryRepository, connector)
	err := mediatr.RegisterRequestHandler(handler)
	if err != nil {
		log.Fatal(err)
	}
}

func registerGetCarOfferQueryHandler(
	offerQueryRepository car_offer_repository.CarOfferQueryRepository,
) {
	handler := get_car_offer.NewGetCarOfferQueryHandler(offerQueryRepository)
	err := mediatr.RegisterRequestHandler(handler)
	if err != nil {
		log.Fatal(err)
	}
}

func registerGetCarTagsQueryHandler(
	tagQueryRepository car_tag_repository.CarTagQueryRepository,
) {
	handler := get_car_tags.NewGetTagsQueryHandler(tagQueryRepository)
	err := mediatr.RegisterRequestHandler(handler)
	if err != nil {
		log.Fatal(err)
	}
}
func registerGetCarImagesQueryHandler(
	imageQueryRepository car_image_repository.CarImageQueryRepository,
) {
	handler := get_car_images.NewGetImagesQueryHandler(imageQueryRepository)
	err := mediatr.RegisterRequestHandler(handler)
	if err != nil {
		log.Fatal(err)
	}
}

func RegisterQueryHandlers(
	carOrderQueryRepository car_order_repository.CarOrderQueryRepository,
	carOfferQueryRepository car_offer_repository.CarOfferQueryRepository,
	connector connector.MicroserviceConnector,
	carTagQueryRepository car_tag_repository.CarTagQueryRepository,
	carImageQueryRepository car_image_repository.CarImageQueryRepository,
) {
	registerGetCarOrdersQueryHandler(carOrderQueryRepository, connector)
	registerGetCarOrderQueryHandler(carOrderQueryRepository, connector)
	registerGetCarOffersQueryHandler(carOfferQueryRepository, connector)
	registerGetCarOfferQueryHandler(carOfferQueryRepository)
	registerGetCarTagsQueryHandler(carTagQueryRepository)
	registerGetCarImagesQueryHandler(carImageQueryRepository)
}
