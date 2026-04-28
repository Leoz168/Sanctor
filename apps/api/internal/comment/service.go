package comment

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

// Service handles business logic for comments.
type Service struct {
	repo Repository
}

// NewService creates a new comment service.
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// CreateComment creates a new comment under a post.
func (s *Service) CreateComment(req CreateCommentRequest) (*Comment, error) {
	if req.PostID == "" {
		return nil, errors.New("post ID is required")
	}
	postUUID, err := uuid.Parse(req.PostID)
	if err != nil {
		return nil, errors.New("invalid post ID format")
	}
	if req.CreatedByUserID == "" {
		return nil, errors.New("created by user ID is required")
	}
	userUUID, err := uuid.Parse(req.CreatedByUserID)
	if err != nil {
		return nil, errors.New("invalid created by user ID format")
	}
	content := strings.TrimSpace(req.Content)
	if content == "" {
		return nil, errors.New("content is required")
	}
	if len(content) > 2000 {
		return nil, errors.New("content must be 2000 characters or fewer")
	}

	if !s.repo.ExistsPost(req.PostID) {
		return nil, errors.New("post not found")
	}
	if !s.repo.ExistsUser(req.CreatedByUserID) {
		return nil, errors.New("user not found")
	}

	now := time.Now()
	comment := &Comment{
		ID:              uuid.New(),
		PostID:          postUUID,
		CreatedByUserID: userUUID,
		Content:         content,
		CreatedAt:       now,
		UpdatedAt:       now,
	}

	if err := s.repo.Create(comment); err != nil {
		return nil, err
	}
	return comment, nil
}

// GetComment retrieves one comment by ID.
func (s *Service) GetComment(id string) (*Comment, error) {
	if id == "" {
		return nil, errors.New("comment ID is required")
	}
	return s.repo.FindByID(id)
}

// GetCommentsByPost retrieves comments for one post.
func (s *Service) GetCommentsByPost(postID string) ([]*Comment, error) {
	if postID == "" {
		return nil, errors.New("post ID is required")
	}
	if _, err := uuid.Parse(postID); err != nil {
		return nil, errors.New("invalid post ID format")
	}
	return s.repo.FindByPostID(postID), nil
}

// UpdateComment updates comment content.
func (s *Service) UpdateComment(id string, req UpdateCommentRequest) (*Comment, error) {
	comment, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("comment not found")
	}

	content := strings.TrimSpace(req.Content)
	if content == "" {
		return nil, errors.New("content is required")
	}
	if len(content) > 2000 {
		return nil, errors.New("content must be 2000 characters or fewer")
	}

	comment.Content = content
	comment.UpdatedAt = time.Now()

	if err := s.repo.Update(comment); err != nil {
		return nil, err
	}
	return comment, nil
}

// DeleteComment soft deletes a comment.
func (s *Service) DeleteComment(id string) error {
	if id == "" {
		return errors.New("comment ID is required")
	}
	return s.repo.Delete(id)
}
