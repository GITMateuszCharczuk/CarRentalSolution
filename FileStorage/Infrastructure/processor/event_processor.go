package processor

import (
	"context"
	"file-storage/Domain/repository_interfaces"
	"file-storage/Infrastructure/mappers"
	"log"
)

type EventProcessorImpl struct {
	fileRepo repository_interfaces.FileRepository
}

func NewEventProcessorImpl(fileRepo repository_interfaces.FileRepository) *EventProcessorImpl {
	return &EventProcessorImpl{fileRepo: fileRepo}
}

func (p *EventProcessorImpl) ProcessUploadEvent(data interface{}) error {
	file, err := mappers.MapToFile(data)
	if err != nil {
		return err
	}
	if err := p.fileRepo.InsertFile(context.Background(), file); err != nil {
		log.Printf("Failed to insert file: %v", err)
		return err
	}

	log.Printf("Successfully processed 'upload' event for file ID: %s", file.ID)
	return nil
}

func (p *EventProcessorImpl) ProcessDeleteEvent(data interface{}) error {
	fileID, err := mappers.MapToFileID(data)
	if err != nil {
		return err
	}

	if err := p.fileRepo.DeleteFileByID(context.Background(), fileID); err != nil {
		log.Printf("Failed to delete file with ID: %s, error: %v", fileID, err)
		return err
	}

	log.Printf("Successfully processed 'delete' event for file ID: %s", fileID)
	return nil
}
