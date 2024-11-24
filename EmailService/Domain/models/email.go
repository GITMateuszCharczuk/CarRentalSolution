package models

type Email struct {
	ID      string `json:"id" swaggertype:"string" example:"12345"`
	From    string `json:"from" swaggertype:"string" example:"example@example.com"`
	To      string `json:"to" swaggertype:"string" example:"recipient@example.com"`
	Subject string `json:"subject" swaggertype:"string" example:"Hello World"`
	Body    string `json:"body" swaggertype:"string" example:"This is the body of the email."`
}
