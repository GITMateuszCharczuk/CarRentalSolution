package queries

import (
	"context"
	contract "email-service/Application.contract/get_emails"
	fetcher "email-service/Domain/fetcher"
	"email-service/Domain/models"
	"email-service/Domain/responses"
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
		return createResponse(500, "Something went wrong", nil), nil
	}

	if len(*emails) == 0 {
		return createResponse(404, "No emails found", &[]models.Email{}), nil
	}

	return createResponse(200, "Emails retrieved successfully", emails), nil
}

func createResponse(statusCode int, message string, emails *[]models.Email) *contract.GetEmailsResponse {
	return &contract.GetEmailsResponse{
		BaseResponse: responses.NewBaseResponse(statusCode, message),
		Emails:       *emails,
	}
}
