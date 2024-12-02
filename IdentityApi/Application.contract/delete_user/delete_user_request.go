package contract

import "identity-api/Domain/models"

type DeleteUserRequest struct {
	ID              string `json:"id" binding:"required" example:"12345" swaggertype:"string"`
	models.JwtToken `json:",inline"`
}
