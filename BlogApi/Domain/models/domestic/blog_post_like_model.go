package models

type BlogPostLikeModel struct {
	Id         string `json:"id" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string"`
	BlogPostId string `json:"blogPostId" example:"456e4567-e89b-12d3-a456-426614174000" swaggertype:"string"`
	UserId     string `json:"userId" example:"789e4567-e89b-12d3-a456-426614174000" swaggertype:"string"`
	CreatedAt  string `json:"createdAt" example:"2023-12-12T10:00:00Z" swaggertype:"string"`
}
