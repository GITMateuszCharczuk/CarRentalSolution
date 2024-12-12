package datafetcher

import models "identity-api/Domain/models/external"

type DataFetcher interface {
	ValidateToken(token models.JwtToken) (*models.TokenInfo, error)
	GetUserInternalInfo(token models.JwtToken) (*models.UserInfo, error)
}
