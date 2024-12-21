package validators

import (
	"time"

	"github.com/go-playground/validator/v10"
)

func validateDateTime(fl validator.FieldLevel) bool {
	dateStr := fl.Field().String()
	_, err := time.Parse(time.RFC3339, dateStr)
	return err == nil
}

func validateFutureDate(fl validator.FieldLevel) bool {
	dateStr := fl.Field().String()
	date, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		return false
	}
	return date.After(time.Now())
}

func validateGreaterThanDate(fl validator.FieldLevel) bool {
	dateStr := fl.Field().String()
	compareToField := fl.Param()

	date, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		return false
	}

	compareToValue := fl.Parent().FieldByName(compareToField).String()
	compareToDate, err := time.Parse(time.RFC3339, compareToValue)
	if err != nil {
		return false
	}

	return date.After(compareToDate)
}
