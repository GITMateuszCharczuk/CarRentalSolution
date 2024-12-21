package queries

import (
	"context"
	contract "email-service/Application.contract/get_emails"
	"email-service/Application/utils"
	fetcher "email-service/Domain/fetcher"
	"email-service/Domain/models"
	"email-service/Domain/responses"
	service_interfaces "email-service/Domain/service_interfaces"
)

type GetEmailsQueryHandler struct {
	fetcher               fetcher.DataFetcher
	microserviceConnector service_interfaces.MicroserviceConnector
}

func NewGetEmailsQueryHandler(fetcher fetcher.DataFetcher, microserviceConnector service_interfaces.MicroserviceConnector) *GetEmailsQueryHandler {
	return &GetEmailsQueryHandler{fetcher: fetcher, microserviceConnector: microserviceConnector}
}

func (h *GetEmailsQueryHandler) Handle(ctx context.Context, query *GetEmailsQuery) (*contract.GetEmailsResponse, error) {
	user, err := h.microserviceConnector.GetUserInternalInfo(query.JwtToken)
	if err != nil {
		return createResponse(401, "Invalid JWT token", &[]models.Email{}), nil
	}

	if !utils.IsAdminOrSuperAdmin(user.Roles) {
		return createResponse(403, "You are not authorized retrive emails", &[]models.Email{}), nil
	}

	emails, err := h.fetcher.GetEmails(query.Pagination)
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
