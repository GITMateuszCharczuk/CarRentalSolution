package contract

import models "identity-api/Domain/models/token"

type DeleteUserRequest struct {
	ID              string `json:"id" binding:"required" example:"12345" swaggertype:"string"`
	models.JwtToken `json:",inline"`
}
