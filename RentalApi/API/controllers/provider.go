package controllers

import (
	base "rental-api/API/controllers/base"
	car_image "rental-api/API/controllers/car_image"
	car_offer "rental-api/API/controllers/car_offer"
	car_order "rental-api/API/controllers/car_order"
	car_tag "rental-api/API/controllers/car_tag"

	"github.com/google/wire"
)

var ControllerSet = wire.NewSet(
	NewControllers,
	ProvideControllers,
	// Car Offer Controllers
	car_offer.NewCreateCarOfferController,
	car_offer.NewUpdateCarOfferController,
	car_offer.NewDeleteCarOfferController,
	car_offer.NewGetCarOfferController,
	car_offer.NewGetCarOffersController,
	// Car Order Controllers
	car_order.NewCreateCarOrderController,
	car_order.NewUpdateCarOrderController,
	car_order.NewDeleteCarOrderController,
	car_order.NewGetCarOrderController,
	car_order.NewGetCarOrdersController,
	// Car Image Controllers
	car_image.NewAddImageController,
	car_image.NewDeleteImageController,
	car_image.NewGetImagesController,
	// Car Tag Controllers
	car_tag.NewGetTagsController,
)

type Controllers struct {
	All []base.Controller
}

func NewControllers(all []base.Controller) *Controllers {
	return &Controllers{All: all}
}

func ProvideControllers(
	// Car Offer Controllers
	createCarOfferController *car_offer.CreateCarOfferController,
	updateCarOfferController *car_offer.UpdateCarOfferController,
	deleteCarOfferController *car_offer.DeleteCarOfferController,
	getCarOfferController *car_offer.GetCarOfferController,
	getCarOffersController *car_offer.GetCarOffersController,
	// Car Order Controllers
	createCarOrderController *car_order.CreateCarOrderController,
	updateCarOrderController *car_order.UpdateCarOrderController,
	deleteCarOrderController *car_order.DeleteCarOrderController,
	getCarOrderController *car_order.GetCarOrderController,
	getCarOrdersController *car_order.GetCarOrdersController,
	// Car Image Controllers
	addImageController *car_image.AddImageController,
	deleteImageController *car_image.DeleteImageController,
	getImagesController *car_image.GetImagesController,
	// Car Tag Controllers
	getTagsController *car_tag.GetTagsController,
) []base.Controller {
	return []base.Controller{
		// Car Offer Controllers
		createCarOfferController,
		updateCarOfferController,
		deleteCarOfferController,
		getCarOfferController,
		getCarOffersController,
		// Car Order Controllers
		createCarOrderController,
		updateCarOrderController,
		deleteCarOrderController,
		getCarOrderController,
		getCarOrdersController,
		// Car Image Controllers
		addImageController,
		deleteImageController,
		getImagesController,
		// Car Tag Controllers
		getTagsController,
	}
}

var WireSet = wire.NewSet(
	ProvideControllers,
	NewControllers,
	// Car Offer Controllers
	car_offer.NewCreateCarOfferController,
	car_offer.NewUpdateCarOfferController,
	car_offer.NewDeleteCarOfferController,
	car_offer.NewGetCarOfferController,
	car_offer.NewGetCarOffersController,
	// Car Order Controllers
	car_order.NewCreateCarOrderController,
	car_order.NewUpdateCarOrderController,
	car_order.NewDeleteCarOrderController,
	car_order.NewGetCarOrderController,
	car_order.NewGetCarOrdersController,
	// Car Image Controllers
	car_image.NewAddImageController,
	car_image.NewDeleteImageController,
	car_image.NewGetImagesController,
	// Car Tag Controllers
	car_tag.NewGetTagsController,
)
