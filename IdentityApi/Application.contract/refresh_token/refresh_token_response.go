package contract

import (
	"identity-api/Domain/models"
	"identity-api/Domain/responses"
)

type RefreshTokenResponse struct {
	responses.BaseResponse
	models.JwtToken `json:",inline"`
}
