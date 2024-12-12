package contract

import (
	models "identity-api/Domain/models/user"
	"identity-api/Domain/responses"
)

type GetUserInfoResponse struct { //TODO
	responses.BaseResponse
	UserSecureInfo models.UserSecureInfo `json:"user_info" swaggertype:"object"`
}

type GetUserInfoResponse200 struct {
	Success        bool                  `json:"success" example:"true" swaggertype:"boolean"`
	Message        string                `json:"message" example:"User info retrieved successfully" swaggertype:"string"`
	UserSecureInfo models.UserSecureInfo `json:"user_info" swaggertype:"object"`
}

type GetUserInfoResponse400 struct {
	Success        bool                  `json:"success" example:"false" swaggertype:"boolean"`
	Message        string                `json:"message" example:"Invalid request parameters" swaggertype:"string"`
	UserSecureInfo models.UserSecureInfo `json:"user_info" swaggertype:"object"`
}

type GetUserInfoResponse401 struct {
	Success        bool                  `json:"success" example:"false" swaggertype:"boolean"`
	Message        string                `json:"message" example:"Invalid or expired token" swaggertype:"string"`
	UserSecureInfo models.UserSecureInfo `json:"user_info" swaggertype:"object"`
}

type GetUserInfoResponse404 struct {
	Success        bool                  `json:"success" example:"false" swaggertype:"boolean"`
	Message        string                `json:"message" example:"User not found" swaggertype:"string"`
	UserSecureInfo models.UserSecureInfo `json:"user_info" swaggertype:"object"`
}

type GetUserInfoResponse500 struct {
	Success        bool                  `json:"success" example:"false" swaggertype:"boolean"`
	Message        string                `json:"message" example:"Internal server error during user info retrieval" swaggertype:"string"`
	UserSecureInfo models.UserSecureInfo `json:"user_info" swaggertype:"object"`
}
