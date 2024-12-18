package services

import (
	"blog-api/Domain/responses"
	"fmt"
	"log"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateRequest[TResponse any](req interface{}, validatorInstance *validator.Validate) *TResponse {
	var errorMessage string
	if err := validatorInstance.Struct(req); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			var errorMessages []string
			for _, fieldError := range validationErrors {
				errorMessages = append(errorMessages, fmt.Sprintf("Field '%s' failed validation: %s", fieldError.Field(), fieldError.Tag()))
			}
			errorMessage = strings.Join(errorMessages, ", ")
		} else {
			errorMessage = "Invalid request parameters"
		}
		response := responses.NewResponse[TResponse](400, errorMessage)
		log.Println(response)
		return &response
	}
	return nil
}
