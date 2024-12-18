package contract

import (
	responses "blog-api/Domain/responses"
)

type DeleteLikeForBlogPostResponse struct {
	responses.BaseResponse
}

type DeleteLikeForBlogPostResponse200 struct {
	Success bool   `json:"success" example:"true" swaggertype:"boolean"`
	Message string `json:"message" example:"Like removed successfully" swaggertype:"string"`
}

type DeleteLikeForBlogPostResponse400 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Invalid request parameters" swaggertype:"string"`
}

type DeleteLikeForBlogPostResponse401 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Unauthorized" swaggertype:"string"`
}

type DeleteLikeForBlogPostResponse404 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Like not found" swaggertype:"string"`
}

type DeleteLikeForBlogPostResponse500 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Internal server error while removing like" swaggertype:"string"`
}
