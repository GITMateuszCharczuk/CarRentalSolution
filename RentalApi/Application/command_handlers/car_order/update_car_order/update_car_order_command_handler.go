package commands

import (
	"context"
	contract "rental-api/Application.contract/car_orders/UpdateCarOrder"
	"rental-api/Application/services"
	models "rental-api/Domain/models/domestic"
	car_offer_repository "rental-api/Domain/repository_interfaces/car_offer_repository"
	repository_interfaces "rental-api/Domain/repository_interfaces/car_order_repository"
	"rental-api/Domain/responses"
	data_fetcher "rental-api/Domain/service_interfaces"
)

type UpdateCarOrderCommandHandler struct {
	carOrderCommandRepository repository_interfaces.CarOrderCommandRepository
	carOrderQueryRepository   repository_interfaces.CarOrderQueryRepository
	carOfferQueryRepository   car_offer_repository.CarOfferQueryRepository
	dataFetcher               data_fetcher.MicroserviceConnector
}

func NewUpdateCarOrderCommandHandler(
	carOrderCommandRepository repository_interfaces.CarOrderCommandRepository,
	carOrderQueryRepository repository_interfaces.CarOrderQueryRepository,
	carOfferQueryRepository car_offer_repository.CarOfferQueryRepository,
	dataFetcher data_fetcher.MicroserviceConnector,
) *UpdateCarOrderCommandHandler {
	return &UpdateCarOrderCommandHandler{
		carOrderCommandRepository: carOrderCommandRepository,
		carOrderQueryRepository:   carOrderQueryRepository,
		carOfferQueryRepository:   carOfferQueryRepository,
		dataFetcher:               dataFetcher,
	}
}

func (h *UpdateCarOrderCommandHandler) Handle(ctx context.Context, command *UpdateCarOrderCommand) (*contract.UpdateCarOrderResponse, error) {
	userInfo, err := h.dataFetcher.GetUserInternalInfo(command.JwtToken)
	if err != nil {
		response := responses.NewResponse[contract.UpdateCarOrderResponse](401, "Unauthorized")
		return &response, nil
	}

	existingOrder, err := h.carOrderQueryRepository.GetCarOrderByID(command.Id)
	if err != nil || existingOrder == nil {
		response := responses.NewResponse[contract.UpdateCarOrderResponse](404, "Car order not found")
		return &response, nil
	}

	if existingOrder.UserId != userInfo.ID && !services.IsAdminOrSuperAdmin(userInfo.Roles) {
		response := responses.NewResponse[contract.UpdateCarOrderResponse](403, "Not authorized to update this car order")
		return &response, nil
	}

	carOffer, err := h.carOfferQueryRepository.GetCarOfferByID(command.CarOfferId)
	if err != nil || carOffer == nil {
		response := responses.NewResponse[contract.UpdateCarOrderResponse](404, "Car offer not found")
		return &response, nil
	}

	carOrder := &models.CarOrderModel{
		Id:               command.Id,
		UserId:           existingOrder.UserId,
		CarOfferId:       command.CarOfferId,
		StartDate:        command.StartDate,
		EndDate:          command.EndDate,
		DeliveryLocation: command.DeliveryLocation,
		ReturnLocation:   command.ReturnLocation,
		NumOfDrivers:     command.NumOfDrivers,
		TotalCost:        command.TotalCost,
	}

	err = h.carOrderCommandRepository.UpdateCarOrder(ctx, carOrder)
	if err != nil {
		response := responses.NewResponse[contract.UpdateCarOrderResponse](500, "Failed to update car order")
		return &response, nil
	}

	return &contract.UpdateCarOrderResponse{
		BaseResponse: responses.NewBaseResponse(200, "Car order updated successfully"),
	}, nil
}
