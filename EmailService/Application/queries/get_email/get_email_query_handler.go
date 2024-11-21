package queries

import (
	"file-storage/API/services"
	"fmt"
)

type GetEmailQueryHandler struct{}

func NewGetEmailQueryHandler() *GetEmailQueryHandler {
	return &GetEmailQueryHandler{}
}

func (h *GetEmailQueryHandler) Execute(command GetEmailQuery) (*GetEmailResponse, error) {
	emails, err := services.GetAllEmails()

	if err != nil {
		return nil, fmt.Errorf("failed to retrieve emails: %w", err)
	}

	for _, email := range emails {
		if email.ID == command.ID {
			return &GetEmailResponse{
				ID:      email.ID,
				From:    email.From,
				To:      email.To,
				Subject: email.Subject,
				Body:    email.Body,
			}, nil
		}
	}

	return nil, ErrEmailNotFound
}
