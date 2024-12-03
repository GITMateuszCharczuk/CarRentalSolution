package contract

import (
	models "identity-api/Domain/models/token"
)

type GetUserIDRequest struct {
	models.JwtToken `json:",inline"`
}
