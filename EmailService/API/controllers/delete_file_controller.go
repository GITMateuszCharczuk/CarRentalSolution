package controllers

import (
	"email-service/API/mappers"
	"email-service/API/services"
	contract "email-service/Application.contract/DeleteFile"
	commands "email-service/Application/commands/delete_file"
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
// @Param fileId path string true "Unique File ID to be deleted"
// @Success 200 {object} contract.DeleteFileResponse "File deletion was successful"
// @Failure 400 {object} contract.DeleteFileResponse "Invalid request format or parameters"
// @Failure 404 {object} contract.DeleteFileResponse "File not found with the given ID"
// @Failure 500 {object} contract.DeleteFileResponse "Server encountered an error during file deletion"
// @Router /files/delete [delete]
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
	resp, err := h.commandHandler.Execute(command)
	if err != nil {
		responseSender.Send(contract.DeleteFileResponse{
			Title:   "StatusInternalServerError",
			Message: fmt.Sprintf("Something went wrong: %v", err),
		})
		return
	}

	responseSender.Send(resp)
}

func (h *DeleteFileController) Route() string {
	return "/files/delete"
}

func (h *DeleteFileController) Methods() []string {
	return []string{"DELETE"}
}
