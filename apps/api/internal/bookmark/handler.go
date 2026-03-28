package bookmark

import (
	"encoding/json"
	"net/http"

	"sanctor/internal/database"
)

// Initialize repository and service (defaults to in-memory).
var (
	repo    Repository = NewRepository()
	service            = NewService(repo)
)

// InitWithDatabase initializes the bookmark module with a database connection.
func InitWithDatabase(db *database.DB) {
	repo = NewPostgresRepository(db)
	service = NewService(repo)
}

// GetBookmarks returns bookmarks for one user.
func GetBookmarks(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	userID := r.URL.Query().Get("userId")
	posts, err := service.GetBookmarkedPostsByUser(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(posts)
}

// CheckBookmark returns whether a user bookmarked a specific post.
func CheckBookmark(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	userID := r.URL.Query().Get("userId")
	postID := r.URL.Query().Get("postId")
	isBookmarked, err := service.IsBookmarked(userID, postID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(BookmarkStatusResponse{
		UserID:       userID,
		PostID:       postID,
		IsBookmarked: isBookmarked,
	})
}

// CreateBookmark creates a new bookmark.
func CreateBookmark(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var req CreateBookmarkRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	bookmark, err := service.CreateBookmark(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(bookmark)
}

// DeleteBookmark deletes an existing bookmark.
func DeleteBookmark(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "DELETE" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID := r.URL.Query().Get("userId")
	postID := r.URL.Query().Get("postId")
	if err := service.DeleteBookmark(userID, postID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}
