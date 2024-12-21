package contract

import models "email-service/Domain/models/external"

type GetEmailRequest struct {
	ID              string `json:"id" binding:"required" validate:"required,len=60"`
	models.JwtToken `json:"-"`
}
