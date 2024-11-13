package controllers

import (
	"file-storage/API/mappers"
	"file-storage/API/services"
	contract "file-storage/Application.contract/GetFile"
	queries "file-storage/Application/queries/get_file"
	"fmt"

	"github.com/gin-gonic/gin"
)

type GetFileController struct {
	queryHandler *queries.GetFileQueryHandler
}

func NewGetFileController(qrs *queries.GetFileQueryHandler) *GetFileController {
	return &GetFileController{queryHandler: qrs}
}

// Handle godoc
// @Summary Get a file
// @Description Retrieves a file from storage by its unique identifier. The ID should refer to a valid, stored file.
// @Tags files
// @Accept json
// @Produce json
// @Param file_id query string true "File ID"
// @Param owner_id query string true "Owner ID"
// @Success 200 {object} contract.GetFileResponse "Successful retrieval of file details"
// @Failure 400 {object} contract.GetFileResponse "Request contains invalid format or parameters"
// @Failure 404 {object} contract.GetFileResponse "File not found with the given ID"
// @Failure 500 {object} contract.GetFileResponse "Server encountered an error during file retrieval"
// @Router /files/get [get]
func (h *GetFileController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)

	fileID := c.Query("file_id")
	ownerID := c.Query("owner_id")

	if fileID == "" || ownerID == "" {
		responseSender.Send(contract.GetFileResponse{
			Title:   "StatusBadRequest",
			Message: "Missing required query parameters: file_id or owner_id",
		})
		return
	}

	req := contract.GetFileRequest{
		FileID:  fileID,
		OwnerID: ownerID,
	}

	query := mappers.MapToGetFileQuery(&req)
	resp, err := h.queryHandler.Execute(query)
	if err != nil {
		responseSender.Send(contract.GetFileResponse{
			Title:   "StatusInternalServerError",
			Message: fmt.Sprintf("Something went wrong: %v", err),
		})
		return
	}

	responseSender.Send(resp)
}

func (h *GetFileController) Route() string {
	return "/files/get"
}

func (h *GetFileController) Methods() []string {
	return []string{"GET"}
}
