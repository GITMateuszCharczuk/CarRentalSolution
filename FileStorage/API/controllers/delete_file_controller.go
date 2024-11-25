package controllers

import (
	"file-storage/API/mappers"
	"file-storage/API/services"
	contract "file-storage/Application.contract/DeleteFile"
	commands "file-storage/Application/commands/delete_file"
	"fmt"

	"github.com/gin-gonic/gin"
)

type DeleteFileController struct {
	commandHandler *commands.DeleteFileCommandHandler
}

func NewDeleteFileController(cmd *commands.DeleteFileCommandHandler) *DeleteFileController {
	return &DeleteFileController{commandHandler: cmd}
}

// Handle godoc
// @Summary Delete a file
// @Description Deletes a file from storage by its unique ID. The file ID should be a valid identifier for an existing file.
// @Tags files
// @Accept json
// @Produce json
// @Param request body contract.DeleteFileRequest true "Delete file request object"
// @Success 200 {object} contract.DeleteFileResponse200 "File deletion was successful"
// @Failure 400 {object} contract.DeleteFileResponse400 "Invalid request format or parameters"
// @Failure 404 {object} contract.DeleteFileResponse404 "File not found with the given ID"
// @Failure 500 {object} contract.DeleteFileResponse500 "Server encountered an error during file deletion"
// @Router /file-storage/api/files/delete [delete]
func (h *DeleteFileController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)

	var req contract.DeleteFileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		responseSender.Send(contract.DeleteFileResponse{
			Title:   "StatusBadRequest",
			Message: fmt.Sprintf("Invalid request: %v", err),
		})
		return
	}

	command := mappers.MapToDeleteFileCommand(&req)
	resp := h.commandHandler.Execute(command)

	responseSender.Send(resp)
}

func (h *DeleteFileController) Route() string {
	return "/files/delete"
}

func (h *DeleteFileController) Methods() []string {
	return []string{"DELETE"}
}
