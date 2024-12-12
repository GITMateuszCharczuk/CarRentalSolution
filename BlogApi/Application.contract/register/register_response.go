package contract

import "identity-api/Domain/responses"

type RegisterUserResponse struct {
	responses.BaseResponse
}

type RegisterUserResponse201 struct {
	Success bool   `json:"success" example:"true" swaggertype:"boolean"`
	Message string `json:"message" example:"User registered successfully" swaggertype:"string"`
	UserID  string `json:"user_id" example:"12345" swaggertype:"string"`
}

type RegisterUserResponse400 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Invalid request parameters" swaggertype:"string"`
}

type RegisterUserResponse409 struct { //TODO
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"User with this email already exists" swaggertype:"string"`
}

type RegisterUserResponse500 struct {
	Success bool   `json:"success" example:"false" swaggertype:"boolean"`
	Message string `json:"message" example:"Internal server error during registration" swaggertype:"string"`
}
