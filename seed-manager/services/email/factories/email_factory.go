package factories

import (
	"fmt"
	"math/rand"
	"time"

	"seeder-manager/models"
	"seeder-manager/reference_store"

	"github.com/google/uuid"
)

type EmailFactory struct {
	refStore *reference_store.InMemoryStore
}

func NewEmailFactory(refStore *reference_store.InMemoryStore) *EmailFactory {
	return &EmailFactory{
		refStore: refStore,
	}
}

func (f *EmailFactory) getRandomUserEmail() (string, error) {
	userEmails := f.refStore.GetAllUserEmails()
	if len(userEmails) == 0 {
		return "", fmt.Errorf("no users available in reference store")
	}
	return userEmails[rand.Intn(len(userEmails))], nil
}

func (f *EmailFactory) Create() (*models.EmailModel, error) {
	// Get a random user for the ToEmail
	userEmail, err := f.getRandomUserEmail()
	if err != nil {
		return nil, fmt.Errorf("error getting random user: %w", err)
	}

	_, err = f.refStore.GetUserID(userEmail)
	if err != nil {
		return nil, fmt.Errorf("error getting user ID: %w", err)
	}

	return &models.EmailModel{
		ID:        uuid.New().String(),
		Subject:   "Welcome to Car Rental Service",
		Body:      "Thank you for joining our car rental service. We hope you enjoy your experience!",
		FromEmail: "noreply@carrentalservice.com",
		ToEmail:   userEmail,
		Status:    "pending",
		CreatedAt: time.Now(),
		SentAt:    nil,
	}, nil
}

func (f *EmailFactory) CreateWithTemplate(template *models.EmailTemplateModel, toEmail string) (*models.EmailModel, error) {
	return &models.EmailModel{
		ID:         uuid.New().String(),
		Subject:    template.Subject,
		Body:       template.Body,
		FromEmail:  "noreply@carrentalservice.com",
		ToEmail:    toEmail,
		Status:     "pending",
		CreatedAt:  time.Now(),
		SentAt:     nil,
		TemplateID: template.ID,
	}, nil
}
