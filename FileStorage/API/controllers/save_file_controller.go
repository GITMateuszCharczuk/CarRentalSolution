package controllers

import (
	"file-storage/API/mappers"
	contract "file-storage/Application.contract/SaveFile"
	command "file-storage/Application/commands/save_file"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SaveFileController struct {
	commandHandler *command.SaveFileCommandHandler
}

func NewSaveFileController(cmd *command.SaveFileCommandHandler) *SaveFileController {
	return &SaveFileController{commandHandler: cmd}
}

func (h *SaveFileController) Handle(c *gin.Context) {
	var req contract.SaveFileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	command := mappers.MapToSaveFileCommand(&req)
	if command == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating GUID"})
		return
	}

	resp, err := h.commandHandler.Execute(*command)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func (h *SaveFileController) Route() string {
	return "/files"
}

func (h *SaveFileController) Methods() []string {
	return []string{"POST"}
}
