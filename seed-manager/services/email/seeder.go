package email

import (
	"fmt"
	"log"
	"seeder-manager/api"
	"seeder-manager/config"
	"seeder-manager/models"
	"seeder-manager/reference_store"
	"seeder-manager/services/email/factories"

	"github.com/go-faker/faker/v4"
)

type EmailSeeder struct {
	emailFactory         *factories.EmailFactory
	emailTemplateFactory *factories.EmailTemplateFactory
	apiClient            *api.APIClient
	cfg                  *config.Config
}

func NewEmailSeeder(apiBaseURL string) *EmailSeeder {
	return &EmailSeeder{
		emailTemplateFactory: factories.NewEmailTemplateFactory(),
		apiClient:            api.NewAPIClient(apiBaseURL),
		cfg:                  config.GetConfig(),
	}
}

type SendEmailRequest struct {
	Subject string `json:"subject"`
	Body    string `json:"body"`
	From    string `json:"from"`
}

type SendInternalEmailRequest struct {
	Subject string `json:"subject"`
	Body    string `json:"body"`
	To      string `json:"to"`
}

func (s *EmailSeeder) sendEmail(email *models.EmailModel, token string) error {
	// For template-based emails, use the internal endpoint
	if email.TemplateID != "" {
		request := SendInternalEmailRequest{
			Subject: email.Subject,
			Body:    email.Body,
			To:      email.ToEmail,
		}
		_, err := s.apiClient.Post("/email-service/api/send-internal-email", request, token)
		return err
	}

	// For regular emails, use the standard endpoint
	request := SendEmailRequest{
		Subject: email.Subject,
		Body:    email.Body,
		From:    email.FromEmail,
	}
	_, err := s.apiClient.Post("/email-service/api/send-email", request, token)
	return err
}

func (s *EmailSeeder) Seed(store *reference_store.InMemoryStore, token string) error {
	log.Printf("Starting to seed %d email templates with %d emails each, and %d standalone emails...",
		s.cfg.SeedCount.EmailTemplates, s.cfg.SeedCount.EmailsPerTemplate, s.cfg.SeedCount.StandaloneEmails)

	s.emailFactory = factories.NewEmailFactory(store)

	// Create email templates first
	templates, err := s.emailTemplateFactory.CreateMany(s.cfg.SeedCount.EmailTemplates)
	if err != nil {
		return err
	}

	log.Printf("Starting to seed template-based emails...")
	// Create some example emails using the templates
	for _, template := range templates {
		// Create emails for each template
		for i := 0; i < s.cfg.SeedCount.EmailsPerTemplate; i++ {
			userEmail := faker.Email()
			_, err := store.GetUserID(userEmail)
			if err != nil {
				continue // Skip if user not found
			}

			email, err := s.emailFactory.CreateWithTemplate(template, userEmail)
			if err != nil {
				return err
			}

			// Send the email through the API
			if err := s.sendEmail(email, token); err != nil {
				return fmt.Errorf("error sending template email: %w", err)
			}
		}
	}
	log.Printf("Completed seeding template-based emails")

	log.Printf("Starting to seed %d standalone emails...", s.cfg.SeedCount.StandaloneEmails)
	// Create some standalone emails
	for i := 0; i < s.cfg.SeedCount.StandaloneEmails; i++ {
		email, err := s.emailFactory.Create()
		if err != nil {
			return err
		}

		// Send the email through the API
		if err := s.sendEmail(email, token); err != nil {
			return fmt.Errorf("error sending standalone email: %w", err)
		}
	}
	log.Printf("Completed seeding standalone emails")

	log.Printf("Completed seeding all emails")
	return nil
}

func (s *EmailSeeder) Cleanup() error {
	// TODO: Implement cleanup logic if needed
	return nil
}
