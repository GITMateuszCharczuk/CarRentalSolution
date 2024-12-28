package mappers

import (
	"rental-api/API/services"
	create_contract "rental-api/Application.contract/car_orders/CreateCarOrder"
	delete_contract "rental-api/Application.contract/car_orders/DeleteCarOrder"
	get_car_order_contract "rental-api/Application.contract/car_orders/GetCarOrder"
	get_car_orders_contract "rental-api/Application.contract/car_orders/GetCarOrders"
	update_contract "rental-api/Application.contract/car_orders/UpdateCarOrder"
	create_commands "rental-api/Application/command_handlers/car_order/create_car_order"
	delete_commands "rental-api/Application/command_handlers/car_order/delete_car_order"
	update_commands "rental-api/Application/command_handlers/car_order/update_car_order"
	get_car_order_queries "rental-api/Application/query_handlers/car_order/get_car_order"
	get_car_orders_queries "rental-api/Application/query_handlers/car_order/get_car_orders"
	"rental-api/Domain/constants"
)

func MapToCreateCarOrderCommand(request *create_contract.CreateCarOrderRequest) create_commands.CreateCarOrderCommand {
	return create_commands.CreateCarOrderCommand{
		CarOfferId:       request.CarOfferId,
		StartDate:        request.StartDate,
		EndDate:          request.EndDate,
		DeliveryLocation: request.DeliveryLocation,
		ReturnLocation:   request.ReturnLocation,
		NumOfDrivers:     request.NumOfDrivers,
		TotalCost:        request.TotalCost,
		Status:           string(constants.OrderStatusPending),
		JwtToken:         request.JwtToken,
	}
}

func MapToUpdateCarOrderCommand(request *update_contract.UpdateCarOrderRequest) update_commands.UpdateCarOrderCommand {
	return update_commands.UpdateCarOrderCommand{
		Id:               request.Id,
		CarOfferId:       request.CarOfferId,
		StartDate:        request.StartDate,
		EndDate:          request.EndDate,
		DeliveryLocation: request.DeliveryLocation,
		ReturnLocation:   request.ReturnLocation,
		NumOfDrivers:     request.NumOfDrivers,
		TotalCost:        request.TotalCost,
		Status:           request.Status,
		JwtToken:         request.JwtToken,
	}
}

func MapToDeleteCarOrderCommand(request *delete_contract.DeleteCarOrderRequest) delete_commands.DeleteCarOrderCommand {
	return delete_commands.DeleteCarOrderCommand{
		ID:       request.CarOrderId,
		JwtToken: request.JwtToken,
	}
}

func MapToGetCarOrderQuery(request *get_car_order_contract.GetCarOrderRequest) get_car_order_queries.GetCarOrderQuery {
	return get_car_order_queries.GetCarOrderQuery{
		ID:       request.CarOrderId,
		JwtToken: request.JwtToken,
	}
}

func MapToGetCarOrdersQuery(request *get_car_orders_contract.GetCarOrdersRequest) get_car_orders_queries.GetCarOrdersQuery {
	return get_car_orders_queries.GetCarOrdersQuery{
		Pagination:     request.Pagination,
		Sortable:       services.ExtractSorting(request.SortQuery),
		StartDate:      request.StartDate,
		EndDate:        request.EndDate,
		UserId:         request.UserId,
		CarOfferId:     request.CarOfferId,
		Statuses:       request.Statuses,
		DateFilterType: request.DateFilterType,
		JwtToken:       request.JwtToken,
	}
}
