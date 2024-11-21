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

	titleField := value.FieldByName("Title")

	if titleField.IsValid() && titleField.Kind() == reflect.String {
		switch titleField.String() {
		case "StatusOK":
			s.c.JSON(http.StatusOK, obj)
			return
		case "StatusBadRequest":
			s.c.JSON(http.StatusBadRequest, obj)
			return
		case "StatusNotFound":
			s.c.JSON(http.StatusNotFound, obj)
			return
		case "StatusInternalServerError":
			s.c.JSON(http.StatusInternalServerError, obj)
			return
		default:
			s.c.JSON(http.StatusInternalServerError, obj)
			return
		}
	}

	s.c.JSON(http.StatusInternalServerError, obj)
}
