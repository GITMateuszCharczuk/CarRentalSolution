package services

import (
	"fmt"
	"identity-api/Domain/responses"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateRequest[TResponse any](req interface{}, validatorInstance *validator.Validate) *TResponse {
	if err := validatorInstance.Struct(req); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			var errorMessages []string
			for _, fieldError := range validationErrors {
				errorMessages = append(errorMessages, fmt.Sprintf("Field '%s' failed validation: %s", fieldError.Field(), fieldError.Tag()))
			}
			errorMessage := strings.Join(errorMessages, ", ")

			responsePtr := new(TResponse)
			reflect.ValueOf(responsePtr).Elem().FieldByName("BaseResponse").Set(reflect.ValueOf(responses.NewBaseResponse(400, fmt.Sprintf("Validation errors: %s", errorMessage))))

			return responsePtr
		} else {
			responsePtr := new(TResponse)
			reflect.ValueOf(responsePtr).Elem().FieldByName("BaseResponse").Set(reflect.ValueOf(responses.NewBaseResponse(400, "Invalid request parameters")))

			return responsePtr
		}
	}
	return nil
}
