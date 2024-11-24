package contract

type SendEmailRequest struct {
	From    string `json:"from" binding:"required,email" example:"test@test.com" swaggertype:"string"`
	To      string `json:"to" binding:"required,email" example:"recipient@example.com" swaggertype:"string"`
	Subject string `json:"subject" binding:"required" example:"Hello" swaggertype:"string"`
	Body    string `json:"body" binding:"required" example:"This is the body of the email." swaggertype:"string"`
}
