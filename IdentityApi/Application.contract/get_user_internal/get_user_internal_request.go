package contract

import (
	models "identity-api/Domain/models/token"
)

type GetUserInternalRequest struct {
	models.JwtToken `json:",inline"`
}
