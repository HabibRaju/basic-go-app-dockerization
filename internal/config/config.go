package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds all the configuration for the application
type Config struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
}

// LoadConfig loads the configuration from environment variables
func LoadConfig() *Config {
	// Load .env file (optional, for local development)
	//if err := godotenv.Load(); err != nil {
	//	log.Println("No .env file found, relying on environment variables")
	//}

	if err := godotenv.Load(); err != nil {
		log.Printf("Error loading .env file: %v\n", err)
	} else {
		log.Println(".env file loaded successfully")
	}
	return &Config{
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "root"),
		DBHost:     getEnv("DB_HOST", "mysql-5"),
		DBPort:     getEnv("DB_PORT", "3306"),
		DBName:     getEnv("DB_NAME", "go"),
	}
}

// getEnv retrieves an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
