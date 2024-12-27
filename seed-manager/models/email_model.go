package models

import "time"

type EmailModel struct {
	ID          string
	Subject     string
	Body        string
	FromEmail   string
	ToEmail     string
	Status      string // "pending", "sent", "failed"
	CreatedAt   time.Time
	SentAt      *time.Time
	TemplateID  string
	Attachments []string // File IDs
}

type EmailTemplateModel struct {
	ID        string
	Name      string
	Subject   string
	Body      string
	CreatedAt time.Time
}
