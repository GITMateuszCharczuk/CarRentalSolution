package contract

import (
	responses "identity-api/Domain/responses"
)

type DeleteBlogPostResponse struct {
	responses.BaseResponse
}

type DeleteBlogPostResponse200 struct {
	Success bool   `json:"success" example:"true" swaggertype:"boolean"`
	Message string `json:"message" example:"Blog post deleted successfully" swaggertype:"string"`
}

type DeleteBlogPostResponse404 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Blog post not found" swaggertype:"string"`
}

type DeleteBlogPostResponse500 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Internal server error while deleting blog post" swaggertype:"string"`
}
