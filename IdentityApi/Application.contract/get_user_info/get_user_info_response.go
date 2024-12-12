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
	Success        bool   `json:"success" example:"true" swaggertype:"boolean"`
	Message        string `json:"message" example:"User info retrieved successfully" swaggertype:"string"`
	UserSecureInfo struct {
		ID           string   `json:"id" example:"12345"`
		Name         string   `json:"name" example:"John"`
		Surname      string   `json:"surname" example:"Doe"`
		PhoneNumber  string   `json:"phone_number" example:"+1234567890"`
		EmailAddress string   `json:"email_address" example:"user@example.com"`
		Address      string   `json:"address" example:"123 Main St"`
		PostalCode   string   `json:"postal_code" example:"12345"`
		City         string   `json:"city" example:"New York"`
		Roles        []string `json:"roles" example:"user,admin"`
	} `json:"user_info"`
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
