package commands

import (
	"context"
	contract "rental-api/Application.contract/car_offers/create_car_offer"
	models "rental-api/Domain/models/domestic"
	repository_interfaces "rental-api/Domain/repository_interfaces/car_offer_repository"
	"rental-api/Domain/responses"
	data_fetcher "rental-api/Domain/service_interfaces"
	"time"
)

type CreateCarOfferCommandHandler struct {
	carOfferCommandRepository repository_interfaces.CarOfferCommandRepository
	dataFetcher               data_fetcher.MicroserviceConnector
}

func NewCreateCarOfferCommandHandler(
	carOfferCommandRepository repository_interfaces.CarOfferCommandRepository,
	dataFetcher data_fetcher.MicroserviceConnector,
) *CreateCarOfferCommandHandler {
	return &CreateCarOfferCommandHandler{
		carOfferCommandRepository: carOfferCommandRepository,
		dataFetcher:               dataFetcher,
	}
}

func (h *CreateCarOfferCommandHandler) Handle(ctx context.Context, command *CreateCarOfferCommand) (*contract.CreateCarOfferResponse, error) {
	userInfo, err := h.dataFetcher.GetUserInternalInfo(command.JwtToken)
	if err != nil {
		response := responses.NewResponse[contract.CreateCarOfferResponse](401, "Unauthorized")
		return &response, nil
	}

	now := time.Now()
	carOffer := &models.CarOfferModel{
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
		CustodianId:        userInfo.ID,
		CustodianEmail:     userInfo.Email,
		CreatedAt:          now.Format(time.RFC3339),
	}

	result, err := h.carOfferCommandRepository.CreateCarOffer(ctx, carOffer, command.Tags, command.ImageUrls)
	if err != nil {
		response := responses.NewResponse[contract.CreateCarOfferResponse](500, "Failed to create car offer")
		return &response, nil
	}

	return &contract.CreateCarOfferResponse{
		BaseResponse: responses.NewBaseResponse(200, "Car offer created successfully"),
		Id:           *result,
	}, nil
}
