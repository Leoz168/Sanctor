package bookmark

import (
	"errors"
	"fmt"
	"sort"

	"github.com/google/uuid"
)

// InMemoryRepository handles bookmark persistence in memory.
type InMemoryRepository struct {
	bookmarks map[string]*Bookmark
}

// NewRepository creates a new in-memory bookmark repository.
func NewRepository() Repository {
	return &InMemoryRepository{bookmarks: make(map[string]*Bookmark)}
}

func (r *InMemoryRepository) Create(bookmark *Bookmark) error {
	if bookmark == nil {
		return errors.New("bookmark cannot be nil")
	}
	key := bookmarkKey(bookmark.UserID, bookmark.PostID)
	if _, exists := r.bookmarks[key]; exists {
		return errors.New("post is already bookmarked")
	}
	r.bookmarks[key] = bookmark
	return nil
}

func (r *InMemoryRepository) Delete(userID, postID uuid.UUID) error {
	key := bookmarkKey(userID, postID)
	delete(r.bookmarks, key)
	return nil
}

func (r *InMemoryRepository) FindByUserID(userID uuid.UUID) ([]*Bookmark, error) {
	bookmarks := make([]*Bookmark, 0)
	for _, bookmark := range r.bookmarks {
		if bookmark.UserID == userID {
			bookmarks = append(bookmarks, bookmark)
		}
	}

	sort.Slice(bookmarks, func(i, j int) bool {
		return bookmarks[i].CreatedAt.After(bookmarks[j].CreatedAt)
	})

	return bookmarks, nil
}

func (r *InMemoryRepository) FindPostsByUserID(userID uuid.UUID) ([]*BookmarkedPost, error) {
	// In-memory mode does not track post entities in this repository.
	return []*BookmarkedPost{}, nil
}

func (r *InMemoryRepository) Exists(userID, postID uuid.UUID) (bool, error) {
	_, exists := r.bookmarks[bookmarkKey(userID, postID)]
	return exists, nil
}

// ExistsPost returns true in memory mode because posts are not tracked in this repository.
func (r *InMemoryRepository) ExistsPost(postID uuid.UUID) bool {
	return postID != uuid.Nil
}

// ExistsUser returns true in memory mode because users are not tracked in this repository.
func (r *InMemoryRepository) ExistsUser(userID uuid.UUID) bool {
	return userID != uuid.Nil
}

func bookmarkKey(userID, postID uuid.UUID) string {
	return fmt.Sprintf("%s:%s", userID, postID)
}
