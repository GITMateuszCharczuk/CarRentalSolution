package contract

import (
	"identity-api/Domain/models"
	"identity-api/Domain/responses"
)

type GetAllUsersResponse struct {
	responses.BaseResponse
	Users []models.UserSecureInfo `json:"users" swaggertype:"array"`
}
