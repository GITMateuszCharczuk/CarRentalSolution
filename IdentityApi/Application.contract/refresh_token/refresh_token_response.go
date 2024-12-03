package contract

import (
	models "identity-api/Domain/models/token"
	"identity-api/Domain/responses"
)

type RefreshTokenResponse struct {
	responses.BaseResponse
	models.JwtToken `json:",inline"`
}
