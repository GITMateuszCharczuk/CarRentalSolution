package commands

import (
	models "identity-api/Domain/models/token"
)

type ValidateTokenCommand struct {
	models.JwtToken `json:",inline"`
}
