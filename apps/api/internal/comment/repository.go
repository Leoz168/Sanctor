package comment

import (
	"errors"
	"sort"
	"time"
)

// InMemoryRepository handles comment persistence in memory.
type InMemoryRepository struct {
	comments map[string]*Comment
}

// NewRepository creates a new in-memory comment repository.
func NewRepository() Repository {
	return &InMemoryRepository{comments: make(map[string]*Comment)}
}

func (r *InMemoryRepository) Create(comment *Comment) error {
	if comment == nil {
		return errors.New("comment cannot be nil")
	}
	r.comments[comment.ID.String()] = comment
	return nil
}

func (r *InMemoryRepository) FindByID(id string) (*Comment, error) {
	comment, exists := r.comments[id]
	if !exists || comment.DeletedAt != nil {
		return nil, errors.New("comment not found")
	}
	return comment, nil
}

func (r *InMemoryRepository) FindByPostID(postID string) []*Comment {
	result := make([]*Comment, 0)
	for _, comment := range r.comments {
		if comment.PostID.String() == postID && comment.DeletedAt == nil {
			result = append(result, comment)
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].CreatedAt.Before(result[j].CreatedAt)
	})

	return result
}

func (r *InMemoryRepository) Update(comment *Comment) error {
	if comment == nil {
		return errors.New("comment cannot be nil")
	}
	existing, exists := r.comments[comment.ID.String()]
	if !exists || existing.DeletedAt != nil {
		return errors.New("comment not found")
	}
	r.comments[comment.ID.String()] = comment
	return nil
}

func (r *InMemoryRepository) Delete(id string) error {
	comment, exists := r.comments[id]
	if !exists || comment.DeletedAt != nil {
		return errors.New("comment not found")
	}
	now := time.Now()
	comment.DeletedAt = &now
	comment.UpdatedAt = now
	return nil
}

// ExistsPost returns true in memory mode because posts are not tracked in this repository.
func (r *InMemoryRepository) ExistsPost(postID string) bool {
	return postID != ""
}

// ExistsUser returns true in memory mode because users are not tracked in this repository.
func (r *InMemoryRepository) ExistsUser(userID string) bool {
	return userID != ""
}
