package commands

import models "identity-api/Domain/models/token"

type ModifyUserCommand struct {
	UserID          string   `json:"user_id"`
	Name            string   `json:"name" binding:"required" example:"John"`
	Surname         string   `json:"surname" binding:"required" example:"Doe"`
	PhoneNumber     string   `json:"phone_number" binding:"required" example:"+1234567890"`
	Address         string   `json:"address" binding:"required" example:"123 Main St"`
	PostalCode      string   `json:"postal_code" binding:"required" example:"12345"`
	City            string   `json:"city" binding:"required" example:"New York"`
	Roles           []string `json:"roles" binding:"required" example:"admin"`
	models.JwtToken `json:",inline"`
}
