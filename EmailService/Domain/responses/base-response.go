package responses

import "email-service/Domain/constants"

type BaseResponse struct {
	Success    bool                 `json:"success"`
	Message    string               `json:"message"`
	StatusCode constants.StatusCode `json:"status_code"`
}

func NewBaseResponse(statusCode int, customMessage ...string) BaseResponse {
	var message string

	switch constants.StatusCode(statusCode) {
	case constants.StatusOK:
		message = "Operation successful"
	case constants.StatusCreated:
		message = "Resource created successfully"
	case constants.StatusBadRequest:
		message = "Bad request"
	case constants.StatusNotFound:
		message = "Resource not found"
	case constants.StatusConflict:
		message = "Resource already exists"
	case constants.StatusUnauthorized:
		message = "Unauthorized"
	case constants.StatusForbidden:
		message = "Forbidden"
	case constants.StatusInternalServerError:
		message = "Internal server error"
	default:
		message = "Unknown status"
	}

	if len(customMessage) > 0 {
		message = customMessage[0]
	}

	return BaseResponse{
		Success:    statusCode == int(constants.StatusOK) || statusCode == int(constants.StatusCreated),
		Message:    message,
		StatusCode: constants.StatusCode(statusCode),
	}
}
