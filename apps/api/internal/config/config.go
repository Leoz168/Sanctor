package config

import (
	"os"
	"strconv"
)

// Config holds application configuration
type Config struct {
	Server      ServerConfig
	Database    DatabaseConfig
	Auth        AuthConfig
	DatabaseURL string
}

// ServerConfig holds server-specific configuration
type ServerConfig struct {
	Port string
	Host string
	Env  string
}

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

// AuthConfig holds authentication configuration
type AuthConfig struct {
	JWTSecret     string
	TokenExpiry   int // in hours
	RefreshExpiry int // in days
}

// Load loads configuration from environment variables
func Load() *Config {
	return &Config{
		Server: ServerConfig{
			Port: getEnv("PORT", "8080"),
			Host: getEnv("HOST", "0.0.0.0"),
			Env:  getEnv("GO_ENV", "development"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnvInt("DB_PORT", 5432),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", ""),
			DBName:   getEnv("DB_NAME", "sanctor"),
		},
		Auth: AuthConfig{
			JWTSecret:     getEnv("JWT_SECRET", "your-secret-key"),
			TokenExpiry:   getEnvInt("TOKEN_EXPIRY", 24),
			RefreshExpiry: getEnvInt("REFRESH_EXPIRY", 7),
		},
		DatabaseURL: getEnv("DATABASE_URL", ""),
	}
}

// getEnv gets an environment variable with a fallback default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvInt gets an integer environment variable with a fallback default value
func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
