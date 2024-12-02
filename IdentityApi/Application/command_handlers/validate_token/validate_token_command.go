package commands

import "identity-api/Domain/models"

type ValidateTokenCommand struct {
	models.JwtToken `json:",inline"`
}
