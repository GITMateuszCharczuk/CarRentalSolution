package queries

import models "identity-api/Domain/models/token"

type GetAllUsersQuery struct {
	models.JwtToken `json:",inline"`
}
