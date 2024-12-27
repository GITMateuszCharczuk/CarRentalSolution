package identity

import (
	"fmt"
	"log"
	"seeder-manager/api"
	"seeder-manager/config"
	"seeder-manager/models"
	"seeder-manager/reference_store"
	"seeder-manager/services/identity/factories"
)

type IdentitySeeder struct {
	userFactory *factories.UserFactory
	apiClient   *api.APIClient
	cfg         *config.Config
}

func NewIdentitySeeder(apiBaseURL string) *IdentitySeeder {
	return &IdentitySeeder{
		userFactory: factories.NewUserFactory(),
		apiClient:   api.NewAPIClient(apiBaseURL),
		cfg:         config.GetConfig(),
	}
}

type RegisterUserRequest struct {
	EmailAddress string `json:"email_address"`
	Password     string `json:"password"`
	Name         string `json:"name"`
	Surname      string `json:"surname"`
	PhoneNumber  string `json:"phone_number"`
	Address      string `json:"address"`
	City         string `json:"city"`
	PostalCode   string `json:"postal_code"`
}

func (s *IdentitySeeder) createUser(user *models.UserModel) error {
	// Generate a secure password for the test user
	password := fmt.Sprintf("Test123!%s", user.ID[:8])

	request := RegisterUserRequest{
		EmailAddress: user.Email,
		Password:     password,
		Name:         user.FirstName,
		Surname:      user.LastName,
		PhoneNumber:  "+1234567890", // Default test phone number
		Address:      "123 Test St", // Default test address
		City:         "Test City",
		PostalCode:   "12345",
	}

	_, err := s.apiClient.Post("/identity-api/api/register", request, "")
	if err != nil {
		return fmt.Errorf("error registering user: %w", err)
	}

	return nil
}

func (s *IdentitySeeder) Seed(store *reference_store.InMemoryStore, token string) error {
	log.Printf("Starting to seed %d users...", s.cfg.SeedCount.Users)

	users, err := s.userFactory.CreateMany(s.cfg.SeedCount.Users)
	if err != nil {
		return err
	}

	// Create users and store references
	for _, user := range users {
		// Create user in the Identity API
		err := s.createUser(user)
		if err != nil {
			return err
		}

		// Store user reference for other services
		if err := store.StoreUserID(user.Email, user.ID); err != nil {
			return err
		}
	}

	log.Printf("Completed seeding %d users", s.cfg.SeedCount.Users)
	return nil
}

func (s *IdentitySeeder) Cleanup() error {
	// TODO: Implement cleanup logic if needed
	return nil
}
