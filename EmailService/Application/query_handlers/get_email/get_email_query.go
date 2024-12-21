package queries

import models "email-service/Domain/models/external"

type GetEmailQuery struct {
	ID       string
	JwtToken models.JwtToken
}
