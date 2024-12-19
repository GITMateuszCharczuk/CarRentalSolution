package queries

import (
	"context"
	contract "rental-api/Application.contract/car_orders/GetCarOrder"
	repository_interfaces "rental-api/Domain/repository_interfaces/car_order_repository"
	"rental-api/Domain/responses"
)

type GetCarOrderQueryHandler struct {
	carOrderQueryRepository repository_interfaces.CarOrderQueryRepository
}

func NewGetCarOrderQueryHandler(
	carOrderQueryRepository repository_interfaces.CarOrderQueryRepository,
) *GetCarOrderQueryHandler {
	return &GetCarOrderQueryHandler{
		carOrderQueryRepository: carOrderQueryRepository,
	}
}

func (h *GetCarOrderQueryHandler) Handle(ctx context.Context, query *GetCarOrderQuery) (*contract.GetCarOrderResponse, error) {
	carOrder, err := h.carOrderQueryRepository.GetCarOrderByID(query.ID)
	if err != nil {
		response := responses.NewResponse[contract.GetCarOrderResponse](500, "Failed to retrieve car order")
		return &response, nil
	}

	if carOrder == nil {
		response := responses.NewResponse[contract.GetCarOrderResponse](404, "Car order not found")
		return &response, nil
	}

	return &contract.GetCarOrderResponse{
		BaseResponse: responses.NewBaseResponse(200, "Car order retrieved successfully"),
		CarOrder:     *carOrder,
	}, nil
}
