package mappers

import (
	"rental-api/API/services"
	create_contract "rental-api/Application.contract/car_offers/create_car_offer"
	delete_contract "rental-api/Application.contract/car_offers/delete_car_offer"
	get_car_offer_contract "rental-api/Application.contract/car_offers/get_car_offer"
	get_car_offers_contract "rental-api/Application.contract/car_offers/get_car_offers"
	update_contract "rental-api/Application.contract/car_offers/update_car_offer"
	create_commands "rental-api/Application/command_handlers/car_offer/create_car_offer"
	delete_commands "rental-api/Application/command_handlers/car_offer/delete_car_offer"
	update_commands "rental-api/Application/command_handlers/car_offer/update_car_offer"
	get_car_offer_queries "rental-api/Application/query_handlers/car_offer/get_car_offer"
	get_car_offers_queries "rental-api/Application/query_handlers/car_offer/get_car_offers"
)

func MapToCreateCarOfferCommand(request *create_contract.CreateCarOfferRequest) create_commands.CreateCarOfferCommand {
	return create_commands.CreateCarOfferCommand{
		Heading:            request.Heading,
		ShortDescription:   request.ShortDescription,
		FeaturedImageUrl:   request.FeaturedImageUrl,
		UrlHandle:          request.UrlHandle,
		Horsepower:         request.Horsepower,
		YearOfProduction:   request.YearOfProduction,
		EngineDetails:      request.EngineDetails,
		DriveDetails:       request.DriveDetails,
		GearboxDetails:     request.GearboxDetails,
		Visible:            request.Visible,
		OneNormalDayPrice:  request.OneNormalDayPrice,
		OneWeekendDayPrice: request.OneWeekendDayPrice,
		OneWeekPrice:       request.OneWeekPrice,
		OneMonthPrice:      request.OneMonthPrice,
		Tags:               request.Tags,
		ImageUrls:          request.ImageUrls,
		JwtToken:           request.JwtToken,
	}
}

func MapToUpdateCarOfferCommand(request *update_contract.UpdateCarOfferRequest) update_commands.UpdateCarOfferCommand {
	return update_commands.UpdateCarOfferCommand{
		Id:                 request.Id,
		Heading:            request.Heading,
		ShortDescription:   request.ShortDescription,
		FeaturedImageUrl:   request.FeaturedImageUrl,
		UrlHandle:          request.UrlHandle,
		Horsepower:         request.Horsepower,
		YearOfProduction:   request.YearOfProduction,
		EngineDetails:      request.EngineDetails,
		DriveDetails:       request.DriveDetails,
		GearboxDetails:     request.GearboxDetails,
		Visible:            request.Visible,
		OneNormalDayPrice:  request.OneNormalDayPrice,
		OneWeekendDayPrice: request.OneWeekendDayPrice,
		OneWeekPrice:       request.OneWeekPrice,
		OneMonthPrice:      request.OneMonthPrice,
		Tags:               request.Tags,
		JwtToken:           request.JwtToken,
	}
}

func MapToDeleteCarOfferCommand(request *delete_contract.DeleteCarOfferRequest) delete_commands.DeleteCarOfferCommand {
	return delete_commands.DeleteCarOfferCommand{
		ID:       request.CarOfferId,
		JwtToken: request.JwtToken,
	}
}

func MapToGetCarOffersQuery(request *get_car_offers_contract.GetCarOffersRequest) get_car_offers_queries.GetCarOffersQuery {
	return get_car_offers_queries.GetCarOffersQuery{
		Pagination:   request.Pagination,
		Sortable:     services.ExtractSorting(request.SortQuery),
		Ids:          request.Ids,
		DateTimeFrom: request.DateTimeFrom,
		DateTimeTo:   request.DateTimeTo,
		Tags:         request.Tags,
		Visible:      request.Visible,
		Status:       request.Status,
	}
}

func MapToGetCarOfferQuery(request *get_car_offer_contract.GetCarOfferRequest) get_car_offer_queries.GetCarOfferQuery {
	return get_car_offer_queries.GetCarOfferQuery{
		ID: request.CarOfferId,
	}
}
