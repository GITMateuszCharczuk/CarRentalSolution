package contract

import (
	models "identity-api/Domain/models/token"
	responses "identity-api/Domain/responses"
)

type LoginResponse struct {
	responses.BaseResponse
	models.JwtToken        `json:",inline"`
	models.JwtRefreshToken `json:",inline"`
	Roles                  []string `json:"roles" example:"[user, admin]" swaggertype:"array,string"`
}

type LoginResponse200 struct {
	Success      bool     `json:"success" example:"true" swaggertype:"boolean"`
	Message      string   `json:"message" example:"Login successful" swaggertype:"string"`
	Token        string   `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." swaggertype:"string"`
	RefreshToken string   `json:"refresh_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." swaggertype:"string"`
	Roles        []string `json:"roles" example:"[user, admin]" swaggertype:"array,string"`
}

type LoginResponse400 struct {
	Success      bool     `json:"success" example:"false" swaggertype:"boolean"`
	Message      string   `json:"message" example:"Invalid request parameters" swaggertype:"string"`
	Token        string   `json:"token" example:"" swaggertype:"string"`
	RefreshToken string   `json:"refresh_token" example:"" swaggertype:"string"`
	Roles        []string `json:"roles" example:"[]" swaggertype:"array,string"`
}

type LoginResponse401 struct {
	Success      bool     `json:"success" example:"false" swaggertype:"boolean"`
	Message      string   `json:"message" example:"Invalid credentials" swaggertype:"string"`
	Token        string   `json:"token" example:"" swaggertype:"string"`
	RefreshToken string   `json:"refresh_token" example:"" swaggertype:"string"`
	Roles        []string `json:"roles" example:"[]" swaggertype:"array,string"`
}

type LoginResponse500 struct {
	Success      bool     `json:"success" example:"false" swaggertype:"boolean"`
	Message      string   `json:"message" example:"Internal server error during login" swaggertype:"string"`
	Token        string   `json:"token" example:"" swaggertype:"string"`
	RefreshToken string   `json:"refresh_token" example:"" swaggertype:"string"`
	Roles        []string `json:"roles" example:"[]" swaggertype:"array,string"`
}
