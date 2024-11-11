package controllers

import (
	"file-storage/API/mappers"
	contract "file-storage/Application.contract/DeleteFile"
	commands "file-storage/Application/commands/delete_file"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteFileController handles file deletion requests
type DeleteFileController struct {
	commandHandler *commands.DeleteFileCommandHandler
}

func NewDeleteFileController(cmd *commands.DeleteFileCommandHandler) *DeleteFileController {
	return &DeleteFileController{commandHandler: cmd}
}

// Handle godoc
// @Summary Delete a file
// @Description Delete a file by ID
// @Tags files
// @Accept json
// @Produce json
// @Param fileId path string true "File ID"
// @Success 200 {object} contract.DeleteFileResponse
// @Failure 400 {object} ErrorResponse "Invalid request"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /files/delete [delete]
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
