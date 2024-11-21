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
