package queries

import (
	"context"
	contract "rental-api/Application.contract/car_orders/GetCarOrders"
	repository_interfaces "rental-api/Domain/repository_interfaces/car_order_repository"
	"rental-api/Domain/responses"
)

type GetCarOrdersQueryHandler struct {
	carOrderQueryRepository repository_interfaces.CarOrderQueryRepository
}

func NewGetCarOrdersQueryHandler(
	carOrderQueryRepository repository_interfaces.CarOrderQueryRepository,
) *GetCarOrdersQueryHandler {
	return &GetCarOrdersQueryHandler{
		carOrderQueryRepository: carOrderQueryRepository,
	}
}

func (h *GetCarOrdersQueryHandler) Handle(ctx context.Context, query *GetCarOrdersQuery) (*contract.GetCarOrdersResponse, error) {
	result, err := h.carOrderQueryRepository.GetCarOrders(
		&query.Pagination,
		&query.Sortable,
		query.StartDate,
		query.EndDate,
		query.UserId,
		query.CarOfferId,
		query.Status,
		query.DateFilterType,
	)

	if err != nil {
		response := responses.NewResponse[contract.GetCarOrdersResponse](500, "Failed to retrieve car orders")
		return &response, nil
	}

	return &contract.GetCarOrdersResponse{
		BaseResponse:    responses.NewBaseResponse(200, "Car orders retrieved successfully"),
		PaginatedResult: *result,
	}, nil
}
