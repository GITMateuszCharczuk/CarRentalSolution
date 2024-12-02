package contract

import "identity-api/Domain/models"

type GetUserIDRequest struct {
	models.JwtToken `json:",inline"`
}
