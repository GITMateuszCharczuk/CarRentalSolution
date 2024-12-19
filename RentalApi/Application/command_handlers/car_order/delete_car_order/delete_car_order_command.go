package commands

import (
	models "rental-api/Domain/models/external"
)

type DeleteCarOrderCommand struct {
	ID string `json:"id" validate:"required"`
	models.JwtToken
}
