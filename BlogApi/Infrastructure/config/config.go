package config

import (
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	ServiceAddress       string
	Env                  string
	ServicePort          string
	PostgresUser         string
	PostgresPassword     string
	PostgresDBName       string
	PostgresHost         string
	PostgresPort         string
	RunPostgresMigration bool
	IdentityApiUrl       string
}

var (
	instance *Config
	once     sync.Once
)

func NewConfig(path string) (*Config, error) {
	var err error
	once.Do(func() {
		env := getEnv("ENV", "")
		if strings.ToLower(env) == "test" || env == "" {
			err = godotenv.Load(path)
		}

		if err != nil && strings.ToLower(env) != "prod" {
			log.Printf("Warning: could not load .env file: %v, %v", err, env)
		}

		instance = &Config{
			Env:                  getEnv("ENV", "test"),
			ServiceAddress:       getEnv("SERVICE_ADDRESS", "identity-api:8080"),
			ServicePort:          getEnv("SERVICE_PORT", "8080"),
			PostgresUser:         getEnv("POSTGRES_USER", "postgres"),
			PostgresPassword:     getEnv("POSTGRES_PASSWORD", "postgres"),
			PostgresDBName:       getEnv("POSTGRES_DB", "postgres"),
			PostgresHost:         getEnv("POSTGRES_HOST", "localhost"),
			PostgresPort:         getEnv("POSTGRES_PORT", "5432"),
			RunPostgresMigration: getEnvBool("RUN_POSTGRES_MIGRATION", false),
			IdentityApiUrl:       getEnv("IDENTITY_API_URL", "http://localhost:8080"),
		}
	})
	return instance, err
}

func GetConfig() *Config {
	return instance
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func getEnvBool(key string, fallback bool) bool {
	if value, exists := os.LookupEnv(key); exists {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return fallback
}
