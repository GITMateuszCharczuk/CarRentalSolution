package contract

import (
	models "identity-api/Domain/models/token"
)

type RefreshTokenRequest struct {
	models.JwtRefreshToken `json:",inline"`
}
