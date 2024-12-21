package config

import (
	"log"
	"os"
	"strings"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoDBUrl      string
	MongoDBName     string
	MongoDBCollName string
	StreamName      string
	NatsUrl         string
	StreamSubjects  string
	ServerAddress   string
	ServiceAddress  string
	IdentityApiUrl  string
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
			MongoDBUrl:      getEnv("MONGO_DB_URL", "mongodb://localhost:27017"),
			MongoDBName:     getEnv("MONGO_DB_NAME", "FileDB"),
			MongoDBCollName: getEnv("MONGO_DB_COLLECTION_NAME", "files"),
			NatsUrl:         getEnv("NATS_URL", "nats://localhost:4222"),
			StreamName:      getEnv("STREAM_NAME", "file_stream"),
			StreamSubjects:  getEnv("STREAM_SUBJECTS", "file-events.*"),
			ServerAddress:   getEnv("SERVER_ADDRESS", ":8081"),
			ServiceAddress:  getEnv("SERVICE_ADDRESS", "/file-storage/api"),
			IdentityApiUrl:  getEnv("IDENTITY_API_URL", "http://localhost:8080"),
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
