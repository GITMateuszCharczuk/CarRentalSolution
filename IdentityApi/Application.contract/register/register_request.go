package contract

type RegisterUserRequest struct {
	Name         string `json:"name" validate:"required" example:"John" swaggertype:"string"`
	Surname      string `json:"surname" validate:"required" example:"Doe" swaggertype:"string"`
	PhoneNumber  string `json:"phone_number" validate:"required" example:"+1234567890" swaggertype:"string"`
	EmailAddress string `json:"email_address" validate:"required,email" example:"user@example.com" swaggertype:"string"`
	Address      string `json:"address" validate:"required" example:"123 Main St" swaggertype:"string"`
	PostalCode   string `json:"postal_code" validate:"required" example:"12345" swaggertype:"string"`
	City         string `json:"city" validate:"required" example:"New York" swaggertype:"string"`
	Password     string `json:"password" validate:"required,min=8,containsany=0123456789,containsany=!@#$%^&*(),containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ,containsany=abcdefghijklmnopqrstuvwxyz" example:"P@ssw0rd123" swaggertype:"string"`
}
