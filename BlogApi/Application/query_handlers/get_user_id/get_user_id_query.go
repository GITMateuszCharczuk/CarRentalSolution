package queries

import models "identity-api/Domain/models/token"

type GetUserIDQuery struct {
	models.JwtToken `json:",inline"`
}
