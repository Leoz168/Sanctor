package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sanctor/internal/group"
	"sanctor/internal/user"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	w.Header().Set("Content-Type", "application/json")

	response := Response{
		Message: "Sanctor API is running",
		Status:  "healthy",
	}

	json.NewEncoder(w).Encode(response)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Load environment variables (e.g., database credentials)
	dsn := os.Getenv("DATABASE_URL") // Example: "host=localhost user=postgres password=secret dbname=mydb port=5432 sslmode=disable"
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	// Initialize GORM with PostgreSQL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Example: AutoMigrate your models
	// db.AutoMigrate(&YourModel{})

	log.Println("Database connection established successfully!")

	// Health check endpoints
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/api/health", healthHandler)

	// User endpoints
	http.HandleFunc("/api/users", user.GetUsers)
	http.HandleFunc("/api/users/get", user.GetUser)
	http.HandleFunc("/api/users/create", user.CreateUser)
	http.HandleFunc("/api/users/update", user.UpdateUser)
	http.HandleFunc("/api/users/delete", user.DeleteUser)

	// Group endpoints
	http.HandleFunc("/api/groups", group.GetGroups)
	http.HandleFunc("/api/groups/get", group.GetGroup)
	http.HandleFunc("/api/groups/create", group.CreateGroup)
	http.HandleFunc("/api/groups/update", group.UpdateGroup)
	http.HandleFunc("/api/groups/delete", group.DeleteGroup)

	// Group membership endpoints
	http.HandleFunc("/api/groups/members/add", group.AddUserToGroup)
	http.HandleFunc("/api/groups/members/remove", group.RemoveUserFromGroup)
	http.HandleFunc("/api/groups/members", group.GetGroupMembers)
	http.HandleFunc("/api/users/groups", group.GetUserGroups)

	// Group messaging endpoints
	http.HandleFunc("/api/groups/messages/send", group.SendGroupMessage)

	fmt.Printf("Server starting on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}