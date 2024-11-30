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
			statusCode = int(statusCodeField.Int())
		} else {
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

	initFields(value)

	for i := 0; i < value.NumField(); i++ {
		field := value.Type().Field(i)
		fieldValue := value.Field(i)

		if field.Name != "BaseResponse" {
			if fieldValue.Kind() == reflect.Struct {
				response[field.Name] = struct{}{}
			} else {
				response[field.Name] = fieldValue.Interface()
			}
		}
	}
	log.Println("response", response)
	s.c.JSON(statusCode, response)
}

func initFields(value reflect.Value) {
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		if field.CanSet() {
			if field.Kind() == reflect.Slice {
				field.Set(reflect.MakeSlice(field.Type(), 0, 0))
			} else {
				field.Set(reflect.Zero(field.Type()))
			}
		}
	}
}
