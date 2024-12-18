package validators

import (
	"blog-api/Domain/constants"
	"log"
	"reflect"

	"github.com/go-playground/validator/v10"
)

func ValidateRole(fl validator.FieldLevel) bool {
	role, ok := fl.Field().Interface().(string)
	if !ok {
		log.Println("Validation failed: Field is not of type string")
		return false
	}

	isValid := constants.JWTRole(role).IsValid()
	if !isValid {
		log.Printf("Validation failed: Role '%s' is not valid", role)
	}
	return isValid
}

func ValidateRoles(fl validator.FieldLevel) bool {
	field := fl.Field()
	if field.Kind() != reflect.Slice {
		log.Printf("Validation failed: Field is not a slice, got %v", field.Kind())
		return false
	}

	if field.Len() == 0 {
		log.Println("Validation failed: No roles provided")
		return false
	}

	for i := 0; i < field.Len(); i++ {
		role := field.Index(i).String()
		if !constants.JWTRole(role).IsValid() {
			log.Printf("Validation failed: Role '%s' is not valid", role)
			return false
		}
	}

	log.Printf("Validation succeeded: All %d roles are valid", field.Len())
	return true
}
