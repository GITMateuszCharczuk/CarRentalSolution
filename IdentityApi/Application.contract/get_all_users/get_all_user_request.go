package contract

import models "identity-api/Domain/models/token"

type GetAllUsersRequest struct {
	models.JwtToken `json:",inline"`
}
