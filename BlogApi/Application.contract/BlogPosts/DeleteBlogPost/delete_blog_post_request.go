package contract

import models "blog-api/Domain/models/external"

type DeleteBlogPostRequest struct {
	BlogPostId      string `json:"blogPostId" validate:"required"`
	models.JwtToken `json:",inline"`
}
