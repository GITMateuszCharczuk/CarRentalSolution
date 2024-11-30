package controllers

import (
	"email-service/API/mappers"
	"email-service/API/services"
	contract "email-service/Application.contract/get_emails"
	queries "email-service/Application/query_handlers/get_emails"

	"github.com/gin-gonic/gin"
)

type GetEmailsController struct {
}

func NewGetEmailsController() *GetEmailsController {
	return &GetEmailsController{}
}

// Handle godoc
// @Summary Get all emails
// @Description Retrieves a list of all emails.
// @Tags emails
// @Accept json
// @Produce json
// @Success 200 {object} contract.GetEmailsResponse200 "List of emails retrieved successfully"
// @Failure 400 {object} contract.GetEmailsResponse400 "Invalid request parameters"
// @Failure 500 {object} contract.GetEmailsResponse500 "Server error during emails retrieval"
// @Router /email-service/api/emails [get]
func (h *GetEmailsController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)
	req := contract.GetEmailsRequest{}
	query := mappers.MapToGetEmailsQuery(&req)
	resp := services.SendToMediator[*queries.GetEmailsQuery, *contract.GetEmailsResponse](c.Request.Context(), &query)
	responseSender.Send(resp)
}

func (h *GetEmailsController) Route() string {
	return "/emails"
}

func (h *GetEmailsController) Methods() []string {
	return []string{"GET"}
}
