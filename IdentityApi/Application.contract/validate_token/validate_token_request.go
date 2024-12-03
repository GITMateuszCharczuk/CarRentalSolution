package contract

import (
	models "identity-api/Domain/models/token"
)

type ValidateTokenRequest struct {
	models.JwtToken `json:",inline"`
}
