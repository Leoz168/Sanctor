package config

import (
	"os"
	"strconv"
)

// Config holds application configuration
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Redis    RedisConfig
	Auth     AuthConfig
	Supabase SupabaseConfig
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

// RedisConfig holds Redis configuration
type RedisConfig struct {
	URL              string
	Host             string
	Port             int
	Username         string
	Password         string
	DB               int
	SentinelEnabled  bool
	MasterName       string
	SentinelAddrs    string
	SentinelUsername string
	SentinelPassword string
}

// AuthConfig holds authentication configuration
type AuthConfig struct {
	JWTSecret     string
	TokenExpiry   int // in hours
	RefreshExpiry int // in days
}

// SupabaseConfig holds Supabase Storage configuration.
type SupabaseConfig struct {
	URL            string
	ServiceRoleKey string
	StorageBucket string
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
		Redis: RedisConfig{
			URL:              getEnv("REDIS_URL", "redis://localhost:6379/0"),
			Host:             getEnv("REDIS_HOST", "localhost"),
			Port:             getEnvInt("REDIS_PORT", 6379),
			Username:         getEnv("REDIS_USERNAME", ""),
			Password:         getEnv("REDIS_PASSWORD", ""),
			DB:               getEnvInt("REDIS_DB", 0),
			SentinelEnabled:  getEnvBool("REDIS_SENTINEL_ENABLED", false),
			MasterName:       getEnv("REDIS_MASTER_NAME", "mymaster"),
			SentinelAddrs:    getEnv("REDIS_SENTINEL_ADDRS", ""),
			SentinelUsername: getEnv("REDIS_SENTINEL_USERNAME", ""),
			SentinelPassword: getEnv("REDIS_SENTINEL_PASSWORD", ""),
		},
		Auth: AuthConfig{
			JWTSecret:     getEnv("JWT_SECRET", "your-secret-key"),
			TokenExpiry:   getEnvInt("TOKEN_EXPIRY", 24),
			RefreshExpiry: getEnvInt("REFRESH_EXPIRY", 7),
		},
		Supabase: SupabaseConfig{
			URL:            getEnv("SUPABASE_URL", ""),
			ServiceRoleKey: getEnv("SUPABASE_SERVICE_ROLE_KEY", ""),
			StorageBucket: getEnv("SUPABASE_STORAGE_BUCKET", "listing-images"),
		},
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

// getEnvBool gets a boolean environment variable with a fallback default value
func getEnvBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return defaultValue
}
