package contract

import "identity-api/Domain/responses"

type DeleteUserResponse struct {
	responses.BaseResponse
}

type DeleteUserResponse200 struct {
	Success bool   `json:"success" example:"true" swaggertype:"boolean"`
	Message string `json:"message" example:"User deleted successfully" swaggertype:"string"`
}

type DeleteUserResponse400 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Invalid request parameters" swaggertype:"string"`
}

type DeleteUserResponse401 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Invalid or expired token" swaggertype:"string"`
}

type DeleteUserResponse404 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"User not found" swaggertype:"string"`
}

type DeleteUserResponse500 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Internal server error during user deletion" swaggertype:"string"`
}
