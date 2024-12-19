package commands

import (
	"context"
	contract "rental-api/Application.contract/images/delete_image"
	"rental-api/Application/services"
	"rental-api/Domain/constants"
	repository_interfaces "rental-api/Domain/repository_interfaces/car_image_repository"
	car_offer_repository "rental-api/Domain/repository_interfaces/car_offer_repository"
	"rental-api/Domain/responses"
	data_fetcher "rental-api/Domain/service_interfaces"
)

type DeleteImageCommandHandler struct {
	carImageCommandRepository repository_interfaces.CarImageCommandRepository
	carOfferQueryRepository   car_offer_repository.CarOfferQueryRepository
	dataFetcher               data_fetcher.MicroserviceConnector
}

func NewDeleteImageCommandHandler(
	carImageCommandRepository repository_interfaces.CarImageCommandRepository,
	carOfferQueryRepository car_offer_repository.CarOfferQueryRepository,
	dataFetcher data_fetcher.MicroserviceConnector,
) *DeleteImageCommandHandler {
	return &DeleteImageCommandHandler{
		carImageCommandRepository: carImageCommandRepository,
		carOfferQueryRepository:   carOfferQueryRepository,
		dataFetcher:               dataFetcher,
	}
}

func (h *DeleteImageCommandHandler) Handle(ctx context.Context, command *DeleteImageCommand) (*contract.DeleteImageFromCarOfferResponse, error) {
	userInfo, err := h.dataFetcher.GetUserInternalInfo(command.JwtToken)
	if err != nil {
		response := responses.NewResponse[contract.DeleteImageFromCarOfferResponse](401, "Unauthorized")
		return &response, nil
	}

	carOffer, err := h.carOfferQueryRepository.GetCarOfferByID(command.CarOfferId)
	if err != nil || carOffer == nil {
		response := responses.NewResponse[contract.DeleteImageFromCarOfferResponse](404, "Car offer not found")
		return &response, nil
	}

	if !services.IsRole(constants.SuperAdmin, userInfo.Roles) {
		if carOffer.CustodianId != userInfo.ID && !services.IsRole(constants.Admin, userInfo.Roles) {
			response := responses.NewResponse[contract.DeleteImageFromCarOfferResponse](403, "Not authorized to delete this car offer")
			return &response, nil
		}
	}

	err = h.carImageCommandRepository.DeleteImageFromCarOffer(ctx, command.CarOfferId, command.Id)
	if err != nil {
		response := responses.NewResponse[contract.DeleteImageFromCarOfferResponse](500, "Failed to delete image")
		return &response, nil
	}

	return &contract.DeleteImageFromCarOfferResponse{
		BaseResponse: responses.NewBaseResponse(200, "Image deleted successfully"),
	}, nil
}
