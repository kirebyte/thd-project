package settings

import (
	"os"
)

// Config struct to hold the environment variables
type Config struct {
	Port           string
	DBPath         string
	CreationScript string
}

// Load reads the environment variables and returns a Config struct
func Load() Config {
	return Config{
		Port:           getEnv("PORT", "8080"),
		DBPath:         getEnv("DB_PATH", "thd-challenge.db"),
		CreationScript: getEnv("DB_CREATION_SCRIPT", "scripts/schema.sql"),
	}
}

// getEnv returns the value of an environment variable or a fallback value
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
