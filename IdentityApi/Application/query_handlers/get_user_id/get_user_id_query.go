package queries

import "identity-api/Domain/models"

type GetUserIDQuery struct {
	models.JwtToken `json:",inline"`
}
