package controllers

import (
	"file-storage/API/mappers"
	contract "file-storage/Application.contract/GetFile"
	queries "file-storage/Application/queries/get_file"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetFileController struct {
	queryHandler *queries.GetFileQueryHandler
}

func NewGetFileController(qrs *queries.GetFileQueryHandler) *GetFileController {
	return &GetFileController{queryHandler: qrs}
}

func (h *GetFileController) Handle(c *gin.Context) {
	var req contract.GetFileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	query := mappers.MapToGetFileQuery(&req)
	resp, err := h.queryHandler.Execute(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *GetFileController) Route() string {
	return "/files/get"
}

func (h *GetFileController) Methods() []string {
	return []string{"GET"}
}
