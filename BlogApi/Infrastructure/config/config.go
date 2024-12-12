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
	RedisHost            string
	RedisPort            string
	RedisPassword        string
	AccessTokenTTL       int
	RefreshTokenTTL      int
	SecretKey            string
	PostgresUser         string
	PostgresPassword     string
	PostgresDBName       string
	PostgresHost         string
	PostgresPort         string
	RunPostgresMigration bool
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
			RedisHost:            getEnv("REDIS_HOST", "localhost"),
			RedisPort:            getEnv("REDIS_PORT", "6379"),
			RedisPassword:        getEnv("REDIS_PASSWORD", ""),
			AccessTokenTTL:       getEnvInt("ACCESS_TOKEN_TTL", 15),
			RefreshTokenTTL:      getEnvInt("REFRESH_TOKEN_TTL", 30),
			SecretKey:            getEnv("SECRET_KEY", ""),
			PostgresUser:         getEnv("POSTGRES_USER", "postgres"),
			PostgresPassword:     getEnv("POSTGRES_PASSWORD", "postgres"),
			PostgresDBName:       getEnv("POSTGRES_DB", "postgres"),
			PostgresHost:         getEnv("POSTGRES_HOST", "localhost"),
			PostgresPort:         getEnv("POSTGRES_PORT", "5432"),
			RunPostgresMigration: getEnvBool("RUN_POSTGRES_MIGRATION", false),
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

func getEnvInt(key string, fallback int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
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
