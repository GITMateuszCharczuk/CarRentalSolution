package smtp

import (
	"fmt"
	"strconv"

	"github.com/go-mail/mail"
)

type EmailSenderImpl struct {
	dialer *mail.Dialer
}

func NewEmailSenderImpl(smtpHost string, smtpPort string, username, password string) (*EmailSenderImpl, error) {
	mailhogPort, err := strconv.Atoi(smtpPort)

	if err != nil {
		fmt.Println("Error converting string to integer:", err)
		return nil, err
	}

	dialer := mail.NewDialer(smtpHost, mailhogPort, username, password)
	return &EmailSenderImpl{
		dialer: dialer,
	}, nil
}

func (es *EmailSenderImpl) SendEmail(from, to, subject, body string) error {
	msg := mail.NewMessage()
	msg.SetHeader("From", from)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/plain", body)

	if err := es.dialer.DialAndSend(msg); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}
	return nil
}
