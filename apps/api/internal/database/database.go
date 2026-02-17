package database

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// DB represents a database connection
type DB struct {
	*sql.DB
}

// Config holds database configuration
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// New creates a new database connection
func New(config Config) (*DB, error) {
	// Build connection string
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.DBName,
		config.SSLMode,
	)

	return connect(connStr)
}

// NewFromURL creates a new database connection from a DATABASE_URL string
func NewFromURL(databaseURL string) (*DB, error) {
	// Ensure sslmode is set (Supabase requires SSL)
	connStr := ensureSSLMode(databaseURL)
	return connect(connStr)
}

// ensureSSLMode adds sslmode=require to the connection URL if not already set
func ensureSSLMode(databaseURL string) string {
	parsed, err := url.Parse(databaseURL)
	if err != nil {
		// If we can't parse it, just append and hope for the best
		if strings.Contains(databaseURL, "?") {
			return databaseURL + "&sslmode=require"
		}
		return databaseURL + "?sslmode=require"
	}

	query := parsed.Query()
	if query.Get("sslmode") == "" {
		query.Set("sslmode", "require")
		parsed.RawQuery = query.Encode()
	}

	return parsed.String()
}

// connect establishes a database connection with the given connection string
func connect(connStr string) (*DB, error) {
	// Open connection
	sqlDB, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Configure connection pool
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	// Retry connection with backoff (DNS/network may take a moment in containers)
	var pingErr error
	for attempt := 1; attempt <= 3; attempt++ {
		pingErr = sqlDB.Ping()
		if pingErr == nil {
			break
		}
		log.Printf("Database connection attempt %d/3 failed: %v", attempt, pingErr)
		if attempt < 3 {
			time.Sleep(time.Duration(attempt) * 2 * time.Second)
		}
	}

	if pingErr != nil {
		sqlDB.Close()
		return nil, fmt.Errorf("failed to ping database after 3 attempts: %w", pingErr)
	}

	log.Println("✅ Database connection established")

	return &DB{sqlDB}, nil
}

// Close closes the database connection
func (db *DB) Close() error {
	log.Println("Closing database connection...")
	return db.DB.Close()
}

// Ping checks if the database is reachable
func (db *DB) Ping() error {
	return db.DB.Ping()
}