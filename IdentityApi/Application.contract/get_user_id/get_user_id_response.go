package contract

import "identity-api/Domain/responses"

type GetUserIDResponse struct {
	responses.BaseResponse
	UserID string   `json:"user_id" example:"12345" swaggertype:"string"`
	Roles  []string `json:"roles" example:"[user, admin]" swaggertype:"array,string"`
}

type GetUserIDResponse200 struct {
	Success bool     `json:"success" example:"true" swaggertype:"boolean"`
	Message string   `json:"message" example:"User ID retrieved successfully" swaggertype:"string"`
	Roles   []string `json:"roles" example:"[user, admin]" swaggertype:"array,string"`
	UserID  string   `json:"user_id" example:"12345" swaggertype:"string"`
}

type GetUserIDResponse400 struct {
	Success bool     `json:"success" example:"false" swaggertype:"boolean"`
	Message string   `json:"message" example:"Invalid request parameters" swaggertype:"string"`
	UserID  string   `json:"user_id" example:"" swaggertype:"string"`
	Roles   []string `json:"roles" example:"[]" swaggertype:"array,string"`
}

type GetUserIDResponse401 struct {
	Success bool     `json:"success" example:"false" swaggertype:"boolean"`
	Message string   `json:"message" example:"Invalid or expired token" swaggertype:"string"`
	UserID  string   `json:"user_id" example:"" swaggertype:"string"`
	Roles   []string `json:"roles" example:"[]" swaggertype:"array,string"`
}

type GetUserIDResponse404 struct {
	Success bool     `json:"success" example:"false" swaggertype:"boolean"`
	Message string   `json:"message" example:"User not found" swaggertype:"string"`
	UserID  string   `json:"user_id" example:"" swaggertype:"string"`
	Roles   []string `json:"roles" example:"[]" swaggertype:"array,string"`
}

type GetUserIDResponse500 struct {
	Success bool     `json:"success" example:"false" swaggertype:"boolean"`
	Message string   `json:"message" example:"Internal server error during user ID retrieval" swaggertype:"string"`
	UserID  string   `json:"user_id" example:"" swaggertype:"string"`
	Roles   []string `json:"roles" example:"[]" swaggertype:"array,string"`
}
