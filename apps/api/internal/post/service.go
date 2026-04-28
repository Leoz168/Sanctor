package post

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// PostRepository defines the methods required for a post repository
type PostRepository interface {
	FindByID(id string) (*Post, error)
	FindAll() ([]*Post, error)
	Search(filters PostSearchFilters) ([]*Post, error)
	CreateWithLinks(post *Post, groupIDs []uuid.UUID, institutionIDs []uuid.UUID) (*Post, error)
	Update(post *Post) error
	Delete(id string) error
}

// Service handles business logic for posts
type Service struct {
	repo PostRepository
}

// NewService creates a new post service
func NewService(repo PostRepository) *Service {
	return &Service{repo: repo}
}

// validatePostInput validates the required fields for a post
// Updated validatePostInput to allow partial updates
func validatePostInput(post *Post) error {
	if post.Title != "" && len(post.Title) < 3 {
		return errors.New("title must be at least 3 characters")
	}
	if post.Content != "" && len(post.Content) < 10 {
		return errors.New("content must be at least 10 characters")
	}
	if post.Price < 0 {
		return errors.New("price cannot be negative")
	}
	if post.Rooms < 0 {
		return errors.New("rooms cannot be negative")
	}
	if post.Bathrooms < 0 {
		return errors.New("bathrooms cannot be negative")
	}
	if post.RoomsOccupied < 0 || post.RoomsOccupied > post.Rooms {
		return errors.New("rooms occupied cannot be negative or exceed total rooms")
	}
	return nil
}

// CreatePost creates a new post
func (s *Service) CreatePost(req *CreatePostRequest) (*Post, error) {
	if req.UserID == "" {
		return nil, errors.New("userId is required")
	}
	userUUID, err := uuid.Parse(req.UserID)
	if err != nil {
		return nil, errors.New("invalid userId format")
	}

	post := &Post{
		ID:              uuid.New(),
		UserID:          userUUID,
		CreatedByUserID: userUUID,
		UpdatedByUserID: userUUID,
	}

	if req.Address != nil {
		post.Address = *req.Address
	}
	if req.IsSublet != nil {
		post.IsSublet = *req.IsSublet
	}
	if req.Price != nil {
		post.Price = *req.Price
	}
	if req.Rooms != nil {
		post.Rooms = *req.Rooms
	}
	if req.RoomsOccupied != nil {
		post.RoomsOccupied = *req.RoomsOccupied
	}
	if req.Bathrooms != nil {
		post.Bathrooms = *req.Bathrooms
	}
	if req.Description != nil {
		post.Description = *req.Description
	}
	if req.Gender != nil {
		post.Gender = *req.Gender
	}
	if req.PropertyType != nil {
		post.PropertyType = *req.PropertyType
	}
	if req.Term != nil {
		post.Term = *req.Term
	}

	// Set timestamps
	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()

	groupIDs, err := parseUUIDs(uniqueIDs(req.GroupIDs))
	if err != nil {
		return nil, err
	}
	institutionIDs, err := parseUUIDs(uniqueIDs(req.InstitutionIDs))
	if err != nil {
		return nil, err
	}

	return s.repo.CreateWithLinks(post, groupIDs, institutionIDs)
}

func parseUUIDs(ids []string) ([]uuid.UUID, error) {
	if len(ids) == 0 {
		return nil, nil
	}
	parsed := make([]uuid.UUID, 0, len(ids))
	for _, id := range ids {
		value, err := uuid.Parse(id)
		if err != nil {
			return nil, errors.New("invalid UUID in list")
		}
		parsed = append(parsed, value)
	}
	return parsed, nil
}

func uniqueIDs(ids []string) []string {
	if len(ids) == 0 {
		return nil
	}

	seen := make(map[string]struct{}, len(ids))
	unique := make([]string, 0, len(ids))
	for _, id := range ids {
		if id == "" {
			continue
		}
		if _, exists := seen[id]; exists {
			continue
		}
		seen[id] = struct{}{}
		unique = append(unique, id)
	}

	return unique
}

// GetPost retrieves a post by ID
func (s *Service) GetPost(id string) (*Post, error) {
	if s.repo != nil {
		return s.repo.FindByID(id)
	}
	return nil, fmt.Errorf("post not found")
}

// GetAllPosts retrieves all posts
func (s *Service) GetAllPosts() ([]*Post, error) {
	if s.repo != nil {
		return s.repo.FindAll()
	}
	return []*Post{}, nil
}

// SearchPosts retrieves posts using backend filters, sorting, and pagination.
func (s *Service) SearchPosts(filters PostSearchFilters) ([]*Post, error) {
	if s.repo == nil {
		return nil, fmt.Errorf("repository not initialized")
	}

	if filters.MinPrice != nil && filters.MaxPrice != nil && *filters.MinPrice > *filters.MaxPrice {
		return nil, errors.New("minPrice cannot be greater than maxPrice")
	}

	if filters.Limit <= 0 {
		filters.Limit = 20
	}
	if filters.Limit > 100 {
		filters.Limit = 100
	}
	if filters.Offset < 0 {
		filters.Offset = 0
	}

	return s.repo.Search(filters)
}

// UpdatePost updates an existing post
func (s *Service) UpdatePost(id string, req UpdatePostRequest, userID string, userRole string) (*Post, error) {
	if s.repo == nil {
		return nil, fmt.Errorf("repository not initialized")
	}

	// Get existing post
	post, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if post == nil {
		return nil, fmt.Errorf("post not found")
	}

	// Check if the user is allowed to update the post
	if userRole != "admin" && post.CreatedByUserID.String() != userID {
		return nil, errors.New("you are not allowed to update this post")
	}

	// Update fields if provided (pointer fields are nil when omitted)
	if req.Address != nil {
		post.Address = *req.Address
	}
	if req.IsSublet != nil {
		post.IsSublet = *req.IsSublet
	}
	if req.Price != nil {
		post.Price = *req.Price
	}
	if req.Rooms != nil {
		post.Rooms = *req.Rooms
	}
	if req.RoomsOccupied != nil {
		post.RoomsOccupied = *req.RoomsOccupied
	}
	if req.Bathrooms != nil {
		post.Bathrooms = *req.Bathrooms
	}
	if req.Description != nil {
		post.Description = *req.Description
	}
	if req.Gender != nil {
		post.Gender = *req.Gender
	}
	if req.PropertyType != nil {
		post.PropertyType = *req.PropertyType
	}
	if req.Term != nil {
		post.Term = *req.Term
	}

	// Update metadata fields
	if userID == "" {
		return nil, errors.New("user ID is required")
	}
	updaterUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, errors.New("invalid user ID format")
	}
	post.UpdatedByUserID = updaterUUID
	post.UpdatedAt = time.Now()

	// Validate required fields
	if err := validatePostInput(post); err != nil {
		return nil, err
	}

	// Save to repository
	if err := s.repo.Update(post); err != nil {
		return nil, err
	}

	return post, nil
}

// DeletePost deletes a post
func (s *Service) DeletePost(id string) error {
	if s.repo != nil {
		return s.repo.Delete(id)
	}
	return fmt.Errorf("not implemented")
}

// Add a middleware function to check user roles and permissions
func Authorize(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("userRole")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
			c.Abort()
			return
		}

		role := userRole.(string)
		for _, allowedRole := range allowedRoles {
			if role == allowedRole {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		c.Abort()
	}
}
