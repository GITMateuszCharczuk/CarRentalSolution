package controllers

import (
	"file-storage/API/mappers"
	contract "file-storage/Application.contract/DeleteFile"
	commands "file-storage/Application/commands/delete_file"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DeleteFileController struct {
	commandHandler *commands.DeleteFileCommandHandler
}

func NewDeleteFileController(cmd *commands.DeleteFileCommandHandler) *DeleteFileController {
	return &DeleteFileController{commandHandler: cmd}
}

func (h *DeleteFileController) Handle(c *gin.Context) {
	var req contract.DeleteFileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	command := mappers.MapToDeleteFileCommand(&req)
	resp, err := h.commandHandler.Execute(command)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *DeleteFileController) Route() string {
	return "/files/delete"
}

func (h *DeleteFileController) Methods() []string {
	return []string{"DELETE"}
}
