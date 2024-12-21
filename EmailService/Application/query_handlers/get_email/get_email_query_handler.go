package queries

import (
	"context"
	contract "email-service/Application.contract/get_email"
	"email-service/Application/utils"
	fetcher "email-service/Domain/fetcher"
	"email-service/Domain/models"
	pagination "email-service/Domain/requests"
	"email-service/Domain/responses"
	service_interfaces "email-service/Domain/service_interfaces"
)

type GetEmailQueryHandler struct {
	fetcher               fetcher.DataFetcher
	microserviceConnector service_interfaces.MicroserviceConnector
}

func NewGetEmailQueryHandler(fetcher fetcher.DataFetcher, microserviceConnector service_interfaces.MicroserviceConnector) *GetEmailQueryHandler {
	return &GetEmailQueryHandler{fetcher: fetcher, microserviceConnector: microserviceConnector}
}

func (h *GetEmailQueryHandler) Handle(ctx context.Context, query *GetEmailQuery) (*contract.GetEmailResponse, error) {
	user, err := h.microserviceConnector.GetUserInternalInfo(query.JwtToken)
	if err != nil {
		return createResponse(401, "Invalid JWT token", &models.Email{}), nil
	}

	if !utils.IsAdminOrSuperAdmin(user.Roles) {
		return createResponse(403, "You are not authorized retrive emails", &models.Email{}), nil
	}

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
