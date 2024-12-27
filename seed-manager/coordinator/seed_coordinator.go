package coordinator

import (
	"fmt"
	"sync"

	"seeder-manager/reference_store"
	"seeder-manager/services/blog"
	"seeder-manager/services/email"
	"seeder-manager/services/file"
	"seeder-manager/services/identity"
	"seeder-manager/services/rental"
)

type SeedCoordinator struct {
	identitySeeder *identity.IdentitySeeder
	blogSeeder     *blog.BlogSeeder
	rentalSeeder   *rental.RentalSeeder
	fileSeeder     *file.FileSeeder
	emailSeeder    *email.EmailSeeder
	refStore       *reference_store.InMemoryStore
}

func NewSeedCoordinator(identitySeeder *identity.IdentitySeeder, blogSeeder *blog.BlogSeeder, rentalSeeder *rental.RentalSeeder, fileSeeder *file.FileSeeder, emailSeeder *email.EmailSeeder, refStore *reference_store.InMemoryStore) *SeedCoordinator {
	return &SeedCoordinator{
		identitySeeder: identitySeeder,
		blogSeeder:     blogSeeder,
		rentalSeeder:   rentalSeeder,
		fileSeeder:     fileSeeder,
		emailSeeder:    emailSeeder,
		refStore:       refStore,
	}
}

func (c *SeedCoordinator) SeedAll(token string) error {
	// 1. Seed Identity Service first to create users
	err := c.identitySeeder.Seed(c.refStore, token)
	if err != nil {
		return fmt.Errorf("error seeding identity service: %w", err)
	}

	// 2. Seed File Storage Service next (as other services might need files)
	if err := c.fileSeeder.Seed(c.refStore, token); err != nil {
		return fmt.Errorf("error seeding file storage service: %w", err)
	}

	// 3. Seed other services in parallel
	errChan := make(chan error, 3)
	var wg sync.WaitGroup
	wg.Add(3)

	// Seed Blog Service
	go func() {
		defer wg.Done()
		if err := c.blogSeeder.Seed(c.refStore, token); err != nil {
			errChan <- fmt.Errorf("error seeding blog service: %w", err)
		}
	}()

	// Seed Rental Service
	go func() {
		defer wg.Done()
		if err := c.rentalSeeder.Seed(c.refStore, token); err != nil {
			errChan <- fmt.Errorf("error seeding rental service: %w", err)
		}
	}()

	// Seed Email Service
	go func() {
		defer wg.Done()
		if err := c.emailSeeder.Seed(c.refStore, token); err != nil {
			errChan <- fmt.Errorf("error seeding email service: %w", err)
		}
	}()

	// Wait for all goroutines to finish
	wg.Wait()
	close(errChan)

	// Check for any errors
	for err := range errChan {
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *SeedCoordinator) Cleanup() error {
	// Cleanup in reverse order
	if err := c.emailSeeder.Cleanup(); err != nil {
		return fmt.Errorf("error cleaning up email service: %w", err)
	}

	if err := c.rentalSeeder.Cleanup(); err != nil {
		return fmt.Errorf("error cleaning up rental service: %w", err)
	}

	if err := c.blogSeeder.Cleanup(); err != nil {
		return fmt.Errorf("error cleaning up blog service: %w", err)
	}

	if err := c.fileSeeder.Cleanup(); err != nil {
		return fmt.Errorf("error cleaning up file storage service: %w", err)
	}

	if err := c.identitySeeder.Cleanup(); err != nil {
		return fmt.Errorf("error cleaning up identity service: %w", err)
	}

	return nil
}
