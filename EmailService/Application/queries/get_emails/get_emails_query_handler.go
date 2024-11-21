package queries

import (
	contract "email-service/Application.contract/get_emails"
	datafetcher "email-service/Infrastructure/data_fetcher"
)

type GetEmailsQueryHandler struct {
	fetcher datafetcher.DataFetcher
}

func NewGetEmailsQueryHandler(fetcher datafetcher.DataFetcher) *GetEmailsQueryHandler {
	return &GetEmailsQueryHandler{fetcher: fetcher}
}

func (h *GetEmailsQueryHandler) Execute(query GetEmailsQuery) *contract.GetEmailsResponse {
	emails, err := h.fetcher.GetEmails()
	if err != nil {
		println("failed to retrieve emails:", err.Error())
		return &contract.GetEmailsResponse{
			Title:   "StatusInternalServerError",
			Message: "Something went wrong",
		}
	}

	if len(*emails) == 0 {
		println("No emails found")
		return &contract.GetEmailsResponse{
			Title:   "StatusNotFound",
			Message: "No emails found",
		}
	}

	return &contract.GetEmailsResponse{
		Title:   "StatusOK",
		Message: "Emails retrieved successfully",
		Emails:  *emails,
	}
}
