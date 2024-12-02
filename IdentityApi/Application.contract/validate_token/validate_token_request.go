package contract

import "identity-api/Domain/models"

type ValidateTokenRequest struct {
	models.JwtToken `json:",inline"`
}
