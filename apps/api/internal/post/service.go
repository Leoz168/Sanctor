package post

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gin-gonic/gin"
)

// Service handles business logic for post operations
type Service struct {
	repo RepositoryInterface
}

// NewService creates a new post service with in-memory repository
func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

// NewServiceWithGorm creates a new post service with GORM repository
func NewServiceWithGorm(repo *GormRepository) *Service {
	return &Service{repo: repo}
}

// validatePostInput validates the required fields for a post
func validatePostInput(post *Post) error {
	if post.Title == "" {
		return errors.New("title is required")
	}
	if post.Content == "" {
		return errors.New("content is required")
	}
	return nil
}

// CreatePost creates a new post
func (s *Service) CreatePost(post *Post) (*Post, error) {
	// Generate ID if not provided
	if post.ID == "" {
		post.ID = uuid.New().String()
	}

	// Set timestamps
	now := time.Now()
	post.CreatedAt = now
	post.UpdatedAt = now

	// Validate required fields
	if err := validatePostInput(post); err != nil {
		return nil, err
	}

	// If repository exists, save to database
	if s.repo != nil {
		return s.repo.Create(post)
	}

	// Return the post with generated values (in-memory mode)
	return post, nil
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
		post.Term = *req.Term
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
