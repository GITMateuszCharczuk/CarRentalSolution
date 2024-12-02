package contract

import "identity-api/Domain/responses"

type GetUserIDResponse struct {
	responses.BaseResponse
	UserID string `json:"user_id" example:"12345" swaggertype:"string"`
}
