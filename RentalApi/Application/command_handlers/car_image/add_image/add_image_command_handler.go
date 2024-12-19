package commands

import (
	"context"
	contract "rental-api/Application.contract/images/add_image"
	"rental-api/Application/services"
	"rental-api/Domain/constants"
	repository_interfaces "rental-api/Domain/repository_interfaces/car_image_repository"
	car_offer_repository "rental-api/Domain/repository_interfaces/car_offer_repository"
	"rental-api/Domain/responses"
	data_fetcher "rental-api/Domain/service_interfaces"
)

type AddImageCommandHandler struct {
	carImageCommandRepository repository_interfaces.CarImageCommandRepository
	carOfferQueryRepository   car_offer_repository.CarOfferQueryRepository
	dataFetcher               data_fetcher.MicroserviceConnector
}

func NewAddImageCommandHandler(
	carImageCommandRepository repository_interfaces.CarImageCommandRepository,
	carOfferQueryRepository car_offer_repository.CarOfferQueryRepository,
	dataFetcher data_fetcher.MicroserviceConnector,
) *AddImageCommandHandler {
	return &AddImageCommandHandler{
		carImageCommandRepository: carImageCommandRepository,
		carOfferQueryRepository:   carOfferQueryRepository,
		dataFetcher:               dataFetcher,
	}
}

func (h *AddImageCommandHandler) Handle(ctx context.Context, command *AddImageCommand) (*contract.AddUrlToCarOfferResponse, error) {
	userInfo, err := h.dataFetcher.GetUserInternalInfo(command.JwtToken)
	if err != nil {
		response := responses.NewResponse[contract.AddUrlToCarOfferResponse](401, "Unauthorized")
		return &response, nil
	}

	carOffer, err := h.carOfferQueryRepository.GetCarOfferByID(command.CarOfferId)
	if err != nil || carOffer == nil {
		response := responses.NewResponse[contract.AddUrlToCarOfferResponse](404, "Car offer not found")
		return &response, nil
	}

	if !services.IsRole(constants.SuperAdmin, userInfo.Roles) {
		if carOffer.CustodianId != userInfo.ID && !services.IsRole(constants.Admin, userInfo.Roles) {
			response := responses.NewResponse[contract.AddUrlToCarOfferResponse](403, "Not authorized to delete this car offer")
			return &response, nil
		}
	}

	result, err := h.carImageCommandRepository.AddImageToCarOffer(ctx, command.CarOfferId, command.ImageId)
	if err != nil {
		response := responses.NewResponse[contract.AddUrlToCarOfferResponse](500, "Failed to add image")
		return &response, nil
	}

	return &contract.AddUrlToCarOfferResponse{
		BaseResponse: responses.NewBaseResponse(200, "Image added successfully"),
		Id:           *result,
	}, nil
}
