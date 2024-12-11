package contract

import (
	models "identity-api/Domain/models/token"
	"identity-api/Domain/responses"
)

type RefreshTokenResponse struct {
	responses.BaseResponse
	models.JwtToken `json:",inline"`
}

type RefreshTokenResponse200 struct {
	Success bool   `json:"success" example:"true" swaggertype:"boolean"`
	Message string `json:"message" example:"Token refreshed successfully" swaggertype:"string"`
	Token   string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." swaggertype:"string"`
}

type RefreshTokenResponse400 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Invalid request parameters" swaggertype:"string"`
	Token   string `json:"token" example:"" swaggertype:"string"`
}

type RefreshTokenResponse401 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Invalid refresh token" swaggertype:"string"`
	Token   string `json:"token" example:"" swaggertype:"string"`
}

type RefreshTokenResponse500 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Internal server error during token refresh" swaggertype:"string"`
	Token   string `json:"token" example:"" swaggertype:"string"`
}
