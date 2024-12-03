package contract

import (
	models "identity-api/Domain/models/token"
)

type GetUserInfoRequest struct {
	Id              string `json:"id" example:"user.id.here" swaggertype:"string"`
	models.JwtToken `json:",inline"`
}
