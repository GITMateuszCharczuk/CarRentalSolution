package queries

import "identity-api/Domain/models"

type GetAllUsersQuery struct {
	models.JwtToken `json:",inline"`
}
