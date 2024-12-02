package commands

type RegisterUserCommand struct {
	Name         string `json:"name" binding:"required"`
	Surname      string `json:"surname" binding:"required"`
	PhoneNumber  string `json:"phone_number" binding:"required"`
	EmailAddress string `json:"email_address" binding:"required,email"`
	Address      string `json:"address" binding:"required"`
	PostalCode   string `json:"postal_code" binding:"required"`
	City         string `json:"city" binding:"required"`
	Password     string `json:"password" binding:"required"`
	Role         string `json:"role"`
}
