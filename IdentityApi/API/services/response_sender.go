package services

import (
	"log"
	"net/http"
	"reflect"

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
			log.Println("statusCodeField", statusCodeField)
			statusCode = int(statusCodeField.Int())
		} else {
			log.Println("statusCodeField is not valid or not int")
			statusCode = http.StatusInternalServerError
		}
	} else {
		statusCode = http.StatusInternalServerError
	}

	response := make(map[string]interface{})

	if baseResponseField.IsValid() {
		response["success"] = baseResponseField.FieldByName("Success").Bool()
		response["message"] = baseResponseField.FieldByName("Message").String()
	}
	log.Println("value", value)
	initFields(value)
	log.Println("value", value)
	for i := 0; i < value.NumField(); i++ {
		field := value.Type().Field(i)
		fieldValue := value.Field(i)
		if field.Name != "BaseResponse" {
			if fieldValue.Kind() == reflect.Struct && isZeroValue(fieldValue) {
				response[field.Name] = struct{}{}
			} else {
				response[field.Name] = fieldValue.Interface()
			}
		}
	}
	log.Println("value", value)
	s.c.JSON(statusCode, response)
}

func initFields(value reflect.Value) {
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		if field.CanSet() && isZeroValue(field) {
			switch field.Kind() {
			case reflect.Slice:
				field.Set(reflect.MakeSlice(field.Type(), 0, 0))
			default:
				field.Set(reflect.Zero(field.Type()))
			}
		}
	}
}

func isZeroValue(value reflect.Value) bool {
	if value.Kind() == reflect.Slice {
		return value.Len() == 0
	}
	return value.Interface() == reflect.Zero(value.Type()).Interface()
}
