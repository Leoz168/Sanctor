package comment

import (
	"encoding/json"
	"net/http"

	"sanctor/internal/database"

	"github.com/google/uuid"
)

// Initialize repository and service (defaults to in-memory).
var (
	repo    Repository = NewRepository()
	service            = NewService(repo)
)

// InitWithDatabase initializes the comment module with a database connection.
func InitWithDatabase(db *database.DB) {
	repo = NewPostgresRepository(db)
	service = NewService(repo)
}

// GetComments returns all comments for a specific post.
func GetComments(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	postID := r.URL.Query().Get("postId")
	postUUID, err := uuid.Parse(postID)
	if err != nil {
		http.Error(w, "invalid post ID format", http.StatusBadRequest)
		return
	}
	comments, err := service.GetCommentsByPost(postUUID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(comments)
}

// GetComment returns a single comment by ID.
func GetComment(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	id := r.URL.Query().Get("id")
	commentID, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "invalid comment ID format", http.StatusBadRequest)
		return
	}
	comment, err := service.GetComment(commentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(comment)
}

// CreateComment creates a new comment.
func CreateComment(w http.ResponseWriter, r *http.Request) {
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

	var req CreateCommentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	comment, err := service.CreateComment(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(comment)
}

// UpdateComment updates an existing comment.
func UpdateComment(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "PUT" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Comment ID is required", http.StatusBadRequest)
		return
	}

	var req UpdateCommentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	commentID, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "invalid comment ID format", http.StatusBadRequest)
		return
	}
	comment, err := service.UpdateComment(commentID, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(comment)
}

// DeleteComment soft deletes a comment.
func DeleteComment(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "DELETE" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	commentID, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "invalid comment ID format", http.StatusBadRequest)
		return
	}
	if err := service.DeleteComment(commentID); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}
