package config

import (
	"log"
	"os"
	"strings"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	StreamName      string
	NatsUrl         string
	StreamSubjects  string
	MailhogHost     string
	Env             string
	MailhogPort     string
	MailhogUsername string
	MailhogPassword string
	MailhogUrl      string
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
			NatsUrl:         getEnv("NATS_URL", "nats://localhost:4222"),
			StreamName:      getEnv("STREAM_NAME", "file_stream"),
			StreamSubjects:  getEnv("STREAM_SUBJECTS", "events.*"),
			Env:             getEnv("ENV", "test"),
			MailhogHost:     getEnv("MAILHOG_HOST", "mailhog"),
			MailhogPort:     getEnv("MAILHOG_PORT", "1025"),
			MailhogUsername: getEnv("MAILHOG_USERNAME", ""),
			MailhogPassword: getEnv("MAILHOG_PASSWORD", ""),
			MailhogUrl:      getEnv("MAILHOG_URL", "localhost:8025"),
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
