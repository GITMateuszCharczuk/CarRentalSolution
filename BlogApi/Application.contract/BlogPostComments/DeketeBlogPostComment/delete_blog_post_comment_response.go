package contract

import (
	responses "blog-api/Domain/responses"
)

type DeleteBlogPostCommentResponse struct {
	responses.BaseResponse
}

type DeleteBlogPostCommentResponse200 struct {
	Success bool   `json:"success" example:"true" swaggertype:"boolean"`
	Message string `json:"message" example:"Comment deleted successfully" swaggertype:"string"`
}

type DeleteBlogPostCommentResponse400 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Invalid request parameters" swaggertype:"string"`
}

type DeleteBlogPostCommentResponse401 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Unauthorized" swaggertype:"string"`
}

type DeleteBlogPostCommentResponse403 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Forbidden" swaggertype:"string"`
}

type DeleteBlogPostCommentResponse404 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Comment not found" swaggertype:"string"`
}

type DeleteBlogPostCommentResponse500 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Internal server error while deleting comment" swaggertype:"string"`
}
