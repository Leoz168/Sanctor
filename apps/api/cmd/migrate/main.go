package main

import (
	"log"

	"Sanctor/apps/api/internal/database"
	"Sanctor/apps/api/internal/post"
)

func main() {
	// Initialize the database connection
	db, err := database.Connect() // Replace with your actual database connection logic
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Run the migration
	err = db.AutoMigrate(&post.Post{})
	if err != nil {
		log.Fatalf("Failed to run migration: %v", err)
	}

	log.Println("Migration completed successfully!")
}