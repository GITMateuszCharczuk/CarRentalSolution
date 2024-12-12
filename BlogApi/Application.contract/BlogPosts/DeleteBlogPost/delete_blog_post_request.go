package contract

import models "identity-api/Domain/models/external"

type DeleteBlogPostRequest struct {
	BlogPostId      string `json:"blogPostId" validate:"required"`
	models.JwtToken `json:",inline"`
}
