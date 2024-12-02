package commands

import "identity-api/Domain/models"

type DeleteUserCommand struct {
	models.JwtToken `json:",inline"`
}
