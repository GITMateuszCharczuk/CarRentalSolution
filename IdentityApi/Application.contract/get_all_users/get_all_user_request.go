package contract

import "identity-api/Domain/models"

type GetAllUsersRequest struct {
	models.JwtToken `json:",inline"`
}
