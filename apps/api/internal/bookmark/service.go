package bookmark

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

// Service handles business logic for bookmarks.
type Service struct {
	repo Repository
}

// NewService creates a new bookmark service.
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// CreateBookmark creates a bookmark for one post.
func (s *Service) CreateBookmark(req CreateBookmarkRequest) (*Bookmark, error) {
	userID := strings.TrimSpace(req.UserID)
	postID := strings.TrimSpace(req.PostID)

	if userID == "" {
		return nil, errors.New("user ID is required")
	}
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, errors.New("invalid user ID format")
	}
	if postID == "" {
		return nil, errors.New("post ID is required")
	}
	postUUID, err := uuid.Parse(postID)
	if err != nil {
		return nil, errors.New("invalid post ID format")
	}
	if !s.repo.ExistsUser(userUUID) {
		return nil, errors.New("user not found")
	}
	if !s.repo.ExistsPost(postUUID) {
		return nil, errors.New("post not found")
	}

	exists, err := s.repo.Exists(userUUID, postUUID)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("post is already bookmarked")
	}

	bookmark := &Bookmark{
		UserID:    userUUID,
		PostID:    postUUID,
		CreatedAt: time.Now(),
	}

	if err := s.repo.Create(bookmark); err != nil {
		return nil, err
	}
	return bookmark, nil
}

// DeleteBookmark removes a bookmark for one post.
func (s *Service) DeleteBookmark(userID, postID uuid.UUID) error {
	if userID == uuid.Nil {
		return errors.New("user ID is required")
	}
	if postID == uuid.Nil {
		return errors.New("post ID is required")
	}
	return s.repo.Delete(userID, postID)
}

// GetBookmarksByUser returns bookmarks for one user.
func (s *Service) GetBookmarksByUser(userID uuid.UUID) ([]*Bookmark, error) {
	if userID == uuid.Nil {
		return nil, errors.New("user ID is required")
	}
	return s.repo.FindByUserID(userID)
}

// GetBookmarkedPostsByUser returns full posts bookmarked by one user.
func (s *Service) GetBookmarkedPostsByUser(userID uuid.UUID) ([]*BookmarkedPost, error) {
	if userID == uuid.Nil {
		return nil, errors.New("user ID is required")
	}
	return s.repo.FindPostsByUserID(userID)
}

// IsBookmarked checks whether a post is bookmarked by a user.
func (s *Service) IsBookmarked(userID, postID uuid.UUID) (bool, error) {
	if userID == uuid.Nil {
		return false, errors.New("user ID is required")
	}
	if postID == uuid.Nil {
		return false, errors.New("post ID is required")
	}
	return s.repo.Exists(userID, postID)
}
