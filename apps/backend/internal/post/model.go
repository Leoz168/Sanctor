package post

import "time"

// Model represents a post in the system
type Model struct {
	ID        string    `json:"id"`
	UserID    string    `json:"userId"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Status    string    `json:"status"` // draft, published, archived
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// CreatePostRequest represents post creation data
type CreatePostRequest struct {
	UserID  string `json:"userId"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// UpdatePostRequest represents post update data
type UpdatePostRequest struct {
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
	Status  string `json:"status,omitempty"`
}
