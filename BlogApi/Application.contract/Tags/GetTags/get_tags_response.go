package contract

import (
	models "identity-api/Domain/models/domestic"
	responses "identity-api/Domain/responses"
)

type GetTagsResponse struct {
	responses.BaseResponse
	Items []models.BlogPostTagModel `json:"items" swaggertype:"array,object"`
}

type GetTagsResponse200 struct {
	Success bool                      `json:"success" example:"true" swaggertype:"boolean"`
	Message string                    `json:"message" example:"Tags retrieved successfully" swaggertype:"string"`
	Items   []models.BlogPostTagModel `json:"items" swaggertype:"array,object"`
}

type GetTagsResponse400 struct {
	Success bool                      `json:"success" example:"false" swaggertype:"boolean"`
	Message string                    `json:"message" example:"Invalid request parameters" swaggertype:"string"`
	Items   []models.BlogPostTagModel `json:"items" swaggertype:"array,object"`
}

type GetTagsResponse404 struct {
	Success bool                      `json:"success" example:"false" swaggertype:"boolean"`
	Message string                    `json:"message" example:"Blog post not found" swaggertype:"string"`
	Items   []models.BlogPostTagModel `json:"items" swaggertype:"array,object"`
}

type GetTagsResponse500 struct {
	Success bool                      `json:"success" example:"false" swaggertype:"boolean"`
	Message string                    `json:"message" example:"Internal server error while retrieving tags" swaggertype:"string"`
	Items   []models.BlogPostTagModel `json:"items" swaggertype:"array,object"`
}
