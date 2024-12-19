package mappers

import (
	add_contract "rental-api/Application.contract/images/add_image"
	delete_contract "rental-api/Application.contract/images/delete_image"
	add_commands "rental-api/Application/command_handlers/car_image/add_image"
	delete_commands "rental-api/Application/command_handlers/car_image/delete_image"
)

func MapToAddImageCommand(request *add_contract.AddUrlToCarOfferRequest) add_commands.AddImageCommand {
	return add_commands.AddImageCommand{
		CarOfferId: request.CarOfferId,
		ImageId:    request.ImageId,
		JwtToken:   request.JwtToken,
	}
}

func MapToDeleteImageCommand(request *delete_contract.DeleteImageFromCarOfferRequest) delete_commands.DeleteImageCommand {
	return delete_commands.DeleteImageCommand{
		Id:         request.Id,
		CarOfferId: request.CarOfferId,
		JwtToken:   request.JwtToken,
	}
}
