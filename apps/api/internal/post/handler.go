package post

import (
	"encoding/json"
	"net/http"
	sharedtypes "sanctor/pkg/types"
	"strconv"
	"strings"

	"github.com/google/uuid"
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

// SearchPosts returns posts filtered by query parameters.
func (h *Handler) SearchPosts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodOptions {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	filters, err := parsePostSearchFilters(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	posts, err := h.service.SearchPosts(filters)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(posts)
}

func parsePostSearchFilters(r *http.Request) (PostSearchFilters, error) {
	q := r.URL.Query()
	communityID := strings.TrimSpace(q.Get("communityId"))
	if communityID == "" {
		communityID = strings.TrimSpace(q.Get("groupId"))
	}

	filters := PostSearchFilters{
		Query:         strings.TrimSpace(q.Get("q")),
		PropertyType:  strings.TrimSpace(q.Get("propertyType")),
		CommunityID:   communityID,
		InstitutionID: strings.TrimSpace(q.Get("institutionId")),
		SortBy:        strings.TrimSpace(q.Get("sortBy")),
		SortOrder:     strings.TrimSpace(q.Get("sortOrder")),
	}

	if q.Get("minPrice") != "" {
		value, err := strconv.ParseInt(q.Get("minPrice"), 10, 64)
		if err != nil {
			return PostSearchFilters{}, err
		}
		filters.MinPrice = &value
	}
	if q.Get("maxPrice") != "" {
		value, err := strconv.ParseInt(q.Get("maxPrice"), 10, 64)
		if err != nil {
			return PostSearchFilters{}, err
		}
		filters.MaxPrice = &value
	}
	if q.Get("minRooms") != "" {
		value, err := strconv.ParseInt(q.Get("minRooms"), 10, 64)
		if err != nil {
			return PostSearchFilters{}, err
		}
		filters.MinRooms = &value
	}
	if q.Get("minBathrooms") != "" {
		value, err := strconv.ParseInt(q.Get("minBathrooms"), 10, 64)
		if err != nil {
			return PostSearchFilters{}, err
		}
		filters.MinBathrooms = &value
	}
	if q.Get("isSublet") != "" {
		value, err := strconv.ParseBool(q.Get("isSublet"))
		if err != nil {
			return PostSearchFilters{}, err
		}
		filters.IsSublet = &value
	}
	if q.Get("gender") != "" {
		value := sharedtypes.Gender(q.Get("gender"))
		filters.Gender = &value
	}
	if q.Get("term") != "" {
		value := sharedtypes.Term(q.Get("term"))
		filters.Term = &value
	}
	if q.Get("limit") != "" {
		value, err := strconv.Atoi(q.Get("limit"))
		if err != nil {
			return PostSearchFilters{}, err
		}
		filters.Limit = value
	}
	if q.Get("offset") != "" {
		value, err := strconv.Atoi(q.Get("offset"))
		if err != nil {
			return PostSearchFilters{}, err
		}
		filters.Offset = value
	}

	return filters, nil
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

	var req CreatePostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdPost, err := h.service.CreatePost(r.Context(), &req)
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

	postID, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "invalid post ID format", http.StatusBadRequest)
		return
	}
	post, err := h.service.GetPost(r.Context(), postID)
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
	postID, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "invalid post ID format", http.StatusBadRequest)
		return
	}

	var req UpdatePostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID := r.Header.Get("userID")     // Replace with actual user ID extraction logic
	userRole := r.Header.Get("userRole") // Replace with actual role extraction logic
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		http.Error(w, "invalid user ID format", http.StatusBadRequest)
		return
	}

	// Update the UpdatePost call to handle both return values
	updatedPost, err := h.service.UpdatePost(r.Context(), postID, req, userUUID, userRole)
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

	postID, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "invalid post ID format", http.StatusBadRequest)
		return
	}
	if err := h.service.DeletePost(r.Context(), postID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
