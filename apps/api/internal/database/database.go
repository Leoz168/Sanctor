package database

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB represents a database connection
type DB struct {
	*sql.DB
	Gorm *gorm.DB // GORM instance for ORM operations
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
	// Let GORM open the connection using its own pgx driver
	var gormDB *gorm.DB
	var err error

	// Retry connection with backoff (DNS/network may take a moment in containers)
	for attempt := 1; attempt <= 3; attempt++ {
		gormDB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{
			SkipDefaultTransaction: true,
		})
		if err == nil {
			break
		}
		log.Printf("Database connection attempt %d/3 failed: %v", attempt, err)
		if attempt < 3 {
			time.Sleep(time.Duration(attempt) * 2 * time.Second)
		}
	}

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database after 3 attempts: %w", err)
	}

	// Get the underlying sql.DB for connection pool configuration
	sqlDB, err := gormDB.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}

	// Configure connection pool
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	// Verify connectivity
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("✅ Database connection established")
	log.Println("✅ GORM initialized")

	return &DB{DB: sqlDB, Gorm: gormDB}, nil
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

// AutoMigrate runs GORM auto-migration for the given models
func (db *DB) AutoMigrate(models ...interface{}) error {
	log.Println("Running database migrations...")
	if err := db.Gorm.AutoMigrate(models...); err != nil {
		return fmt.Errorf("failed to auto-migrate: %w", err)
	}
	log.Println("✅ Database migrations completed")
	return nil
}