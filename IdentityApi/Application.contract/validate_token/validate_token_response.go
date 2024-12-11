package contract

import "identity-api/Domain/responses"

type ValidateTokenResponse struct {
	responses.BaseResponse
	Roles []string `json:"roles" example:"[user, admin]" swaggertype:"array,string"`
	Valid bool     `json:"valid" example:"true" swaggertype:"boolean"`
}

type ValidateTokenResponse200 struct {
	Success bool     `json:"success" example:"true" swaggertype:"boolean"`
	Message string   `json:"message" example:"Token validated successfully" swaggertype:"string"`
	Roles   []string `json:"roles" example:"[user, admin]" swaggertype:"array,string"`
	Valid   bool     `json:"valid" example:"true" swaggertype:"boolean"`
}

type ValidateTokenResponse400 struct {
	Success bool     `json:"success" example:"false" swaggertype:"boolean"`
	Message string   `json:"message" example:"Invalid request parameters" swaggertype:"string"`
	Roles   []string `json:"roles" example:"[]" swaggertype:"array,string"`
	Valid   bool     `json:"valid" example:"false" swaggertype:"boolean"`
}

type ValidateTokenResponse401 struct {
	Success bool     `json:"success" example:"false" swaggertype:"boolean"`
	Message string   `json:"message" example:"Invalid or expired token" swaggertype:"string"`
	Roles   []string `json:"roles" example:"[]" swaggertype:"array,string"`
	Valid   bool     `json:"valid" example:"false" swaggertype:"boolean"`
}

type ValidateTokenResponse500 struct {
	Success bool     `json:"success" example:"false" swaggertype:"boolean"`
	Message string   `json:"message" example:"Internal server error during token validation" swaggertype:"string"`
	Roles   []string `json:"roles" example:"[]" swaggertype:"array,string"`
	Valid   bool     `json:"valid" example:"false" swaggertype:"boolean"`
}
