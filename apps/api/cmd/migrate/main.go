package main

import (
	"log"
	"os"
	"sanctor/internal/database"
	"sanctor/internal/post"
)

func main() {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is required")
	}

	// Initialize the database connection
	db, err := database.NewFromURL(databaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Run the migration
	err = db.AutoMigrate(&post.Post{})
	if err != nil {
		log.Fatalf("Failed to run migration: %v", err)
	}

	log.Println("Migration completed successfully!")
}
