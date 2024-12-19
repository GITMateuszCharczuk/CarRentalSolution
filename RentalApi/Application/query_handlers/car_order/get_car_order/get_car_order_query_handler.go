package queries

import (
	"context"
	contract "rental-api/Application.contract/car_orders/GetCarOrder"
	"rental-api/Application/services"
	repository_interfaces "rental-api/Domain/repository_interfaces/car_order_repository"
	"rental-api/Domain/responses"
	microservice_connector "rental-api/Domain/service_interfaces"
)

type GetCarOrderQueryHandler struct {
	carOrderQueryRepository repository_interfaces.CarOrderQueryRepository
	connector               microservice_connector.MicroserviceConnector
}

func NewGetCarOrderQueryHandler(
	carOrderQueryRepository repository_interfaces.CarOrderQueryRepository,
	connector microservice_connector.MicroserviceConnector,
) *GetCarOrderQueryHandler {
	return &GetCarOrderQueryHandler{
		carOrderQueryRepository: carOrderQueryRepository,
		connector:               connector,
	}
}

func (h *GetCarOrderQueryHandler) Handle(ctx context.Context, query *GetCarOrderQuery) (*contract.GetCarOrderResponse, error) {
	userInfo, err := h.connector.GetUserInternalInfo(query.JwtToken)
	if err != nil {
		response := responses.NewResponse[contract.GetCarOrderResponse](401, "Unauthorized")
		return &response, nil
	}
	carOrder, err := h.carOrderQueryRepository.GetCarOrderByID(query.ID)
	if !services.IsAdminOrSuperAdmin(userInfo.Roles) || carOrder.UserId != userInfo.ID {
		response := responses.NewResponse[contract.GetCarOrderResponse](403, "Forbidden")
		return &response, nil
	}
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			response := responses.NewResponse[contract.GetCarOrderResponse](404, "Car order not found")
			return &response, nil
		}
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
