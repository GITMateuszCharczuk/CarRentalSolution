package commands

import (
	models "rental-api/Domain/models/external"
)

type DeleteCarOfferCommand struct {
	ID string `json:"id" validate:"required"`
	models.JwtToken
}
