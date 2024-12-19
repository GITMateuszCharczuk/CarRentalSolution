package commands

import (
	"context"
	contract "rental-api/Application.contract/car_offers/delete_car_offer"
	"rental-api/Application/services"
	"rental-api/Domain/constants"
	repository_interfaces "rental-api/Domain/repository_interfaces/car_offer_repository"
	"rental-api/Domain/responses"
	data_fetcher "rental-api/Domain/service_interfaces"
)

type DeleteCarOfferCommandHandler struct {
	carOfferCommandRepository repository_interfaces.CarOfferCommandRepository
	carOfferQueryRepository   repository_interfaces.CarOfferQueryRepository
	dataFetcher               data_fetcher.MicroserviceConnector
}

func NewDeleteCarOfferCommandHandler(
	carOfferCommandRepository repository_interfaces.CarOfferCommandRepository,
	carOfferQueryRepository repository_interfaces.CarOfferQueryRepository,
	dataFetcher data_fetcher.MicroserviceConnector,
) *DeleteCarOfferCommandHandler {
	return &DeleteCarOfferCommandHandler{
		carOfferCommandRepository: carOfferCommandRepository,
		carOfferQueryRepository:   carOfferQueryRepository,
		dataFetcher:               dataFetcher,
	}
}

func (h *DeleteCarOfferCommandHandler) Handle(ctx context.Context, command *DeleteCarOfferCommand) (*contract.DeleteCarOfferResponse, error) {
	userInfo, err := h.dataFetcher.GetUserInternalInfo(command.JwtToken)
	if err != nil {
		response := responses.NewResponse[contract.DeleteCarOfferResponse](401, "Unauthorized")
		return &response, nil
	}

	existingOffer, err := h.carOfferQueryRepository.GetCarOfferByID(command.ID)
	if err != nil || existingOffer == nil {
		response := responses.NewResponse[contract.DeleteCarOfferResponse](404, "Car offer not found")
		return &response, nil
	}

	if !services.IsRole(constants.SuperAdmin, userInfo.Roles) {
		if existingOffer.CustodianId != userInfo.ID && !services.IsRole(constants.Admin, userInfo.Roles) {
			response := responses.NewResponse[contract.DeleteCarOfferResponse](403, "Not authorized to delete this car offer")
			return &response, nil
		}
	}

	err = h.carOfferCommandRepository.DeleteCarOffer(ctx, command.ID)
	if err != nil {
		response := responses.NewResponse[contract.DeleteCarOfferResponse](500, "Failed to delete car offer")
		return &response, nil
	}

	return &contract.DeleteCarOfferResponse{
		BaseResponse: responses.NewBaseResponse(200, "Car offer deleted successfully"),
	}, nil
}
