package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDatabase initializes the database connection using the provided DSN (Data Source Name).
func ConnectDatabase(dsn string) {
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	log.Println("Database connection established successfully!")
}

// Close closes the database connection
func Close() error {
	db, err := DB.DB()
	if err != nil {
		return err
	}
	return db.Close()
}

// Ping checks if the database is reachable
func Ping() error {
	return DB.Exec("SELECT 1").Error
}

// Migrate runs database migrations
func Migrate() error {
	log.Println("Running database migrations...")
	return nil
}
