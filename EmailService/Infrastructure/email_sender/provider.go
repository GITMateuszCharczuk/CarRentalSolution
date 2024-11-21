package smtp

import (
	"email-service/Infrastructure/config"

	"github.com/google/wire"
)

func ProvideEmailService(cfg *config.Config) (EmailSender, error) {
	return NewEmailSenderImpl(cfg.MailhogHost, cfg.MailhogPort, cfg.MailhogUsername, cfg.MailhogPassword)
}

var WireSet = wire.NewSet(ProvideEmailService)
