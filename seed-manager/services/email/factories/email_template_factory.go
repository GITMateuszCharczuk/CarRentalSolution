package factories

import (
	"time"

	"seeder-manager/models"

	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
)

type EmailTemplateFactory struct{}

func NewEmailTemplateFactory() *EmailTemplateFactory {
	return &EmailTemplateFactory{}
}

func (f *EmailTemplateFactory) Create() (*models.EmailTemplateModel, error) {
	templateTypes := []string{
		"welcome",
		"order_confirmation",
		"order_cancelled",
		"payment_received",
		"rental_reminder",
	}

	index, err := faker.RandomInt(0, len(templateTypes)-1, 1)
	if err != nil {
		return nil, err
	}
	templateType := templateTypes[index[0]]

	return &models.EmailTemplateModel{
		ID:        uuid.New().String(),
		Name:      templateType,
		Subject:   f.getSubjectByType(templateType),
		Body:      f.getBodyByType(templateType),
		CreatedAt: time.Now(),
	}, nil
}

func (f *EmailTemplateFactory) CreateMany(count int) ([]*models.EmailTemplateModel, error) {
	templates := make([]*models.EmailTemplateModel, 0, count)
	for i := 0; i < count; i++ {
		template, err := f.Create()
		if err != nil {
			return nil, err
		}
		templates = append(templates, template)
	}
	return templates, nil
}

func (f *EmailTemplateFactory) getSubjectByType(templateType string) string {
	switch templateType {
	case "welcome":
		return "Welcome to Car Rental Service"
	case "order_confirmation":
		return "Your Car Rental Order Confirmation"
	case "order_cancelled":
		return "Car Rental Order Cancellation"
	case "payment_received":
		return "Payment Received - Car Rental Service"
	case "rental_reminder":
		return "Upcoming Car Rental Reminder"
	default:
		return "Car Rental Service Notification"
	}
}

func (f *EmailTemplateFactory) getBodyByType(templateType string) string {
	switch templateType {
	case "welcome":
		return "Dear {UserName},\n\nWelcome to Car Rental Service! We're excited to have you on board..."
	case "order_confirmation":
		return "Dear {UserName},\n\nYour car rental order has been confirmed. Order details..."
	case "order_cancelled":
		return "Dear {UserName},\n\nYour car rental order has been cancelled as requested..."
	case "payment_received":
		return "Dear {UserName},\n\nWe've received your payment for the car rental order..."
	case "rental_reminder":
		return "Dear {UserName},\n\nThis is a reminder about your upcoming car rental..."
	default:
		return "Dear {UserName},\n\nThank you for using Car Rental Service..."
	}
}
