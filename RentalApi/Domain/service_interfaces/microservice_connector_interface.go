package service_interfaces

import models "rental-api/Domain/models/external"

type MicroserviceConnector interface {
	ValidateToken(token models.JwtToken) (*models.TokenInfo, error)
	GetUserInternalInfo(token models.JwtToken) (*models.UserInfo, error)
	SendEmail(email models.Email) error
}
