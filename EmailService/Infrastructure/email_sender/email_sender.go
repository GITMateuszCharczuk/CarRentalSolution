package smtp

import (
	"fmt"
	"strconv"
	"time"

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

	done := make(chan error, 1)

	go func() {
		done <- es.dialer.DialAndSend(msg)
	}()

	select {
	case err := <-done:
		if err != nil {
			return fmt.Errorf("failed to send email: %w", err)
		}
		fmt.Printf("Successfully sent email to: %s\n", to)
	case <-time.After(5 * time.Second):
		return fmt.Errorf("sending email timed out")
	}

	return nil
}
