package models

type UserInfo struct {
	Name         string `json:"name" example:"John" swaggertype:"string"`
	Surname      string `json:"surname" example:"Doe" swaggertype:"string"`
	PhoneNumber  string `json:"phone_number" example:"+1234567890" swaggertype:"string"`
	EmailAddress string `json:"email_address" example:"user@example.com" swaggertype:"string"`
	Address      string `json:"address" example:"123 Main St" swaggertype:"string"`
	PostalCode   string `json:"postal_code" example:"12345" swaggertype:"string"`
	City         string `json:"city" example:"New York" swaggertype:"string"`
}
