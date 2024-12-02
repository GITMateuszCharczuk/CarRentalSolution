package contract

import "identity-api/Domain/models"

type GetUserInfoRequest struct {
	Id              string `json:"id" example:"user.id.here" swaggertype:"string"`
	models.JwtToken `json:",inline"`
}
