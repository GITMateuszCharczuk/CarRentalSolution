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
	ServicePort     string
	EmailServiceURL string
	FileServiceURL  string
	Env             string
	BlogApiURL      string
	RentalApiURL    string
	IdentityApiURL  string
	JWTToken        string
	SeedCount       struct {
		Users             int
		BlogPosts         int
		CommentsPerPost   int
		Cars              int
		OrdersPerCar      int
		CarImages         int
		Documents         int
		Avatars           int
		EmailTemplates    int
		EmailsPerTemplate int
		StandaloneEmails  int
		BlogImages        int
	}
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
			ServicePort:     getEnv("SERVICE_PORT", "8080"),
			Env:             getEnv("ENV", "test"),
			BlogApiURL:      getEnv("BLOG_API_URL", "http://blog-api:8080"),
			RentalApiURL:    getEnv("RENTAL_API_URL", "http://rental-api:8080"),
			IdentityApiURL:  getEnv("IDENTITY_API_URL", "http://identity-api:8080"),
			EmailServiceURL: getEnv("EMAIL_SERVICE_URL", "http://email-service:8081"),
			FileServiceURL:  getEnv("FILE_SERVICE_URL", "http://file-service:8080"),
			JWTToken:        getEnv("JWT_TOKEN", ""),
		}

		// Set seed counts
		instance.SeedCount.Users = getEnvInt("SEED_USERS_COUNT", 10)
		instance.SeedCount.BlogPosts = getEnvInt("SEED_BLOG_POSTS_COUNT", 10)
		instance.SeedCount.CommentsPerPost = getEnvInt("SEED_COMMENTS_PER_POST", 3)
		instance.SeedCount.Cars = getEnvInt("SEED_CARS_COUNT", 20)
		instance.SeedCount.OrdersPerCar = getEnvInt("SEED_ORDERS_PER_CAR", 2)
		instance.SeedCount.CarImages = getEnvInt("SEED_CAR_IMAGES_COUNT", 30)
		instance.SeedCount.BlogImages = getEnvInt("SEED_BLOG_IMAGES_COUNT", 30)
		instance.SeedCount.Avatars = getEnvInt("SEED_AVATARS_COUNT", 10)
		instance.SeedCount.EmailTemplates = getEnvInt("SEED_EMAIL_TEMPLATES_COUNT", 5)
		instance.SeedCount.EmailsPerTemplate = getEnvInt("SEED_EMAILS_PER_TEMPLATE", 3)
		instance.SeedCount.StandaloneEmails = getEnvInt("SEED_STANDALONE_EMAILS_COUNT", 10)
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
