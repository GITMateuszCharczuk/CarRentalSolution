package contract

import "identity-api/Domain/models"

type RefreshTokenRequest struct {
	models.JwtRefreshToken `json:",inline"`
}
