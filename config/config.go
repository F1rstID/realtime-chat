package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type DatabaseConfig struct {
	DSN string
}

type Config struct {
	ServerURL  string
	ServerPort string
	Database   DatabaseConfig
	JWTSecret  string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("failed to load .env file")
	}
	return &Config{
		ServerURL:  getEnv("SERVER_URL", "localhost"),
		ServerPort: getEnv("SERVER_PORT", "5000"),
		Database: DatabaseConfig{
			DSN: getEnv("DATABASE_DSN", "sqlite.db"),
		},
		JWTSecret: getEnv("JWT_SECRET", "test"),
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
