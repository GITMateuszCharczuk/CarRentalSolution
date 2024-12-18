package commands

import (
	models "blog-api/Domain/models/external"
)

type CreateBlogPostCommand struct {
	Heading          string   `json:"heading" binding:"required" example:"Blog Post Title" swaggertype:"string" validate:"required"`
	PageTitle        string   `json:"pageTitle" binding:"required" example:"Page Title" swaggertype:"string" validate:"required"`
	Content          string   `json:"content" binding:"required" example:"Blog post content..." swaggertype:"string" validate:"required"`
	ShortDescription string   `json:"shortDescription" binding:"required" example:"Short description" swaggertype:"string" validate:"required"`
	FeaturedImageUrl string   `json:"featuredImageUrl" example:"https://example.com/image.jpg" swaggertype:"string"`
	UrlHandle        string   `json:"urlHandle" binding:"required" example:"blog-post-title" swaggertype:"string" validate:"required"`
	Visible          bool     `json:"visible" example:"true" swaggertype:"boolean"`
	Tags             []string `json:"tags" example:"[\"tech\",\"programming\"]" swaggertype:"array,string"`
	models.JwtToken  `json:",inline"`
}
