package post

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"context"

	"sanctor/internal/events"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// PostRepository defines the methods required for a post repository
type PostRepository interface {
	FindByID(id uuid.UUID) (*Post, error)
	FindAll() ([]*Post, error)
	Search(filters PostSearchFilters) ([]*Post, error)
	CreateWithLinks(post *Post, groupIDs []uuid.UUID, institutionIDs []uuid.UUID) (*Post, error)
	Update(post *Post) error
	Delete(id uuid.UUID) error
}

// Service handles business logic for posts
type Service struct {
	repo           PostRepository
	eventPublisher events.EventPublisher
}

// NewService creates a new post service
func NewService(repo PostRepository, eventPublisher events.EventPublisher) *Service {
	return &Service{
		repo:           repo,
		eventPublisher: eventPublisher,
	}
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

	communityIDs, err := parseUUIDs(uniqueIDs(req.CommunityIDs))
	if err != nil {
		return nil, err
	}
	institutionIDs, err := parseUUIDs(uniqueIDs(req.InstitutionIDs))
	if err != nil {
		return nil, err
	}

	createdPost, err := s.repo.CreateWithLinks(post, communityIDs, institutionIDs)
	if err != nil {
		return nil, err
	}
	if s.eventPublisher != nil {
		event := events.PostCreatedEvent{
			BaseEvent: events.BaseEvent{
				EventID:   uuid.New(),
				Timestamp: time.Now(),
				EventType: events.EventTypePostCreated,
			},
			PostID:   createdPost.ID,
			AuthorID: createdPost.UserID,
			Price:    createdPost.Price,
			Gender:   createdPost.Gender,
		}
		if len(institutionIDs) > 0 {
			event.InstitutionIDs = institutionIDs
		}
		if len(communityIDs) > 0 {
			event.CommunityID = communityIDs[0]
		}
		if err := s.eventPublisher.PublishPostCreated(context.Background(), event); err != nil {
			fmt.Printf("failed to publish post created event for post %s: %v\n", createdPost.ID.String(), err)
		}
	}

	return createdPost, nil
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
func (s *Service) GetPost(id uuid.UUID) (*Post, error) {
	if s.repo == nil {
		return nil, fmt.Errorf("repository not initialized")
	}

	post, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if post == nil {
		return nil, fmt.Errorf("post not found")
	}

	if s.eventPublisher != nil {
		event := events.PostViewedEvent{
			BaseEvent: events.BaseEvent{
				EventID:   uuid.New(),
				Timestamp: time.Now(),
				EventType: events.EventTypePostViewed,
			},
			PostID: post.ID,
		}

		if err := s.eventPublisher.PublishPostViewed(context.Background(), event); err != nil {
			return nil, err
		}
	}
	return post, nil
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
func (s *Service) UpdatePost(id uuid.UUID, req UpdatePostRequest, userID uuid.UUID, userRole string) (*Post, error) {
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
	if userRole != "admin" && post.CreatedByUserID != userID {
		return nil, errors.New("you are not allowed to update this post")
	}

	//save old values before mutation
	oldAddress := post.Address
	oldIsSublet := post.IsSublet
	oldPrice := post.Price
	oldRooms := post.Rooms
	oldRoomsOccupied := post.RoomsOccupied
	oldBathrooms := post.Bathrooms
	oldDescription := post.Description
	oldGender := post.Gender
	oldPropertyType := post.PropertyType
	oldTerm := post.Term

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
	if userID == uuid.Nil {
		return nil, errors.New("user ID is required")
	}
	post.UpdatedByUserID = userID
	post.UpdatedAt = time.Now()

	// Validate required fields
	if err := validatePostInput(post); err != nil {
		return nil, err
	}

	// Save to repository
	if err := s.repo.Update(post); err != nil {
		return nil, err
	}

	if s.eventPublisher != nil {
		updatedFields := make(map[string]events.FieldChange)
		if post.Address != oldAddress {
			updatedFields["address"] = events.FieldChange{Old: oldAddress, New: post.Address}
		}
		if post.IsSublet != oldIsSublet {
			updatedFields["isSublet"] = events.FieldChange{Old: oldIsSublet, New: post.IsSublet}
		}
		if post.Price != oldPrice {
			updatedFields["price"] = events.FieldChange{Old: oldPrice, New: post.Price}
		}
		if post.Rooms != oldRooms {
			updatedFields["rooms"] = events.FieldChange{Old: oldRooms, New: post.Rooms}
		}
		if post.RoomsOccupied != oldRoomsOccupied {
			updatedFields["roomsOccupied"] = events.FieldChange{Old: oldRoomsOccupied, New: post.RoomsOccupied}
		}
		if post.Bathrooms != oldBathrooms {
			updatedFields["bathrooms"] = events.FieldChange{Old: oldBathrooms, New: post.Bathrooms}
		}
		if post.Description != oldDescription {
			updatedFields["description"] = events.FieldChange{Old: oldDescription, New: post.Description}
		}
		if post.Gender != oldGender {
			updatedFields["gender"] = events.FieldChange{Old: oldGender, New: post.Gender}
		}
		if post.PropertyType != oldPropertyType {
			updatedFields["propertyType"] = events.FieldChange{Old: oldPropertyType, New: post.PropertyType}
		}
		if post.Term != oldTerm {
			updatedFields["term"] = events.FieldChange{Old: oldTerm, New: post.Term}
		}
		if len(updatedFields) > 0 {
			event := events.PostUpdatedEvent{
				BaseEvent: events.BaseEvent{
					EventID:   uuid.New(),
					Timestamp: time.Now(),
					EventType: events.EventTypePostUpdated,
				},
				PostID:        post.ID,
				AuthorID:      userID,
				UpdatedFields: updatedFields,
			}

			if err := s.eventPublisher.PublishPostUpdated(context.Background(), event); err != nil {
				fmt.Printf("failed to publish post updated event: %v\n", err)
			}
		}
	}

	return post, nil
}

// DeletePost deletes a post
func (s *Service) DeletePost(id uuid.UUID) error {
	if s.repo == nil {
		return fmt.Errorf("repository not initialized")
	}

	post, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	if post == nil {
		return fmt.Errorf("post not found")
	}

	if err := s.repo.Delete(id); err != nil {
		return err
	}

	if s.eventPublisher != nil {
		event := events.PostDeletedEvent{
			BaseEvent: events.BaseEvent{
				EventID:   uuid.New(),
				Timestamp: time.Now(),
				EventType: events.EventTypePostDeleted,
			},
			PostID:   post.ID,
			AuthorID: post.UserID,
		}

		if err := s.eventPublisher.PublishPostDeleted(context.Background(), event); err != nil {
			fmt.Printf("failed to publish post deleted event: %v\n", err)
		}
	}

	return nil
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
