package post

import (
	"fmt"
	"time"

	"github.com/google/uuid"
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
	if post.UserID == "" {
		return nil, fmt.Errorf("userID is required")
	}
	if post.Address == "" {
		return nil, fmt.Errorf("address is required")
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
func (s *Service) UpdatePost(id string, req UpdatePostRequest) (*Post, error) {
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

	// Update timestamp
	post.UpdatedAt = time.Now()

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
