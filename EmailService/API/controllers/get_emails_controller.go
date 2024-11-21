package controllers

import (
	"file-storage/API/mappers"
	"file-storage/API/services"
	contract "file-storage/Application.contract/get_emails"
	queries "file-storage/Application/queries/get_emails"
	"fmt"

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
	resp, err := h.queryHandler.Execute(query)
	if err != nil {
		responseSender.Send(contract.GetEmailsResponse{
			Title:   "StatusInternalServerError",
			Message: fmt.Sprintf("Something went wrong: %v", err),
		})
		return
	}

	responseSender.Send(resp)
}

func (h *GetEmailsController) Route() string {
	return "/emails"
}

func (h *GetEmailsController) Methods() []string {
	return []string{"GET"}
}
