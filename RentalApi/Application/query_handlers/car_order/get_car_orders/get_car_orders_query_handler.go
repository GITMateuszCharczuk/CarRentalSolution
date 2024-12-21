package queries

import (
	"context"
	contract "rental-api/Application.contract/car_orders/GetCarOrders"
	"rental-api/Application/services"
	repository_interfaces "rental-api/Domain/repository_interfaces/car_order_repository"
	"rental-api/Domain/responses"
	microservice_connector "rental-api/Domain/service_interfaces"
)

type GetCarOrdersQueryHandler struct {
	carOrderQueryRepository repository_interfaces.CarOrderQueryRepository
	connector               microservice_connector.MicroserviceConnector
}

func NewGetCarOrdersQueryHandler(
	carOrderQueryRepository repository_interfaces.CarOrderQueryRepository,
	connector microservice_connector.MicroserviceConnector,
) *GetCarOrdersQueryHandler {
	return &GetCarOrdersQueryHandler{
		carOrderQueryRepository: carOrderQueryRepository,
		connector:               connector,
	}
}

func (h *GetCarOrdersQueryHandler) Handle(ctx context.Context, query *GetCarOrdersQuery) (*contract.GetCarOrdersResponse, error) {
	tokenInfo, err := h.connector.ValidateToken(query.JwtToken)
	if err != nil {
		response := responses.NewResponse[contract.GetCarOrdersResponse](401, "Unauthorized")
		return &response, nil
	}
	if !services.IsAdminOrSuperAdmin(tokenInfo.Roles) {
		response := responses.NewResponse[contract.GetCarOrdersResponse](403, "Forbidden")
		return &response, nil
	}
	result, err := h.carOrderQueryRepository.GetCarOrders(
		&query.Pagination,
		&query.Sortable,
		query.StartDate,
		query.EndDate,
		query.UserId,
		query.CarOfferId,
		query.Statuses,
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
