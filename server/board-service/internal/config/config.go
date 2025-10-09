package config

import (
	"os"
)

type Config struct {
	Port            string
	MongoDBURI      string
	MongoDBDatabase string
	AllowedOrigin   string
}

func Load() *Config {
	return &Config{
		Port:            getEnv("PORT", "8082"),
		MongoDBURI:      getEnv("MONGODB_URI", "mongodb://admin:password@localhost:27017"),
		MongoDBDatabase: getEnv("MONGODB_DATABASE", "boards"),
		AllowedOrigin:   getEnv("ALLOWED_ORIGIN", "http://localhost:3000"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
