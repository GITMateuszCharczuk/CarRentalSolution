package contract

import (
	responses "identity-api/Domain/responses"
)

type UpdateBlogPostResponse struct {
	responses.BaseResponse
	Title string `json:"title" example:"Blog Post Updated" swaggertype:"string"`
}

type UpdateBlogPostResponse200 struct {
	Success bool   `json:"success" example:"true" swaggertype:"boolean"`
	Message string `json:"message" example:"Blog post updated successfully" swaggertype:"string"`
	Title   string `json:"title" example:"Blog Post Updated" swaggertype:"string"`
}

type UpdateBlogPostResponse400 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Invalid request parameters" swaggertype:"string"`
	Title   string `json:"title" example:"" swaggertype:"string"`
}

type UpdateBlogPostResponse404 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Blog post not found" swaggertype:"string"`
	Title   string `json:"title" example:"" swaggertype:"string"`
}

type UpdateBlogPostResponse500 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Internal server error while updating blog post" swaggertype:"string"`
	Title   string `json:"title" example:"" swaggertype:"string"`
}
