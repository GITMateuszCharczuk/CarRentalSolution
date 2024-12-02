package queries

import (
	"context"
	contract "identity-api/Application.contract/get_email"
	fetcher "identity-api/Domain/fetcher"
	"identity-api/Domain/models"
	pagination "identity-api/Domain/requests"
	"identity-api/Domain/responses"
)

type GetEmailQueryHandler struct {
	fetcher fetcher.DataFetcher
}

func NewGetEmailQueryHandler(fetcher fetcher.DataFetcher) *GetEmailQueryHandler {
	return &GetEmailQueryHandler{fetcher: fetcher}
}

func (h *GetEmailQueryHandler) Handle(ctx context.Context, query *GetEmailQuery) (*contract.GetEmailResponse, error) {
	emptyPagination := pagination.Pagination{}
	emails, err := h.fetcher.GetEmails(emptyPagination)
	if err != nil {
		return createResponse(500, "Something went wrong", &models.Email{}), nil
	}

	var resEmail *models.Email = nil

	for _, email := range *emails {
		if email.ID == query.ID {
			resEmail = &email
			break
		}
	}

	if resEmail == nil {
		return createResponse(404, "No emails found", &models.Email{}), nil
	}

	return createResponse(200, "Email retrieved successfully", resEmail), nil
}

func createResponse(statusCode int, message string, email *models.Email) *contract.GetEmailResponse {
	return &contract.GetEmailResponse{
		BaseResponse: responses.NewBaseResponse(statusCode, message),
		Email:        *email,
	}
}
