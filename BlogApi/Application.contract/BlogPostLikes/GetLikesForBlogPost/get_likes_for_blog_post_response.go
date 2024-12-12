package contract

import (
	responses "identity-api/Domain/responses"
)

type GetLikesForBlogPostResponse struct {
	responses.BaseResponse
	TotalCount int `json:"totalCount" example:"42" swaggertype:"integer"`
}

type GetLikesForBlogPostResponse200 struct {
	Success    bool   `json:"success" example:"true" swaggertype:"boolean"`
	Message    string `json:"message" example:"Likes retrieved successfully" swaggertype:"string"`
	TotalCount int    `json:"totalCount" example:"42" swaggertype:"integer"`
}

type GetLikesForBlogPostResponse404 struct {
	Success    bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message    string `json:"message" example:"Blog post not found" swaggertype:"string"`
	TotalCount int    `json:"totalCount" example:"0" swaggertype:"integer"`
}

type GetLikesForBlogPostResponse500 struct {
	Success    bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message    string `json:"message" example:"Internal server error while retrieving likes" swaggertype:"string"`
	TotalCount int    `json:"totalCount" example:"0" swaggertype:"integer"`
}
