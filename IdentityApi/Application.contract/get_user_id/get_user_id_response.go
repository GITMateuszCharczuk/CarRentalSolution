package contract

import "identity-api/Domain/responses"

type GetUserIDResponse struct {
	responses.BaseResponse
	UserID string   `json:"user_id" example:"12345" swaggertype:"string"`
	Roles  []string `json:"roles" example:"[user, admin]" swaggertype:"array,string"`
}
