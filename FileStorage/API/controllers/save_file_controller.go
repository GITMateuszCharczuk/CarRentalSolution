package controllers

import (
	"file-storage/API/mappers"
	"file-storage/API/services"
	contract "file-storage/Application.contract/SaveFile"
	command "file-storage/Application/commands/save_file"
	"fmt"

	"github.com/gin-gonic/gin"
)

type SaveFileController struct {
	commandHandler *command.SaveFileCommandHandler
}

func NewSaveFileController(cmd *command.SaveFileCommandHandler) *SaveFileController {
	return &SaveFileController{commandHandler: cmd}
}

// Handle godoc
// @Summary Save a new file
// @Description Uploads and saves a file in the storage system, including metadata and content.
// @Tags files
// @Accept json
// @Produce json
// @Param file body contract.SaveFileRequest true "File metadata and content for saving"
// @Success 201 {object} contract.SaveFileResponse "File saved successfully with unique ID and details"
// @Failure 400 {object} contract.SaveFileResponse "Invalid request format or file data"
// @Failure 500 {object} contract.SaveFileResponse "Server encountered an error during file save operation"
// @Router /files [post]
func (h *SaveFileController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)

	var req contract.SaveFileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		responseSender.Send(contract.SaveFileResponse{
			Title:   "StatusBadRequest",
			Message: fmt.Sprintf("Invalid JSON format: %v", err),
		})
		return
	}

	command := mappers.MapToSaveFileCommand(&req)

	resp, err := h.commandHandler.Execute(command)
	if err != nil {
		responseSender.Send(contract.SaveFileResponse{
			Title:   "StatusInternalServerError",
			Message: fmt.Sprintf("Something went wrong: %v", err),
		})
		return
	}

	responseSender.Send(resp)
}

func (h *SaveFileController) Route() string {
	return "/files"
}

func (h *SaveFileController) Methods() []string {
	return []string{"POST"}
}
