package commands

import models "identity-api/Domain/models/token"

type DeleteUserCommand struct {
	ID              string `json:"id"`
	models.JwtToken `json:",inline"`
}
