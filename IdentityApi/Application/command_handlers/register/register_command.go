package commands

type RegisterUserCommand struct {
	Name         string `json:"name" binding:"required" example:"John"`
	Surname      string `json:"surname" binding:"required" example:"Doe"`
	PhoneNumber  string `json:"phone_number" binding:"required" example:"+1234567890"`
	EmailAddress string `json:"email_address" binding:"required" example:"user@example.com"`
	Password     string `json:"password" binding:"required" example:"password123"`
	Address      string `json:"address" binding:"required" example:"123 Main St"`
	PostalCode   string `json:"postal_code" binding:"required" example:"12345"`
	City         string `json:"city" binding:"required" example:"New York"`
}
