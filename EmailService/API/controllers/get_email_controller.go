package controllers

import (
	"file-storage/API/mappers"
	"file-storage/API/services"
	contract "file-storage/Application.contract/get_email"
	queries "file-storage/Application/queries/get_email"
	"fmt"

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
	resp, err := h.queryHandler.Execute(query)
	if err != nil {
		responseSender.Send(contract.GetEmailResponse{
			Title:   "StatusInternalServerError",
			Message: fmt.Sprintf("Something went wrong: %v", err),
		})
		return
	}

	responseSender.Send(resp)
}

func (h *GetEmailController) Route() string {
	return "/emails/:id"
}

func (h *GetEmailController) Methods() []string {
	return []string{"GET"}
}
