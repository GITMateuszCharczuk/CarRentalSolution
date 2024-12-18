package commands

import (
	models "blog-api/Domain/models/external"
)

type DeleteBlogPostCommand struct {
	ID string `json:"id" validate:"required"`
	models.JwtToken
}
