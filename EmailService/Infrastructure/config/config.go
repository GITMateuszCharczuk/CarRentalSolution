package config

import (
	"log"
	"os"
	"strings"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	ServiceAddress     string
	StreamName         string
	NatsUrl            string
	StreamSubjects     string
	MailhogHost        string
	Env                string
	MailhogPort        string
	MailhogUsername    string
	MailhogPassword    string
	MailhogUrl         string
	ServicePort        string
	DefaultEmailSender string
	IdentityApiUrl     string
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
			NatsUrl:            getEnv("NATS_URL", "nats://localhost:4222"),
			StreamName:         getEnv("STREAM_NAME", "email_stream"),
			StreamSubjects:     getEnv("STREAM_SUBJECTS", "email-events.*"),
			Env:                getEnv("ENV", "test"),
			MailhogHost:        getEnv("MAILHOG_HOST", "mailhog"),
			MailhogPort:        getEnv("MAILHOG_PORT", "8025"),
			MailhogUsername:    getEnv("MAILHOG_USERNAME", ""),
			MailhogPassword:    getEnv("MAILHOG_PASSWORD", ""),
			MailhogUrl:         getEnv("MAILHOG_URL", "localhost:8025"),
			ServiceAddress:     getEnv("SERVICE_ADDRESS", "email-service:8080"),
			ServicePort:        getEnv("SERVICE_PORT", "8080"),
			DefaultEmailSender: getEnv("DEFAULT_EMAIL_SENDER", "test@test.com"),
			IdentityApiUrl:     getEnv("IDENTITY_API_URL", "http://localhost:8080"),
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
