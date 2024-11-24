package queries

import (
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

func (cmd *GetEmailQueryHandler) Execute(command GetEmailQuery) *contract.GetEmailResponse {
	emails, err := cmd.fetcher.GetEmails()
	if err != nil {
		println("failed to retrieve emails:", err.Error())
		return &contract.GetEmailResponse{
			Title:   "StatusInternalServerError",
			Message: "Something went wrong",
		}
	}

	var resEmail *models.Email = nil

	for _, email := range *emails {
		if email.ID == command.ID {
			resEmail = &email
		}
	}

	if resEmail == nil {
		return &contract.GetEmailResponse{
			Title:   "StatusNotFound",
			Message: "No emails found",
		}
	}

	return &contract.GetEmailResponse{
		Title:   "StatusOK",
		Message: "Emails retrieved successfully",
		Email:   *resEmail,
	}
}
