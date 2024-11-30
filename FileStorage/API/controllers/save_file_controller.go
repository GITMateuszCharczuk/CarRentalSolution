package controllers

import (
	"file-storage/API/mappers"
	"file-storage/API/services"
	contract "file-storage/Application.contract/SaveFile"
	commands "file-storage/Application/commands/save_file"

	"github.com/gin-gonic/gin"
	"github.com/mehdihadeli/go-mediatr"
)

type SaveFileController struct {
}

func NewSaveFileController() *SaveFileController {
	return &SaveFileController{}
}

// Handle godoc
// @Summary Save a new file
// @Description Uploads and saves a file in the storage system, including metadata and content.
// @Tags files
// @Accept multipart/form-data
// @Produce json
// @Param owner_id formData string true "Owner ID associated with the file"
// @Param file formData file true "Binary file content (JPEG, PNG, etc.) to be saved"
// @Success 201 {object} contract.SaveFileResponse201 "File saved successfully with unique ID and details"
// @Failure 400 {object} contract.SaveFileResponse400 "Invalid request format or missing parameters"
// @Failure 500 {object} contract.SaveFileResponse500 "Server encountered an error during file save operation"
// @Router /file-storage/api/files [post]
func (h *SaveFileController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)

	ownerID := c.PostForm("owner_id")
	file, err := c.FormFile("file")

	if err != nil || ownerID == "" {
		responseSender.Send(contract.SaveFileResponse{
			Title:   "StatusBadRequest",
			Message: "Missing file or owner_id",
		})
		return
	}

	req := &contract.SaveFileRequest{
		OwnerID: ownerID,
		File:    file,
	}

	command := mappers.MapToSaveFileCommand(req)
	resp, err := mediatr.Send[*commands.SaveFileCommand, *contract.SaveFileResponse](c.Request.Context(), &command)
	if err != nil {
		responseSender.Send(contract.SaveFileResponse{
			Title:   "StatusInternalServerError",
			Message: "Something went wrong",
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
