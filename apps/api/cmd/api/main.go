package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sanctor/internal/dm"
	"os"
	"sanctor/internal/auth"
	"sanctor/internal/comment"
	"sanctor/internal/config"
	"sanctor/internal/database"
	"sanctor/internal/group"
	"sanctor/internal/institution"
	"sanctor/internal/middleware"
	"sanctor/internal/picture"
	"sanctor/internal/post"
	redisclient "sanctor/internal/redis"
	"sanctor/internal/user"
	"time"
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
	appConfig := config.Load()
	var err error

	// Initialize database connection if DATABASE_URL is set
	var db *database.DB
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL != "" {
		log.Println("Connecting to database...")
		var err error
		db, err = database.NewFromURL(databaseURL)
		if err != nil {
			log.Printf("⚠️  Failed to connect to database: %v", err)
			log.Println("⚠️  Falling back to in-memory storage")
			db = nil
		} else {
			defer db.Close()

			// Run auto-migration for core models
			if err := db.AutoMigrate(&user.User{}, &group.Group{}, &group.UserGroup{}, &picture.Picture{}, &group.GroupInstitution{}); err != nil {
				log.Printf("⚠️  Failed to migrate core tables: %v", err)
			}

			// Migrate posts separately (existing data may have type issues)
			if err := db.AutoMigrate(&post.Post{}, &post.PostGroup{}, &post.PostInstitution{}, &comment.Comment{}, &institution.Institution{}); err != nil {
				log.Printf("⚠️  Failed to migrate posts table: %v", err)
			}

			// Migrate DM tables independently so they always get created
			if err := db.AutoMigrate(&dm.DMGroup{}, &dm.DMGroupUser{}, &dm.DMMessage{}); err != nil {
				log.Printf("⚠️  Failed to migrate DM tables: %v", err)
			} else {
				log.Println("✅ DM tables migrated successfully")
			}

			// Apply SQL migrations after GORM has created base tables
			db.RunSQLMigrations()

			log.Println("Initializing modules with database...")
			user.InitWithDatabase(db)
			group.InitWithDatabase(db)
			dm.InitWithDatabase(db)
			institution.InitWithDatabase(db)
			comment.InitWithDatabase(db)
			log.Println("✅ Database initialized successfully")
		}
	} else {
		log.Println("⚠️  No DATABASE_URL found, using in-memory storage")
	}

	// Initialize Redis client for shared cache/rate-limit use cases.
	var redisClient *redisclient.Client
	redisClient, err = redisclient.New(appConfig.Redis)
	if err != nil {
		log.Printf("⚠️  Failed to initialize Redis client: %v", err)
	} else {
		defer func() {
			if closeErr := redisClient.Close(); closeErr != nil {
				log.Printf("⚠️  Failed to close Redis client: %v", closeErr)
			}
		}()

		pingCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		if pingErr := redisClient.Ping(pingCtx); pingErr != nil {
			log.Printf("⚠️  Redis is not reachable: %v", pingErr)
		} else {
			log.Println("✅ Redis initialized successfully")
		}
	}

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

	// DM endpoints
	http.HandleFunc("/api/dm/groups/direct", dm.CreateDirectGroup)
	http.HandleFunc("/api/dm/groups", dm.GetUserGroups)
	http.HandleFunc("/api/dm/messages", dm.GetMessages)
	http.HandleFunc("/api/dm/messages/send", dm.SendMessage)
	http.HandleFunc("/api/dm/ws", dm.HandleWebSocket)

	// Institution endpoints
	http.HandleFunc("/api/institutions", institution.GetInstitutions)
	http.HandleFunc("/api/institutions/get", institution.GetInstitution)
	http.HandleFunc("/api/institutions/create", institution.CreateInstitution)
	http.HandleFunc("/api/institutions/update", institution.UpdateInstitution)
	http.HandleFunc("/api/institutions/delete", institution.DeleteInstitution)

	// Post endpoints - use database if available
	var postService *post.Service
	if db != nil {
		postRepo := post.NewGormRepository(db)
		postService = post.NewService(postRepo)
		log.Println("✅ Posts initialized with database")
	} else {
		log.Fatal("⚠️  In-memory storage is no longer supported for posts")
	}
	postHandler := post.NewHandler(postService)
	http.HandleFunc("/api/posts", postHandler.GetPosts)
	http.HandleFunc("/api/posts/search", postHandler.SearchPosts)
	http.HandleFunc("/api/posts/get", postHandler.GetPost)
	http.HandleFunc("/api/posts/create", postHandler.CreatePost)
	http.HandleFunc("/posts/", postHandler.UpdatePost) // Updated route for UpdatePost
	http.HandleFunc("/api/posts/delete", postHandler.DeletePost)

	// Comment endpoints
	http.HandleFunc("/api/comments", comment.GetComments)
	http.HandleFunc("/api/comments/get", comment.GetComment)
	http.HandleFunc("/api/comments/create", comment.CreateComment)
	http.HandleFunc("/api/comments/update", comment.UpdateComment)
	http.HandleFunc("/api/comments/delete", comment.DeleteComment)

	// Initialize shared user service
	userRepo := user.NewRepository()
	userService := user.NewService(userRepo)

	// Auth endpoints
	authRepo := auth.NewRepository()
	authService := auth.NewService(authRepo, userService)
	authHandler := auth.NewHandler(authService)
	http.HandleFunc("/api/auth/register", authHandler.Register)
	http.HandleFunc("/api/auth/login", authHandler.Login)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	defaultRateLimit := middleware.DefaultRateLimitConfig()
	defaultRateLimit.MaxRequests = 120

	routeRateLimits := map[string]middleware.RateLimitConfig{
		"/api/auth": {
			Prefix:      "rl:auth",
			Window:      time.Minute,
			MaxRequests: 20,
			FailOpen:    true,
		},
		"/api/posts/search": {
			Prefix:      "rl:search",
			Window:      time.Minute,
			MaxRequests: 60,
			FailOpen:    true,
		},
	}

	rateLimited := middleware.RateLimitWithRedisByRoute(redisClient, routeRateLimits, defaultRateLimit)(http.DefaultServeMux)
	handler := middleware.Logger(middleware.CORS(rateLimited))

	fmt.Printf("Server starting on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
