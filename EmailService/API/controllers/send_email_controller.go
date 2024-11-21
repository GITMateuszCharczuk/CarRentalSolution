package controllers

import (
	"file-storage/API/mappers"
	"file-storage/API/services"
	contract "file-storage/Application.contract/send_email"
	command "file-storage/Application/commands/send_email"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SendEmailController struct {
	commandHandler *command.SendEmailCommandHandler
}

func NewSendEmailController(handler *command.SendEmailCommandHandler) *SendEmailController {
	return &SendEmailController{
		commandHandler: handler,
	}
}

func (h *SendEmailController) Handle(c *gin.Context) {
	responseSender := services.NewResponseSender(c)

	var req contract.SendEmailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	command := mappers.MapToSendEmailCommand(&req)
	resp, err := h.commandHandler.Execute(command)
	if err != nil {
		responseSender.Send(contract.SendEmailResponse{
			Title:   "StatusInternalServerError",
			Message: fmt.Sprintf("Something went wrong: %v", err),
		})
		return
	}

	responseSender.Send(resp)
}

func (h *SendEmailController) Route() string {
	return "/send-email"
}

func (h *SendEmailController) Methods() []string {
	return []string{"POST"}
}
