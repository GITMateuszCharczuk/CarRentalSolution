package contract

import (
	responses "identity-api/Domain/responses"
)

type TagModel struct {
	Id   string `json:"id" example:"123e4567-e89b-12d3-a456-426614174000" swaggertype:"string"`
	Name string `json:"name" example:"Technology" swaggertype:"string"`
}

type GetTagsResponse struct {
	responses.BaseResponse
	Items []TagModel `json:"items" swaggertype:"array,object"`
}

type GetTagsResponse200 struct {
	Success bool       `json:"success" example:"true" swaggertype:"boolean"`
	Message string     `json:"message" example:"Tags retrieved successfully" swaggertype:"string"`
	Items   []TagModel `json:"items" swaggertype:"array,object"`
}

type GetTagsResponse400 struct {
	Success bool       `json:"success" example:"false" swaggertype:"boolean"`
	Message string     `json:"message" example:"Invalid request parameters" swaggertype:"string"`
	Items   []TagModel `json:"items" swaggertype:"array,object"`
}

type GetTagsResponse404 struct {
	Success bool       `json:"success" example:"false" swaggertype:"boolean"`
	Message string     `json:"message" example:"Blog post not found" swaggertype:"string"`
	Items   []TagModel `json:"items" swaggertype:"array,object"`
}

type GetTagsResponse500 struct {
	Success bool       `json:"success" example:"false" swaggertype:"boolean"`
	Message string     `json:"message" example:"Internal server error while retrieving tags" swaggertype:"string"`
	Items   []TagModel `json:"items" swaggertype:"array,object"`
}
