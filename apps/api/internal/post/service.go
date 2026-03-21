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
	CreateWithLinks(post *Post, groupIDs []string, institutionIDs []string) (*Post, error)
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
	return nil
}

// CreatePost creates a new post
func (s *Service) CreatePost(req *CreatePostRequest) (*Post, error) {
	post := &Post{
		ID:     uuid.New().String(),
		UserID: req.UserID,
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
		post.Term = string(*req.Term)
	}

	// Set timestamps
	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()

	groupIDs := uniqueIDs(req.GroupIDs)
	institutionIDs := uniqueIDs(req.InstitutionIDs)

	return s.repo.CreateWithLinks(post, groupIDs, institutionIDs)
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
	if userRole != "admin" && post.CreatedBy != userID {
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
		post.Term = string(*req.Term)
	}

	// Update metadata fields
	post.UpdatedBy = userID
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
