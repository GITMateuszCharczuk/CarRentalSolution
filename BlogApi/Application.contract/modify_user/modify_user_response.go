package contract

import "identity-api/Domain/responses"

type ModifyUserResponse struct {
	responses.BaseResponse
}

type ModifyUserResponse200 struct {
	Success bool   `json:"success" example:"true" swaggertype:"boolean"`
	Message string `json:"message" example:"User modified successfully" swaggertype:"string"`
}

type ModifyUserResponse400 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Invalid request parameters" swaggertype:"string"`
}

type ModifyUserResponse401 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Invalid or expired token" swaggertype:"string"`
}

type ModifyUserResponse404 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"User not found" swaggertype:"string"`
}

type ModifyUserResponse500 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Internal server error during user modification" swaggertype:"string"`
}
