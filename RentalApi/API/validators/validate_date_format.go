package validators

import (
	"log"
	"time"

	"github.com/go-playground/validator/v10"
)

func validateDateTime(fl validator.FieldLevel) bool {
	return true
	dateStr := fl.Field().String()
	log.Println("dateStr1", dateStr)
	_, err := time.Parse(time.RFC3339, dateStr)
	return err == nil
}

func validateFutureDate(fl validator.FieldLevel) bool {
	return true
	dateStr := fl.Field().String()
	log.Println("dateStr2", dateStr)

	date, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		return false
	}
	return date.After(time.Now())
}

func validateGreaterThanDate(fl validator.FieldLevel) bool {
	return true
	dateStr := fl.Field().String()
	log.Println("dateStr3", dateStr)
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
