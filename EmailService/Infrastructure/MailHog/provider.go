package smtp

import (
	"file-storage/Infrastructure/config"

	"github.com/google/wire"
)

func ProvideEmailService(cfg *config.Config) (EmailSender, error) {
	return NewEmailSenderImpl(cfg.MailhogHost, cfg.MaihogPort, cfg.MailhogUsername, cfg.MaihogPassword)
}

var WireSet = wire.NewSet(ProvideEmailService)
