package contract

type RegisterUserRequest struct {
	Name         string `json:"name" binding:"required" example:"John" swaggertype:"string" validate:"required"`
	Surname      string `json:"surname" binding:"required" example:"Doe" swaggertype:"string" validate:"required"`
	PhoneNumber  string `json:"phone_number" binding:"required" example:"+1234567890" swaggertype:"string" validate:"required"`
	EmailAddress string `json:"email_address" binding:"required,email" example:"user@example.com" swaggertype:"string" validate:"required,email"`
	Address      string `json:"address" binding:"required" example:"123 Main St" swaggertype:"string" validate:"required"`
	PostalCode   string `json:"postal_code" binding:"required" example:"12345" swaggertype:"string" validate:"required"`
	City         string `json:"city" binding:"required" example:"New York" swaggertype:"string" validate:"required"`
	Password     string `json:"password" binding:"required" example:"password123" swaggertype:"string" validate:"required,min=8"`
	Role         string `json:"role" example:"user" swaggertype:"string" validate:"oneof=user admin superadmin"`
}
