package commands

import (
	"context"
	contract "rental-api/Application.contract/car_offers/update_car_offer"
	"rental-api/Application/services"
	"rental-api/Domain/constants"
	models "rental-api/Domain/models/domestic"
	repository_interfaces "rental-api/Domain/repository_interfaces/car_offer_repository"
	"rental-api/Domain/responses"
	data_fetcher "rental-api/Domain/service_interfaces"
)

type UpdateCarOfferCommandHandler struct {
	carOfferCommandRepository repository_interfaces.CarOfferCommandRepository
	carOfferQueryRepository   repository_interfaces.CarOfferQueryRepository
	dataFetcher               data_fetcher.MicroserviceConnector
}

func NewUpdateCarOfferCommandHandler(
	carOfferCommandRepository repository_interfaces.CarOfferCommandRepository,
	carOfferQueryRepository repository_interfaces.CarOfferQueryRepository,
	dataFetcher data_fetcher.MicroserviceConnector,
) *UpdateCarOfferCommandHandler {
	return &UpdateCarOfferCommandHandler{
		carOfferCommandRepository: carOfferCommandRepository,
		carOfferQueryRepository:   carOfferQueryRepository,
		dataFetcher:               dataFetcher,
	}
}

func (h *UpdateCarOfferCommandHandler) Handle(ctx context.Context, command *UpdateCarOfferCommand) (*contract.UpdateCarOfferResponse, error) {
	userInfo, err := h.dataFetcher.GetUserInternalInfo(command.JwtToken)
	if err != nil {
		response := responses.NewResponse[contract.UpdateCarOfferResponse](401, "Unauthorized")
		return &response, nil
	}

	existingOffer, err := h.carOfferQueryRepository.GetCarOfferByID(command.Id)
	if err != nil || existingOffer == nil {
		response := responses.NewResponse[contract.UpdateCarOfferResponse](404, "Car offer not found")
		return &response, nil
	}

	if !services.IsRole(constants.SuperAdmin, userInfo.Roles) {
		if existingOffer.CustodianId != userInfo.ID && !services.IsRole(constants.Admin, userInfo.Roles) {
			response := responses.NewResponse[contract.UpdateCarOfferResponse](403, "Not authorized to delete this car offer")
			return &response, nil
		}
	}

	carOffer := &models.CarOfferModel{
		Id:                 command.Id,
		Heading:            command.Heading,
		ShortDescription:   command.ShortDescription,
		FeaturedImageUrl:   command.FeaturedImageUrl,
		UrlHandle:          command.UrlHandle,
		Horsepower:         command.Horsepower,
		YearOfProduction:   command.YearOfProduction,
		EngineDetails:      command.EngineDetails,
		DriveDetails:       command.DriveDetails,
		GearboxDetails:     command.GearboxDetails,
		Visible:            command.Visible,
		OneNormalDayPrice:  command.OneNormalDayPrice,
		OneWeekendDayPrice: command.OneWeekendDayPrice,
		OneWeekPrice:       command.OneWeekPrice,
		OneMonthPrice:      command.OneMonthPrice,
		CustodianId:        existingOffer.CustodianId,
		CustodianEmail:     existingOffer.CustodianEmail,
		CreatedAt:          existingOffer.CreatedAt,
	}

	err = h.carOfferCommandRepository.UpdateCarOffer(ctx, carOffer, command.Tags, nil)
	if err != nil {
		response := responses.NewResponse[contract.UpdateCarOfferResponse](500, "Failed to update car offer")
		return &response, nil
	}

	return &contract.UpdateCarOfferResponse{
		BaseResponse: responses.NewBaseResponse(200, "Car offer updated successfully"),
	}, nil
}
