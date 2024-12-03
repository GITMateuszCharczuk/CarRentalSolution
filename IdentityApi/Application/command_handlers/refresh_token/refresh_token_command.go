package commands

import models "identity-api/Domain/models/token"

type RefreshTokenCommand struct {
	models.JwtRefreshToken `json:",inline"`
}
