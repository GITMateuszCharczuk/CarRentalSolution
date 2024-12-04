package contract

import (
	models "identity-api/Domain/models/token"
)

type ModifyUserRequest struct {
	UserID          string   `json:"user_id,omitempty" example:"1234567890"`
	Name            string   `json:"name,omitempty" example:"John"`
	Surname         string   `json:"surname,omitempty" example:"Doe"`
	PhoneNumber     string   `json:"phone_number,omitempty" example:"+1234567890"`
	EmailAddress    string   `json:"email_address,omitempty" example:"user@example.com"`
	Address         string   `json:"address,omitempty" example:"123 Main St"`
	PostalCode      string   `json:"postal_code,omitempty" example:"12345"`
	City            string   `json:"city,omitempty" example:"New York"`
	Roles           []string `json:"roles,omitempty" example:"[user, admin]"`
	models.JwtToken `json:",inline"`
}
