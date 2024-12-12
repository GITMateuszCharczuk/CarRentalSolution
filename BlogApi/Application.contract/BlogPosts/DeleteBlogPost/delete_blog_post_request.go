package contract

import models "identity-api/Domain/models/token"

type DeleteBlogPostRequest struct {
	BlogPostId      string `json:"blogPostId" validate:"required"`
	models.JwtToken `json:",inline"`
}
