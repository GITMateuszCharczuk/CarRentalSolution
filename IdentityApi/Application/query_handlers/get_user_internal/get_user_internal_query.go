package queries

import models "identity-api/Domain/models/token"

type GetUserInternalQuery struct {
	models.JwtToken `json:",inline"`
}
