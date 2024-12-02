package contract

import (
	"identity-api/Domain/models"
	"identity-api/Domain/responses"
)

type GetUserInfoResponse struct {
	responses.BaseResponse
	models.UserSecureInfo
}
