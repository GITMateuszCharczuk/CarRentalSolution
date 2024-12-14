package contract

import models "identity-api/Domain/models/external"

type UpdateBlogPostRequest struct {
	Id               string   `json:"id" binding:"required" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string" validate:"required"`
	Heading          string   `json:"heading" binding:"required" example:"Updated Blog Post Title" swaggertype:"string" validate:"required"`
	PageTitle        string   `json:"pageTitle" binding:"required" example:"Updated Page Title" swaggertype:"string" validate:"required"`
	Content          string   `json:"content" binding:"required" example:"Updated blog post content..." swaggertype:"string" validate:"required"`
	ShortDescription string   `json:"shortDescription" binding:"required" example:"Updated short description" swaggertype:"string" validate:"required"`
	FeaturedImageUrl string   `json:"featuredImageUrl" example:"https://example.com/updated-image.jpg" swaggertype:"string"`
	UrlHandle        string   `json:"urlHandle" binding:"required" example:"updated-blog-post-title" swaggertype:"string" validate:"required"`
	PublishedDate    string   `json:"publishedDate" example:"2023-12-12" swaggertype:"string" validate:"required"`
	Author           string   `json:"author" binding:"required" example:"John Doe" swaggertype:"string" validate:"required"`
	Visible          bool     `json:"visible" example:"true" swaggertype:"boolean"`
	Tags             []string `json:"tags" example:"[\"tech\",\"programming\"]" swaggertype:"array,string"`
	models.JwtToken  `json:"jwtToken" validate:"required"`
}
