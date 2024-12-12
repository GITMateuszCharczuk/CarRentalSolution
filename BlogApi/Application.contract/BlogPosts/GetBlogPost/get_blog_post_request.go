package contract

import models "identity-api/Domain/models/external"

type GetBlogPostRequest struct {
	BlogPostId      string `json:"blogPostId" validate:"required"`
	models.JwtToken `json:",inline"`
}
