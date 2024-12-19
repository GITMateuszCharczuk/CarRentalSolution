package commands

import (
	"context"
	contract "rental-api/Application.contract/car_orders/DeleteCarOrder"
	"rental-api/Application/services"
	repository_interfaces "rental-api/Domain/repository_interfaces/car_order_repository"
	"rental-api/Domain/responses"
	data_fetcher "rental-api/Domain/service_interfaces"
)

type DeleteCarOrderCommandHandler struct {
	carOrderCommandRepository repository_interfaces.CarOrderCommandRepository
	carOrderQueryRepository   repository_interfaces.CarOrderQueryRepository
	dataFetcher               data_fetcher.MicroserviceConnector
}

func NewDeleteCarOrderCommandHandler(
	carOrderCommandRepository repository_interfaces.CarOrderCommandRepository,
	carOrderQueryRepository repository_interfaces.CarOrderQueryRepository,
	dataFetcher data_fetcher.MicroserviceConnector,
) *DeleteCarOrderCommandHandler {
	return &DeleteCarOrderCommandHandler{
		carOrderCommandRepository: carOrderCommandRepository,
		carOrderQueryRepository:   carOrderQueryRepository,
		dataFetcher:               dataFetcher,
	}
}

func (h *DeleteCarOrderCommandHandler) Handle(ctx context.Context, command *DeleteCarOrderCommand) (*contract.DeleteCarOrderResponse, error) {
	userInfo, err := h.dataFetcher.GetUserInternalInfo(command.JwtToken)
	if err != nil {
		response := responses.NewResponse[contract.DeleteCarOrderResponse](401, "Unauthorized")
		return &response, nil
	}

	existingOrder, err := h.carOrderQueryRepository.GetCarOrderByID(command.ID)
	if err != nil || existingOrder == nil {
		response := responses.NewResponse[contract.DeleteCarOrderResponse](404, "Car order not found")
		return &response, nil
	}

	if existingOrder.UserId != userInfo.ID && !services.IsAdminOrSuperAdmin(userInfo.Roles) {
		response := responses.NewResponse[contract.DeleteCarOrderResponse](403, "Not authorized to delete this car order")
		return &response, nil
	}

	err = h.carOrderCommandRepository.DeleteCarOrder(ctx, command.ID)
	if err != nil {
		response := responses.NewResponse[contract.DeleteCarOrderResponse](500, "Failed to delete car order")
		return &response, nil
	}

	return &contract.DeleteCarOrderResponse{
		BaseResponse: responses.NewBaseResponse(200, "Car order deleted successfully"),
	}, nil
}
