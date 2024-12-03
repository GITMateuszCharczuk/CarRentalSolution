package contract

import (
	models "identity-api/Domain/models/token"
	responses "identity-api/Domain/responses"
)

type LoginResponse struct {
	responses.BaseResponse
	models.JwtToken        `json:",inline"`
	models.JwtRefreshToken `json:",inline"`
	Roles                  []string `json:"roles" example:"[user, admin]" swaggertype:"array,string"`
}
