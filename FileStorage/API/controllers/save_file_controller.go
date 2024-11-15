package controllers

import (
	"file-storage/API/mappers"
	"file-storage/API/services"
	contract "file-storage/Application.contract/SaveFile"
	command "file-storage/Application/commands/save_file"
	"file-storage/Domain/constants"
	"fmt"
	"path/filepath"

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
// @Accept multipart/form-data
// @Produce json
// @Param owner_id formData string true "Owner ID associated with the file"
// @Param file formData file true "Binary file content (JPEG, PNG, etc.) to be saved"
// @Success 201 {object} contract.SaveFileResponse "File saved successfully with unique ID and details"
// @Failure 400 {object} contract.SaveFileResponse "Invalid request format or missing parameters"
// @Failure 500 {object} contract.SaveFileResponse "Server encountered an error during file save operation"
// @Router /files [post]
func (h *SaveFileController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)

	ownerID := c.PostForm("owner_id")
	file, err := c.FormFile("file")

	if err != nil || ownerID == "" {
		responseSender.Send(contract.SaveFileResponse{
			Title:   "StatusBadRequest",
			Message: "Missing file, owner_id, or file_name",
		})
		return
	}

	fileExtension := filepath.Ext(file.Filename)
	if fileExtension == "" {
		responseSender.Send(contract.SaveFileResponse{
			Title:   "StatusBadRequest",
			Message: "File must have a valid extension",
		})
		return
	}

	if !constants.IsAllowedExtension(fileExtension) {
		responseSender.Send(contract.SaveFileResponse{
			Title:   "StatusBadRequest",
			Message: fmt.Sprintf("File extension '%s' is not allowed", fileExtension),
		})
		return
	}

	req := &contract.SaveFileRequest{
		OwnerID: ownerID,
		File:    file,
	}

	command := mappers.MapToSaveFileCommand(req)
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
