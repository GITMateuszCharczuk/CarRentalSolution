package contract

import (
	responses "identity-api/Domain/responses"
)

type CreateLikeForBlogPostResponse struct {
	responses.BaseResponse
	Id string `json:"id" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string"`
}

type CreateLikeForBlogPostResponse200 struct {
	Success bool   `json:"success" example:"true" swaggertype:"boolean"`
	Message string `json:"message" example:"Like added successfully" swaggertype:"string"`
	Id      string `json:"id" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string"`
}

type CreateLikeForBlogPostResponse400 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Invalid request parameters" swaggertype:"string"`
	Id      string `json:"id" example:"" swaggertype:"string"`
}

type CreateLikeForBlogPostResponse401 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Unauthorized" swaggertype:"string"`
	Id      string `json:"id" example:"" swaggertype:"string"`
}

type CreateLikeForBlogPostResponse404 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Blog post not found" swaggertype:"string"`
	Id      string `json:"id" example:"" swaggertype:"string"`
}

type CreateLikeForBlogPostResponse500 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Internal server error while adding like" swaggertype:"string"`
	Id      string `json:"id" example:"" swaggertype:"string"`
}
