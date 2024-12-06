package contract

import (
	models "identity-api/Domain/models/token"
)

type ModifyUserRequest struct {
	UserID          string   `json:"user_id" example:"1234567890"`
	Name            string   `json:"name" validate:"required,min=2,max=50" example:"John" swaggertype:"string"`
	Surname         string   `json:"surname" validate:"required,min=2,max=50" example:"Doe" swaggertype:"string"`
	PhoneNumber     string   `json:"phone_number" validate:"required,e164" example:"+1234567890" swaggertype:"string"`
	EmailAddress    string   `json:"email_address" validate:"omitempty,email" example:"user@example.com" swaggertype:"string"`
	Address         string   `json:"address" validate:"required,min=5,max=100" example:"123 Main St" swaggertype:"string"`
	PostalCode      string   `json:"postal_code" validate:"required,min=5,max=10" example:"12345" swaggertype:"string"`
	City            string   `json:"city" validate:"required,min=2,max=50" example:"New York" swaggertype:"string"`
	Roles           []string `json:"roles" validate:"dive,oneof=user admin superadmin" example:"[user,admin]" swaggertype:"array,string"`
	models.JwtToken `json:",inline"`
}
