package commands

import models "email-service/Domain/models/external"

type SendEmailCommand struct {
	From            string `json:"from" binding:"required,email"`
	Subject         string `json:"subject" binding:"required"`
	Body            string `json:"body" binding:"required"`
	models.JwtToken `json:",inline"`
}
