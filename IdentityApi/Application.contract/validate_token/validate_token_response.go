package contract

import "identity-api/Domain/responses"

type ValidateTokenResponse struct {
	responses.BaseResponse
	Roles []string `json:"roles" example:"[user, admin]" swaggertype:"array,string"`
}
