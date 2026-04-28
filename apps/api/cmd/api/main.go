package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sanctor/internal/auth"
	"sanctor/internal/bookmark"
	"sanctor/internal/comment"
	"sanctor/internal/community"
	"sanctor/internal/config"
	"sanctor/internal/database"
	"sanctor/internal/dm"
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

func authRequired(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		middleware.Authenticate(handler).ServeHTTP(w, r)
	}
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

			// Run auto-migration for all models
			if err := db.AutoMigrate(&user.User{}, &community.Community{}, &community.UserCommunity{}, &community.CommunityInstitution{}, &post.Post{}, &post.PostCommunity{}, &post.PostInstitution{}, &bookmark.Bookmark{}, &comment.Comment{}, &picture.Picture{}, &institution.Institution{}); err != nil {
				log.Printf("⚠️  Failed to migrate database: %v", err)
			}

			// Apply SQL migrations after GORM has created base tables
			db.RunSQLMigrations()

			log.Println("Initializing modules with database...")
			user.InitWithDatabase(db)
			community.InitWithDatabase(db)
			institution.InitWithDatabase(db)
			bookmark.InitWithDatabase(db)
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
	http.HandleFunc("/api/users", authRequired(user.GetUsers))
	http.HandleFunc("/api/users/get", authRequired(user.GetUser))
	http.HandleFunc("/api/users/create", authRequired(authHandler.Register))
	http.HandleFunc("/api/users/update", authRequired(user.UpdateUser))
	http.HandleFunc("/api/users/delete", authRequired(user.DeleteUser))

	// Community endpoints
	http.HandleFunc("/api/communities", authRequired(community.GetGroups))
	http.HandleFunc("/api/communities/get", authRequired(community.GetGroup))
	http.HandleFunc("/api/communities/create", authRequired(community.CreateGroup))
	http.HandleFunc("/api/communities/update", authRequired(community.UpdateGroup))
	http.HandleFunc("/api/communities/delete", authRequired(community.DeleteGroup))

	// Community membership endpoints
	http.HandleFunc("/api/communities/members/add", authRequired(community.AddUserToGroup))
	http.HandleFunc("/api/communities/members/remove", authRequired(community.RemoveUserFromGroup))
	http.HandleFunc("/api/communities/members", authRequired(community.GetGroupMembers))
	http.HandleFunc("/api/users/communities", authRequired(community.GetUserGroups))

	// Community messaging endpoints
	http.HandleFunc("/api/communities/messages/send", authRequired(community.SendGroupMessage))

	// DM endpoints
	http.HandleFunc("/api/dm/groups/direct", authRequired(dm.CreateDirectGroup))
	http.HandleFunc("/api/dm/groups", authRequired(dm.GetUserGroups))
	http.HandleFunc("/api/dm/messages", authRequired(dm.GetMessages))
	http.HandleFunc("/api/dm/messages/send", authRequired(dm.SendMessage))
	http.HandleFunc("/api/dm/ws", dm.HandleWebSocket)

	// Institution endpoints
	http.HandleFunc("/api/institutions", authRequired(institution.GetInstitutions))
	http.HandleFunc("/api/institutions/get", authRequired(institution.GetInstitution))
	http.HandleFunc("/api/institutions/create", authRequired(institution.CreateInstitution))
	http.HandleFunc("/api/institutions/update", authRequired(institution.UpdateInstitution))
	http.HandleFunc("/api/institutions/delete", authRequired(institution.DeleteInstitution))

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

	http.HandleFunc("/api/posts/get", authRequired(postHandler.GetPost))
	http.HandleFunc("/api/posts/create", authRequired(postHandler.CreatePost))
	http.HandleFunc("/posts/", authRequired(postHandler.UpdatePost)) // Updated route for UpdatePost
	http.HandleFunc("/api/posts/delete", authRequired(postHandler.DeletePost))
	http.HandleFunc("/api/posts/bookmarks", authRequired(bookmark.GetBookmarks))
	http.HandleFunc("/api/posts/bookmarks/check", authRequired(bookmark.CheckBookmark))
	http.HandleFunc("/api/posts/bookmarks/create", authRequired(bookmark.CreateBookmark))
	http.HandleFunc("/api/posts/bookmarks/delete", authRequired(bookmark.DeleteBookmark))

	// Comment endpoints
	http.HandleFunc("/api/comments", authRequired(comment.GetComments))
	http.HandleFunc("/api/comments/get", authRequired(comment.GetComment))
	http.HandleFunc("/api/comments/create", authRequired(comment.CreateComment))
	http.HandleFunc("/api/comments/update", authRequired(comment.UpdateComment))
	http.HandleFunc("/api/comments/delete", authRequired(comment.DeleteComment))

	// Initialize shared user service
	var userRepo user.Repository
	if db != nil {
		userRepo = user.NewPostgresRepository(db)
	} else {
		userRepo = user.NewRepository()
	}
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

	rateLimitConfig := middleware.DefaultRateLimitConfig()
	// Keep auth endpoints tighter than general API traffic.
	rateLimitConfig.MaxRequests = 60
	rateLimited := middleware.RateLimitWithRedis(redisClient, rateLimitConfig)(http.DefaultServeMux)
	handler := middleware.Logger(middleware.CORS(rateLimited))

	fmt.Printf("Server starting on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
