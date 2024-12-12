package models

type BlogPostCommentModel struct {
	Id          string `json:"id" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string"`
	Description string `json:"description" example:"This is a great blog post!" swaggertype:"string"`
	BlogPostId  string `json:"blogPostId" example:"456e4567-e89b-12d3-a456-426614174000" swaggertype:"string"`
	UserId      string `json:"userId" example:"789e4567-e89b-12d3-a456-426614174000" swaggertype:"string"`
	CreatedAt   string `json:"createdAt" example:"2023-12-12T10:00:00Z" swaggertype:"string"`
	UpdatedAt   string `json:"updatedAt" example:"2023-12-12T10:00:00Z" swaggertype:"string"`
}
