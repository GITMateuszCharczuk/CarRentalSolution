package commands

type SendEmailCommand struct {
	From    string `json:"from" binding:"required,email"`
	To      string `json:"to" binding:"required,email"`
	Subject string `json:"subject" binding:"required"`
	Body    string `json:"body" binding:"required"`
}
