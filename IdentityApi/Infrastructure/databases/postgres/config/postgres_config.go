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

type PostgresDatabase struct {
	DB *gorm.DB
}

var (
	instance *PostgresDatabase
	once     sync.Once
)

func NewPostgresConfig(user, password, name, host, port string, runPostgresMigration bool) *PostgresDatabase {
	once.Do(func() {
		instance = &PostgresDatabase{}
		instance.connect(user, password, name, host, port)
		instance.createUserRoleEnumType()
		instance.runMigration(runPostgresMigration)
	})
	return instance
}

func (dc *PostgresDatabase) connect(user, password, name, host, port string) {
	var err error

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, name)

	dc.DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	log.Println("Successfully connected to the PostgreSQL database!")
}

func (dc *PostgresDatabase) runMigration(shouldMigrate bool) {
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

func (dc *PostgresDatabase) createUserRoleEnumType() {
	err := dc.DB.Exec(`DO $$ 
	BEGIN 
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_role') THEN 
			CREATE TYPE user_role AS ENUM ('user', 'admin', 'superadmin');
		END IF;
	END $$;`).Error
	if err != nil {
		log.Fatalf("Error creating enum type: %v", err)
	}
}
