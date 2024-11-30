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
		return createResponse(500, "Something went wrong", &[]models.Email{}), nil
	}

	if len(*emails) == 0 {
		return createResponse(404, "No emails found", &[]models.Email{}), nil
	}

	return createResponse(200, "Emails retrieved successfully", emails), nil
}

func createResponse(statusCode int, message string, emails ...*[]models.Email) *contract.GetEmailsResponse {
	var emailList *[]models.Email
	if len(emails) > 0 {
		emailList = emails[0]
	} else {
		emailList = &[]models.Email{}
	}
	return &contract.GetEmailsResponse{
		BaseResponse: responses.NewBaseResponse(statusCode, message),
		Emails:       *emailList,
	}
}
