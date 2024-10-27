package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoDBUrl      string
	MongoDBName     string
	MongoDBCollName string
	StreamName      string
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
			MongoDBUrl:      getEnv("MONGO_DB_URL", "mongodb://localhost:27017"),
			MongoDBName:     getEnv("MONGO_DB_NAME", "FileDB"),
			MongoDBCollName: getEnv("MONGO_DB_COLLECTION_NAME", "files"),
			StreamName:      getEnv("STREAM_NAME", "default-stream"),
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
