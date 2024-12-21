package services

import (
	"fmt"
	"rental-api/Domain/responses"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateRequest[TResponse any](req interface{}, validatorInstance *validator.Validate) *TResponse {
	var errorMessage string
	if err := validatorInstance.Struct(req); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			var errorMessages []string
			for _, e := range validationErrors {
				switch e.Tag() {
				case "datetime":
					errorMessages = append(errorMessages,
						fmt.Sprintf("Field %s must be in format: YYYY-MM-DD HH:MM:SS.NNNNNN+ZZ", e.Field()))
				case "futuredate":
					errorMessages = append(errorMessages,
						fmt.Sprintf("Field %s must be a future date", e.Field()))
				case "gtdate":
					errorMessages = append(errorMessages,
						fmt.Sprintf("Field %s must be after %s", e.Field(), e.Param()))
				default:
					errorMessages = append(errorMessages, e.Error())
				}
			}
			response := responses.NewResponse[TResponse](400, strings.Join(errorMessages, "; "))
			return &response
		}
		response := responses.NewResponse[TResponse](400, errorMessage)
		return &response
	}
	return nil
}
