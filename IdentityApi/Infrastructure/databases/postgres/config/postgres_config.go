package config

import (
	"fmt"
	"identity-api/Infrastructure/databases/postgres/entities"
	"log"
	"sync"

	_ "github.com/lib/pq" // PostgreSQL driver
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	DB *gorm.DB
}

var (
	instance *DatabaseConfig
	once     sync.Once
)

func GetDatabaseConfig(user, password, name, host, port string, runPostgresMigration bool) *DatabaseConfig {
	once.Do(func() {
		instance = &DatabaseConfig{}
		instance.connect(user, password, name, host, port)
		instance.createUserRoleEnumType()
		instance.runMigration(runPostgresMigration)
	})
	return instance
}

func (dc *DatabaseConfig) connect(user, password, name, host, port string) {
	var err error

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, name)

	dc.DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	log.Println("Successfully connected to the PostgreSQL database!")
}

func (dc *DatabaseConfig) runMigration(shouldMigrate bool) {
	if shouldMigrate {
		err := dc.DB.AutoMigrate(&entities.UserEntity{})
		if err != nil {
			log.Fatalf("Error migrating database: %v", err)
		}
		log.Println("Database migrated successfully!")
	} else {
		log.Println("Migrations are not required at this time.")
	}
}

func (dc *DatabaseConfig) createUserRoleEnumType() {
	err := dc.DB.Exec("DO $$ BEGIN IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_role') THEN CREATE TYPE user_role AS ENUM ('user', 'admin', 'superadmin'); END IF; $$;").Error
	if err != nil {
		log.Fatalf("Error creating enum type: %v", err)
	}
}
