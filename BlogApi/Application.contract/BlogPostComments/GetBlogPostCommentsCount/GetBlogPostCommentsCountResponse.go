package contract

import "blog-api/Domain/responses"

type GetBlogPostCommentsCountResponse struct {
	responses.BaseResponse
	Count int `json:"count"`
}

type GetBlogPostCommentsCountResponse200 struct {
	Success bool   `json:"success" example:"true" swaggertype:"boolean"`
	Message string `json:"message" example:"Comments count retrieved successfully" swaggertype:"string"`
	Count   int    `json:"count" example:"100"`
}

type GetBlogPostCommentsCountResponse404 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Blog post not found" swaggertype:"string"`
}

type GetBlogPostCommentsCountResponse500 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Server error during retrieval" swaggertype:"string"`
}

type GetBlogPostCommentsCountResponse400 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Invalid request parameters" swaggertype:"string"`
}
