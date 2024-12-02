package contract

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email" example:"user@example.com" swaggertype:"string" validate:"required,email"`
	Password string `json:"password" binding:"required" example:"password123" swaggertype:"string" validate:"required,min=8"`
}
