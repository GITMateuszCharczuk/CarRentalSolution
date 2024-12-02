package contract

import (
	"identity-api/Domain/models"
	"identity-api/Domain/responses"
)

type LoginResponse struct {
	responses.BaseResponse
	models.JwtToken `json:",inline"`
	Roles           []string `json:"roles" example:"[user, admin]" swaggertype:"array,string"`
}
