package service_interfaces

import models "file-storage/Domain/models/external"

type MicroserviceConnector interface {
	ValidateToken(token models.JwtToken) (*models.TokenInfo, error)
	GetUserInternalInfo(token models.JwtToken) (*models.UserInfo, error)
}
