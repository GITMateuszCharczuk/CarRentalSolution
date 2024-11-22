package controllers

import (
	"email-service/API/mappers"
	"email-service/API/services"
	contract "email-service/Application.contract/get_emails"
	queries "email-service/Application/queries/get_emails"

	"github.com/gin-gonic/gin"
)

type GetEmailsController struct {
	queryHandler *queries.GetEmailsQueryHandler
}

func NewGetEmailsController(handler *queries.GetEmailsQueryHandler) *GetEmailsController {
	return &GetEmailsController{
		queryHandler: handler,
	}
}

func (h *GetEmailsController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)

	req := contract.GetEmailsRequest{}
	query := mappers.MapToGetEmailsQuery(&req)
	resp := h.queryHandler.Execute(query)

	responseSender.Send(resp)
}

// Handle godoc
// @Summary Get all emails
// @Description Retrieves a list of all emails.
// @Tags emails
// @Accept json
// @Produce json
// @Success 200 {object} contract.GetEmailsResponse "List of emails retrieved successfully"
// @Failure 400 {object} contract.GetEmailsResponse "Invalid request parameters"
// @Failure 500 {object} contract.GetEmailsResponse "Server error during emails retrieval"
// @Router /emails [get]
func (h *GetEmailsController) Route() string {
	return "/emails"
}

func (h *GetEmailsController) Methods() []string {
	return []string{"GET"}
}
