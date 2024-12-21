package validators

import (
	"rental-api/Domain/constants"

	"github.com/go-playground/validator/v10"
)

func validateCarOrderStatus(fl validator.FieldLevel) bool {
	status := fl.Field().String()
	return constants.IsValidCarOrderStatus(status)
}

func validateCarOrderStatusArray(fl validator.FieldLevel) bool {
	if statuses, ok := fl.Field().Interface().([]string); ok {
		for _, status := range statuses {
			if !constants.IsValidCarOrderStatus(status) {
				return false
			}
		}
		return true
	}
	return false
}
