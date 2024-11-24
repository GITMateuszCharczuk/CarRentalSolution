package controllers

import (
	"email-service/API/mappers"
	"email-service/API/services"
	contract "email-service/Application.contract/get_email"
	queries "email-service/Application/queries/get_email"

	"github.com/gin-gonic/gin"
)

type GetEmailController struct {
	queryHandler *queries.GetEmailQueryHandler
}

func NewGetEmailController(handler *queries.GetEmailQueryHandler) *GetEmailController {
	return &GetEmailController{
		queryHandler: handler,
	}
}

// Handle godoc
// @Summary Get email by ID
// @Description Retrieves an email based on its unique ID.
// @Tags emails
// @Accept json
// @Produce json
// @Param id path string true "Unique Email ID"
// @Success 200 {object} contract.GetEmailResponse200 "Email details retrieved successfully"
// @Failure 400 {object} contract.GetEmailResponse400 "Invalid request parameters"
// @Failure 404 {object} contract.GetEmailResponse404 "Email not found"
// @Failure 500 {object} contract.GetEmailResponse500 "Server error during email retrieval"
// @Router /email-service/api/emails/{id} [get]
func (h *GetEmailController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)
	emailID := c.Param("id")

	req := contract.GetEmailRequest{ID: emailID}
	query := mappers.MapToGetEmailQuery(&req)
	resp := h.queryHandler.Execute(query)

	responseSender.Send(resp)
}

func (h *GetEmailController) Route() string {
	return "/emails/:id"
}

func (h *GetEmailController) Methods() []string {
	return []string{"GET"}
}
