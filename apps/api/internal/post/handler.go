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
	w.Header().Set("Access-Control-Allow-Origin", "*")

	posts, err := h.service.GetAllPosts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(posts)
}

// CreatePost creates a new post
func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	var post Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Map Post to CreatePostRequest
	req := &CreatePostRequest{
		UserID:        post.UserID,
		Address:       &post.Address,
		IsSublet:      &post.IsSublet,
		Price:         &post.Price,
		Rooms:         &post.Rooms,
		RoomsOccupied: &post.RoomsOccupied,
		Bathrooms:     &post.Bathrooms,
		Description:   &post.Description,
		Gender:        &post.Gender,
		PropertyType:  &post.PropertyType,
		Term:          (*Term)(&post.Term),
	}

	createdPost, err := h.service.CreatePost(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdPost)
}

// GetPost retrieves a single post by ID
func (h *Handler) GetPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "id parameter is required", http.StatusBadRequest)
		return
	}

	post, err := h.service.GetPost(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if post == nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(post)
}

// UpdatePost updates an existing post
func (h *Handler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract the ID from the URL
	id := r.URL.Path[len("/posts/"):] // Adjusted for the correct route
	if id == "" {
		http.Error(w, "Post ID is required", http.StatusBadRequest)
		return
	}

	var req UpdatePostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID := r.Header.Get("userID")     // Replace with actual user ID extraction logic
	userRole := r.Header.Get("userRole") // Replace with actual role extraction logic

	// Update the UpdatePost call to handle both return values
	updatedPost, err := h.service.UpdatePost(id, req, userID, userRole)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	// Return the updated post in the response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedPost)
}

// DeletePost deletes a post
func (h *Handler) DeletePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "id parameter is required", http.StatusBadRequest)
		return
	}

	if err := h.service.DeletePost(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
