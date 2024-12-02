package queries

import "identity-api/Domain/models"

type GetUserInfoQuery struct {
	Id              string `json:"id"`
	models.JwtToken `json:",inline"`
}
