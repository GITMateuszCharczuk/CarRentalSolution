package commands

import (
	"context"
	contract "rental-api/Application.contract/car_orders/CreateCarOrder"
	models "rental-api/Domain/models/domestic"
	car_offer_repository "rental-api/Domain/repository_interfaces/car_offer_repository"
	repository_interfaces "rental-api/Domain/repository_interfaces/car_order_repository"
	"rental-api/Domain/responses"
	data_fetcher "rental-api/Domain/service_interfaces"
)

type CreateCarOrderCommandHandler struct {
	carOrderCommandRepository repository_interfaces.CarOrderCommandRepository
	carOfferQueryRepository   car_offer_repository.CarOfferQueryRepository
	dataFetcher               data_fetcher.MicroserviceConnector
}

func NewCreateCarOrderCommandHandler(
	carOrderCommandRepository repository_interfaces.CarOrderCommandRepository,
	carOfferQueryRepository car_offer_repository.CarOfferQueryRepository,
	dataFetcher data_fetcher.MicroserviceConnector,
) *CreateCarOrderCommandHandler {
	return &CreateCarOrderCommandHandler{
		carOrderCommandRepository: carOrderCommandRepository,
		carOfferQueryRepository:   carOfferQueryRepository,
		dataFetcher:               dataFetcher,
	}
}

func (h *CreateCarOrderCommandHandler) Handle(ctx context.Context, command *CreateCarOrderCommand) (*contract.CreateCarOrderResponse, error) {
	userInfo, err := h.dataFetcher.GetUserInternalInfo(command.JwtToken)
	if err != nil {
		response := responses.NewResponse[contract.CreateCarOrderResponse](401, "Unauthorized")
		return &response, nil
	}

	carOffer, err := h.carOfferQueryRepository.GetCarOfferByID(command.CarOfferId)
	if err != nil || carOffer == nil {
		response := responses.NewResponse[contract.CreateCarOrderResponse](404, "Car offer not found")
		return &response, nil
	}

	carOrder := &models.CarOrderModel{
		UserId:           userInfo.ID,
		CarOfferId:       command.CarOfferId,
		StartDate:        command.StartDate,
		EndDate:          command.EndDate,
		DeliveryLocation: command.DeliveryLocation,
		ReturnLocation:   command.ReturnLocation,
		NumOfDrivers:     command.NumOfDrivers,
		TotalCost:        command.TotalCost,
	}

	result, err := h.carOrderCommandRepository.CreateCarOrder(ctx, carOrder)
	if err != nil {
		response := responses.NewResponse[contract.CreateCarOrderResponse](500, "Failed to create car order")
		return &response, nil
	}

	return &contract.CreateCarOrderResponse{
		BaseResponse: responses.NewBaseResponse(200, "Car order created successfully"),
		Id:           *result,
	}, nil
}
