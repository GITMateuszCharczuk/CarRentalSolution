package services

import (
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

type ResponseSender struct {
	c *gin.Context
}

func NewResponseSender(c *gin.Context) *ResponseSender {
	return &ResponseSender{c: c}
}

func (s *ResponseSender) Send(obj interface{}) {
	value := reflect.ValueOf(obj)

	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	baseResponseField := value.FieldByName("BaseResponse")

	var statusCode int
	if baseResponseField.IsValid() {
		statusCodeField := baseResponseField.FieldByName("StatusCode")
		if statusCodeField.IsValid() && statusCodeField.Kind() == reflect.Int {
			statusCode = int(statusCodeField.Int())
		} else {
			statusCode = http.StatusInternalServerError
		}
	} else {
		statusCode = http.StatusInternalServerError
	}

	response := make(map[string]interface{})

	if baseResponseField.IsValid() {
		// response["status_code"] = statusCode
		response["success"] = baseResponseField.FieldByName("Success").Bool()
		response["message"] = baseResponseField.FieldByName("Message").String()
	}

	initFields(value)

	for i := 0; i < value.NumField(); i++ {
		field := value.Type().Field(i)
		fieldValue := value.Field(i)
		if field.Name != "BaseResponse" {
			if fieldValue.Kind() == reflect.Struct {
				if field.Name == "BlogPost" {
					response[getJSONFieldName(field)] = interface{}(fieldValue.Interface())
					continue
				}
				structFields := extractStructFields(fieldValue)
				for k, v := range structFields {
					response[k] = v
				}
			} else {
				response[field.Name] = fieldValue.Interface()
			}
		}
	}

	s.c.JSON(statusCode, response)
}

func getJSONFieldName(field reflect.StructField) string {
	jsonTag := field.Tag.Get("json")
	if jsonTag == "" {
		return field.Name
	}
	parts := strings.Split(jsonTag, ",")
	return parts[0]
}

func initFields(value reflect.Value) {
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		if field.CanSet() && isEmptyValue(field) {
			switch field.Kind() {
			case reflect.Slice:
				field.Set(reflect.MakeSlice(field.Type(), 0, 0))
			case reflect.Struct:
				initFields(field)
			default:
				field.Set(reflect.Zero(field.Type()))
			}
		}
	}
}

func extractStructFields(value reflect.Value) map[string]interface{} {
	result := make(map[string]interface{})

	for i := 0; i < value.NumField(); i++ {
		field := value.Type().Field(i)
		fieldValue := value.Field(i)

		if field.Anonymous {
			// Handle embedded struct
			embedded := extractStructFields(fieldValue)
			for k, v := range embedded {
				result[k] = v
			}
		} else {
			result[field.Name] = fieldValue.Interface()
		}
	}

	return result
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	case reflect.Struct:
		// For structs, check if all fields are empty
		for i := 0; i < v.NumField(); i++ {
			if !isEmptyValue(v.Field(i)) {
				return false
			}
		}
		return true
	}
	return false
}
