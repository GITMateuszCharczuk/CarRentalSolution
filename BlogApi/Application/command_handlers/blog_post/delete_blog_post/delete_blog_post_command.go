package commands

import (
	models "identity-api/Domain/models/external"
)

type DeleteBlogPostCommand struct {
	ID string `json:"id" validate:"required"`
	models.JwtToken
}
