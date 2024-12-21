// commands/delete_file.go
package commands

import (
	"context"
	contract "file-storage/Application.contract/DeleteFile"
	"file-storage/Application/utils"
	"file-storage/Domain/event"
	"file-storage/Domain/models"
	interfaces "file-storage/Domain/service_interfaces"
	"fmt"
)

type DeleteFileCommandHandler struct {
	eventPublisher        event.EventPublisher
	microserviceConnector interfaces.MicroserviceConnector
}

func NewDeleteFileCommandHandler(eventPublisher event.EventPublisher, microserviceConnector interfaces.MicroserviceConnector) *DeleteFileCommandHandler {
	return &DeleteFileCommandHandler{
		eventPublisher:        eventPublisher,
		microserviceConnector: microserviceConnector,
	}
}

func (cmd *DeleteFileCommandHandler) Handle(ctx context.Context, command *DeleteFileCommand) (*contract.DeleteFileResponse, error) {
	user, err := cmd.microserviceConnector.GetUserInternalInfo(command.JwtToken)
	if err != nil {
		return &contract.DeleteFileResponse{
			Title:   "StatusUnauthorized",
			Message: "Invalid JWT token",
		}, nil
	}

	if !utils.IsAdminOrSuperAdmin(user.Roles) {
		return &contract.DeleteFileResponse{
			Title:   "StatusForbidden",
			Message: "You are not authorized to delete files",
		}, nil
	}

	if err := cmd.eventPublisher.PublishEvent("file-events.delete", command.FileID, models.EventTypeDelete); err != nil {
		return &contract.DeleteFileResponse{
			Title:   "StatusInternalServerError",
			Message: fmt.Sprintf("Failed to delete file: %v", err),
		}, nil
	}

	return &contract.DeleteFileResponse{
		Title:   "StatusOK",
		Message: "File deleted successfully",
	}, nil
}
