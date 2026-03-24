package bookmark

import (
	"errors"
	"fmt"
	"sort"
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

func (r *InMemoryRepository) Delete(userID, postID string) error {
	key := bookmarkKey(userID, postID)
	delete(r.bookmarks, key)
	return nil
}

func (r *InMemoryRepository) FindByUserID(userID string) ([]*Bookmark, error) {
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

func (r *InMemoryRepository) FindPostsByUserID(userID string) ([]*BookmarkedPost, error) {
	// In-memory mode does not track post entities in this repository.
	return []*BookmarkedPost{}, nil
}

func (r *InMemoryRepository) Exists(userID, postID string) (bool, error) {
	_, exists := r.bookmarks[bookmarkKey(userID, postID)]
	return exists, nil
}

// ExistsPost returns true in memory mode because posts are not tracked in this repository.
func (r *InMemoryRepository) ExistsPost(postID string) bool {
	return postID != ""
}

// ExistsUser returns true in memory mode because users are not tracked in this repository.
func (r *InMemoryRepository) ExistsUser(userID string) bool {
	return userID != ""
}

func bookmarkKey(userID, postID string) string {
	return fmt.Sprintf("%s:%s", userID, postID)
}
