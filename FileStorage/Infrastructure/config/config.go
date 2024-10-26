package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoDBURL string
	StreamName string
}

var (
	instance *Config
	once     sync.Once
)

func NewConfig(path string) (*Config, error) {
	var err error
	once.Do(func() {
		err = godotenv.Load(path)
		if err != nil {
			log.Printf("Warning: could not load .env file: %v", err)
		}

		instance = &Config{
			MongoDBURL: getEnv("MONGO_DB_URL", "mongodb://localhost:27017"),
			StreamName: getEnv("STREAM_NAME", "default-stream"),
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
