package main

import (
	"log"
	"os"
	"path/filepath"

	"seeder-manager/config"
	"seeder-manager/coordinator"
	"seeder-manager/reference_store"
	"seeder-manager/services/blog"
	"seeder-manager/services/email"
	"seeder-manager/services/file"
	"seeder-manager/services/identity"
	"seeder-manager/services/rental"
)

func main() {
	// Load configuration
	cfg := config.NewConfig(filepath.Join(".", ".env"))
	if cfg.JWTToken == "" {
		log.Fatal("JWT token is required")
	}

	// Initialize reference store
	refStore := reference_store.NewInMemoryStore()

	// Initialize coordinator with all services
	coordinator := coordinator.NewSeedCoordinator(
		identity.NewIdentitySeeder(cfg.IdentityApiURL),
		blog.NewBlogSeeder(cfg.BlogApiURL, refStore),
		rental.NewRentalSeeder(cfg.RentalApiURL, refStore),
		file.NewFileSeeder(refStore, cfg),
		email.NewEmailSeeder(cfg.EmailServiceURL),
		refStore,
	)

	// Run seeding
	log.Println("Starting data seeding...")
	if err := coordinator.SeedAll(cfg.JWTToken); err != nil {
		log.Printf("Error during seeding: %v\n", err)
		os.Exit(1)
	}
	log.Println("Data seeding completed successfully")

	// Run cleanup if needed
	if cfg.Env == "test" {
		log.Println("Running cleanup...")
		if err := coordinator.Cleanup(); err != nil {
			log.Printf("Error during cleanup: %v\n", err)
			os.Exit(1)
		}
		log.Println("Cleanup completed successfully")
	}
}
