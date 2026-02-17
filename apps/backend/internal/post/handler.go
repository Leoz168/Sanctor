package post

import (
	"encoding/json"
	"net/http"
)

// Handler handles HTTP requests for posts
type Handler struct {
	service *Service
}

// NewHandler creates a new post handler
func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// GetPosts returns all posts
func (h *Handler) GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// TODO: Implement get posts
	json.NewEncoder(w).Encode(map[string]string{"message": "Get posts - TODO"})
}

// CreatePost creates a new post
func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	// TODO: Implement create post
	json.NewEncoder(w).Encode(map[string]string{"message": "Create post - TODO"})
}
