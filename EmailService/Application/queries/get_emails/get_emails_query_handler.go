package queries

import (
	"file-storage/API/services"
	"fmt"
)

type GetEmailsQueryHandler struct{}

func NewGetEmailsQueryHandler() *GetEmailsQueryHandler {
	return &GetEmailsQueryHandler{}
}

func (h *GetEmailsQueryHandler) Execute(query GetEmailsQuery) (*GetAllEmailsResponse, error) {
	emails, err := services.GetAllEmails()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve emails: %w", err)
	}

	return &GetAllEmailsResponse{Emails: emails}, nil
}
