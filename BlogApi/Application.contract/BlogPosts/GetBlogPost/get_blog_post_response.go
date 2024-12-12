package contract

import (
	"identity-api/Domain/models"
	responses "identity-api/Domain/responses"
)

type TagModel struct {
	Id   string `json:"id" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string"`
	Name string `json:"name" example:"Technology" swaggertype:"string"`
}

type GetBlogPostResponse struct {
	responses.BaseResponse
	Id               string                        `json:"id" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string"`
	Heading          string                        `json:"heading" example:"Blog Post Title" swaggertype:"string"`
	PageTitle        string                        `json:"pageTitle" example:"Page Title" swaggertype:"string"`
	Content          string                        `json:"content" example:"Blog post content..." swaggertype:"string"`
	ShortDescription string                        `json:"shortDescription" example:"Short description" swaggertype:"string"`
	FeaturedImageUrl string                        `json:"featuredImageUrl" example:"https://example.com/image.jpg" swaggertype:"string"`
	UrlHandle        string                        `json:"urlHandle" example:"blog-post-title" swaggertype:"string"`
	PublishedDate    string                        `json:"publishedDate" example:"2023-12-12" swaggertype:"string"`
	Author           string                        `json:"author" example:"John Doe" swaggertype:"string"`
	Visible          bool                          `json:"visible" example:"true" swaggertype:"boolean"`
	Tags             []TagModel                    `json:"tags" swaggertype:"array,object"`
	Likes            []models.BlogPostLikeModel    `json:"likes" swaggertype:"array,object"`
	Comments         []models.BlogPostCommentModel `json:"comments" swaggertype:"array,object"`
}

type GetBlogPostResponse200 struct {
	Success          bool                          `json:"success" example:"true" swaggertype:"boolean"`
	Message          string                        `json:"message" example:"Blog post retrieved successfully" swaggertype:"string"`
	Id               string                        `json:"id" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string"`
	Heading          string                        `json:"heading" example:"Blog Post Title" swaggertype:"string"`
	PageTitle        string                        `json:"pageTitle" example:"Page Title" swaggertype:"string"`
	Content          string                        `json:"content" example:"Blog post content..." swaggertype:"string"`
	ShortDescription string                        `json:"shortDescription" example:"Short description" swaggertype:"string"`
	FeaturedImageUrl string                        `json:"featuredImageUrl" example:"https://example.com/image.jpg" swaggertype:"string"`
	UrlHandle        string                        `json:"urlHandle" example:"blog-post-title" swaggertype:"string"`
	PublishedDate    string                        `json:"publishedDate" example:"2023-12-12" swaggertype:"string"`
	Author           string                        `json:"author" example:"John Doe" swaggertype:"string"`
	Visible          bool                          `json:"visible" example:"true" swaggertype:"boolean"`
	Tags             []TagModel                    `json:"tags" swaggertype:"array,object"`
	Likes            []models.BlogPostLikeModel    `json:"likes" swaggertype:"array,object"`
	Comments         []models.BlogPostCommentModel `json:"comments" swaggertype:"array,object"`
}

type GetBlogPostResponse404 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Blog post not found" swaggertype:"string"`
}

type GetBlogPostResponse500 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Internal server error while retrieving blog post" swaggertype:"string"`
}
