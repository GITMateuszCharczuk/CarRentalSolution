package config

import (
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	ServicePort           string
	EmailServiceURL       string
	FileServiceURL        string
	Env                   string
	RequestSentLimit      int
	RequestSentTimeWindow time.Duration
	RequestSizeLimit      int64
	MainApiRoute          string
	BlogApiURL            string
	RentalApiURL          string
	IdentityApiURL        string
}

var (
	instance *Config
	once     sync.Once
)

func NewConfig(path string) *Config {
	var err error
	once.Do(func() {
		env := getEnv("ENV", "")
		if strings.ToLower(env) == "test" || env == "" {
			err = godotenv.Load(path)
		}

		if err != nil && strings.ToLower(env) != "prod" {
			panic(err)
		}

		instance = &Config{
			ServicePort:           getEnv("SERVICE_PORT", "8080"),
			Env:                   getEnv("ENV", "test"),
			RequestSentLimit:      getEnvInt("REQUEST_SENT_LIMIT", 100),
			RequestSentTimeWindow: getEnvDuration("REQUEST_SENT_TIME_WINDOW", 60*time.Second),
			RequestSizeLimit:      getEnvInt64("REQUEST_SIZE_LIMIT", 10),
			MainApiRoute:          getEnv("MAIN_API_ROUTE", "/car-rental/api"),
			BlogApiURL:            getEnv("BLOG_API_URL", "http://blog-api:8080"),
			RentalApiURL:          getEnv("RENTAL_API_URL", "http://rental-api:8080"),
			IdentityApiURL:        getEnv("IDENTITY_API_URL", "http://identity-api:8080"),
			EmailServiceURL:       getEnv("EMAIL_SERVICE_URL", "http://email-service:8081"),
			FileServiceURL:        getEnv("FILE_SERVICE_URL", "http://file-service:8080"),
		}
	})
	return instance
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

func getEnvInt64(key string, fallback int64) int64 {
	if value, exists := os.LookupEnv(key); exists {
		if int64Value, err := strconv.ParseInt(value, 10, 64); err == nil {
			return int64Value
		}
	}
	return fallback
}

func getEnvDuration(key string, fallback time.Duration) time.Duration {
	if value, exists := os.LookupEnv(key); exists {
		if seconds, err := strconv.Atoi(value); err == nil {
			return time.Duration(seconds) * time.Second
		}
	}
	return fallback
}
