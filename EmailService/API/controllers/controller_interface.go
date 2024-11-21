package controllers

import "github.com/gin-gonic/gin"

type Controller interface {
	Handle(c *gin.Context)
	Route() string
	Methods() []string
}
