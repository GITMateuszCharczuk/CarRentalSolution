package controllers

import (
	"file-storage/API/mappers"
	"file-storage/API/services"
	contract "file-storage/Application.contract/DeleteFile"
	commands "file-storage/Application/commands/delete_file"

	"github.com/gin-gonic/gin"
	"github.com/mehdihadeli/go-mediatr"
)

type DeleteFileController struct {
}

func NewDeleteFileController() *DeleteFileController {
	return &DeleteFileController{}
}

// Handle godoc
// @Summary Delete a file
// @Description Deletes a file from storage by its unique ID. The file ID should be a valid identifier for an existing file.
// @Tags files
// @Accept json
// @Produce json
// @Param token query string true "JWT token"
// @Param file_id path string true "File ID"
// @Success 200 {object} contract.DeleteFileResponse200 "File deletion was successful"
// @Failure 400 {object} contract.DeleteFileResponse400 "Invalid request format or parameters"
// @Failure 404 {object} contract.DeleteFileResponse404 "File not found with the given ID"
// @Failure 500 {object} contract.DeleteFileResponse500 "Server encountered an error during file deletion"
// @Router /file-storage/api/files/delete/{file_id} [delete]
func (h *DeleteFileController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)
	req := contract.DeleteFileRequest{
		JwtToken: services.GetJwtTokenFromQuery(c),
		FileID:   services.ExtractFromPath(c, "file_id"),
	}
	command := mappers.MapToDeleteFileCommand(&req)
	resp, err := mediatr.Send[*commands.DeleteFileCommand, *contract.DeleteFileResponse](c.Request.Context(), &command)
	if err != nil {
		responseSender.Send(contract.DeleteFileResponse{
			Title:   "StatusInternalServerError",
			Message: "Something went wrong",
		})
		return
	}

	responseSender.Send(resp)
}

func (h *DeleteFileController) Route() string {
	return "/files/delete/:file_id"
}

func (h *DeleteFileController) Methods() []string {
	return []string{"DELETE"}
}
