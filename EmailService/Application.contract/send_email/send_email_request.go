package contract

import models "email-service/Domain/models/external"

type SendEmailRequest struct {
	From            string `json:"from" binding:"required,email" example:"test@test.com" swaggertype:"string" validate:"required,email"`
	Subject         string `json:"subject" binding:"required" example:"Hello" swaggertype:"string" validate:"required,min=5,max=100"`
	Body            string `json:"body" binding:"required" example:"This is the body of the email." swaggertype:"string" validate:"required,min=5,max=1000"`
	models.JwtToken `json:"-"`
}
