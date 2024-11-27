package queries

import (
	"context"
	contract "email-service/Application.contract/get_emails"
	fetcher "email-service/Domain/fetcher"
)

type GetEmailsQueryHandler struct {
	fetcher fetcher.DataFetcher
}

func NewGetEmailsQueryHandler(fetcher fetcher.DataFetcher) *GetEmailsQueryHandler {
	return &GetEmailsQueryHandler{fetcher: fetcher}
}

func (h *GetEmailsQueryHandler) Handle(ctx context.Context, query *GetEmailsQuery) (*contract.GetEmailsResponse, error) {
	emails, err := h.fetcher.GetEmails()
	if err != nil {
		return &contract.GetEmailsResponse{
			Title:   "StatusInternalServerError",
			Message: "Something went wrong",
		}, err
	}

	if len(*emails) == 0 {
		return &contract.GetEmailsResponse{
			Title:   "StatusNotFound",
			Message: "No emails found",
		}, nil
	}

	return &contract.GetEmailsResponse{
		Title:   "StatusOK",
		Message: "Emails retrieved successfully",
		Emails:  *emails,
	}, nil
}
