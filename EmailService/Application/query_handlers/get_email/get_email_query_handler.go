package queries

import (
	"context"
	contract "email-service/Application.contract/get_email"
	fetcher "email-service/Domain/fetcher"
	"email-service/Domain/models"
)

type GetEmailQueryHandler struct {
	fetcher fetcher.DataFetcher
}

func NewGetEmailQueryHandler(fetcher fetcher.DataFetcher) *GetEmailQueryHandler {
	return &GetEmailQueryHandler{fetcher: fetcher}
}

func (h *GetEmailQueryHandler) Handle(ctx context.Context, query *GetEmailQuery) (*contract.GetEmailResponse, error) {
	emails, err := h.fetcher.GetEmails()
	if err != nil {
		return &contract.GetEmailResponse{
			Title:   "StatusInternalServerError",
			Message: "Something went wrong",
		}, err
	}

	var resEmail *models.Email = nil

	for _, email := range *emails {
		if email.ID == query.ID {
			resEmail = &email
		}
	}

	if resEmail == nil {
		return &contract.GetEmailResponse{
			Title:   "StatusNotFound",
			Message: "No emails found",
		}, nil
	}

	return &contract.GetEmailResponse{
		Title:   "StatusOK",
		Message: "Emails retrieved successfully",
		Email:   *resEmail,
	}, nil
}
