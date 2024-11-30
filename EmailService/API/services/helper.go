package services

import (
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

	for i := 0; i < value.NumField(); i++ {
		field := value.Type().Field(i)
		if field.Name != "BaseResponse" && field.Type.Kind() != reflect.Struct {
			response[field.Name] = value.Field(i).Interface()
		}
	}

	s.c.JSON(statusCode, response)
}
