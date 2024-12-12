package queries

import (
	models "identity-api/Domain/models/token"
)

type GetUserInfoQuery struct {
	Id              string `json:"id"`
	models.JwtToken `json:",inline"`
}
