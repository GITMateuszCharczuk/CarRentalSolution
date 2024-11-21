package controllers

import (
	"email-service/API/mappers"
	"email-service/API/services"
	contract "email-service/Application.contract/GetFile"
	queries "email-service/Application/queries/get_file"
	"fmt"

	"github.com/gin-gonic/gin"
)

type GetFileController struct {
	queryHandler *queries.GetFileQueryHandler
}

func NewGetFileController(qrs *queries.GetFileQueryHandler) *GetFileController {
	return &GetFileController{queryHandler: qrs}
}

// The commented section starting with `// Handle godoc` is a documentation comment written in a
// specific format that is compatible with tools like Swagger and other API documentation generators.
// This comment provides detailed information about the `Handle` method of the `GetFileController`
// struct.
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
// @Failure 400 {object} contract.GetFileResponse "Request contains invalid format or parameters"
// @Failure 404 {object} contract.GetFileResponse "File not found with the given ID"
// @Failure 500 {object} contract.GetFileResponse "Server encountered an error during file retrieval"
// @Router /files/get [get]
func (h *GetFileController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)

	fileID := c.Query("file_id")
	ownerID := c.Query("owner_id")
	download := c.Query("download")

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
	_, fileStream, fileExt, err := h.queryHandler.Execute(query)
	if err != nil {
		responseSender.Send(contract.GetFileResponse{
			Title:   "StatusInternalServerError",
			Message: fmt.Sprintf("Something went wrong: %v", err),
		})
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
