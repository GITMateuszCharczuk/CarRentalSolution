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
// @Description Retrieves a file from storage by its unique identifier. The ID should refer to a valid, stored file, and the file is returned in its original format (e.g., JPEG, PNG).
// @Tags files
// @Accept json
// @Produce octet-stream
// @Param file_id query string true "Unique File ID for retrieval"
// @Param owner_id query string true "Owner ID associated with the file"
// @Param download query string false "File binary content or file to download"
// @Success 200 {file} binary "Successful retrieval of file in binary format or file to download in attachment"
// @Failure 400 {object} contract.GetFileResponse400 "Request contains invalid format or parameters"
// @Failure 404 {object} contract.GetFileResponse404 "File not found with the given ID"
// @Failure 500 {object} contract.GetFileResponse500 "Server encountered an error during file retrieval"
// @Router /file-storage/api/files/get [get]
func (h *GetFileController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)

	fileID := c.Query("file_id")
	ownerID := c.Query("owner_id")
	download := c.Query("download")

	req := contract.GetFileRequest{
		FileID:  fileID,
		OwnerID: ownerID,
	}

	query := mappers.MapToGetFileQuery(&req)
	errorResponse, fileStream, fileExt, err := h.queryHandler.Execute(query)
	if err != nil {
		responseSender.Send(errorResponse)
		return
	}

	c.Header("Content-Type", *fileExt)
	if download == "true" {
		c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", fileStream.FileName))
	} else {
		c.Header("Content-Disposition", fmt.Sprintf("inline; filename=\"%s\"", fileStream.FileName))
	}
	c.DataFromReader(200, fileStream.FileSize, *fileExt, fileStream.Stream, nil)
}

func (h *GetFileController) Route() string {
	return "/files/get"
}

func (h *GetFileController) Methods() []string {
	return []string{"GET"}
}
