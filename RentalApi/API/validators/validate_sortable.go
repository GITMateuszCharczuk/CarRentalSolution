package validators

import (
	"fmt"
	"log"
	"reflect"
	"rental-api/Domain/sorting"
	"sort"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateSortQuery[T any](fl validator.FieldLevel) bool {
	sortQueries, ok := fl.Field().Interface().([]string)
	if !ok {
		log.Printf("Field is not a string slice, got: %T", fl.Field().Interface())
		return false
	}

	if len(sortQueries) == 0 {
		return true
	}

	validFields := extractFields[T]()
	validFieldsList := getValidFieldsList(validFields)

	for i, query := range sortQueries {
		parts := strings.Split(query, ":")
		if len(parts) != 2 {
			log.Printf("Invalid sort query format at index %d: '%s'. Expected format: 'field:direction'", i, query)
			return false
		}

		field := parts[0]
		direction := strings.ToLower(parts[1])

		if !validFields[field] {
			log.Printf("Invalid sort field: '%s'. Valid fields are: %s",
				field,
				strings.Join(validFieldsList, ", "))
			return false
		}

		if direction != string(sorting.ASC) && direction != string(sorting.DESC) {
			log.Printf("Invalid sort direction: '%s' for field '%s'. Must be either 'asc' or 'desc'",
				direction,
				field)
			return false
		}
	}

	return true
}

func getValidFieldsList(validFields map[string]bool) []string {
	fields := make([]string, 0, len(validFields))
	for field := range validFields {
		fields = append(fields, fmt.Sprintf("'%s'", field))
	}
	sort.Strings(fields)
	return fields
}

func extractFields[T any]() map[string]bool {
	validFields := make(map[string]bool)
	t := reflect.TypeOf((*T)(nil)).Elem()

	var extractFieldsFromType func(reflect.Type)
	extractFieldsFromType = func(t reflect.Type) {
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)

			if field.Anonymous {
				extractFieldsFromType(field.Type)
				continue
			}

			jsonTag := field.Tag.Get("json")
			if jsonTag == "" || jsonTag == "-" {
				continue
			}

			tagParts := strings.Split(jsonTag, ",")
			fieldName := tagParts[0]

			description := field.Tag.Get("description")
			if description != "" {
				log.Printf("Field '%s': %s", fieldName, description)
			}

			validFields[fieldName] = true
		}
	}

	extractFieldsFromType(t)
	return validFields
}

func RegisterSortQueryValidator[T any](v *validator.Validate, tag string) {
	validFields := extractFields[T]()
	validFieldsList := getValidFieldsList(validFields)
	log.Printf("Registered sort validator for tag '%s'. Valid fields: %s",
		tag,
		strings.Join(validFieldsList, ", "))

	v.RegisterValidation(tag, func(fl validator.FieldLevel) bool {
		return ValidateSortQuery[T](fl)
	})
}
