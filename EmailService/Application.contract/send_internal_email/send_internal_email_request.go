package contract

type SendInternalEmailRequest struct {
	To      string `json:"to" binding:"required,email" example:"recipient@example.com" swaggertype:"string" validate:"required,email"`
	Subject string `json:"subject" binding:"required" example:"Hello" swaggertype:"string" validate:"required,min=5,max=100"`
	Body    string `json:"body" binding:"required" example:"This is the body of the email." swaggertype:"string" validate:"required,min=5,max=1000"`
}
