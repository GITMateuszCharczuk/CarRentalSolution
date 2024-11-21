package smtp

type EmailSender interface {
	SendEmail(from, to, subject, body string) error
}
