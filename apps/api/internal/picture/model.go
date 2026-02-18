package picture

import "time"

// Model represents a picture in the system
type Model struct {
	ID        string    `json:"id"`
	PostID    string    `json:"postId"`    // Foreign key to Post
	URL       string    `json:"url"`
	Caption   string    `json:"caption"`
	Order     int       `json:"order"`     // Order of picture in the post
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// CreatePictureRequest represents picture creation data
type CreatePictureRequest struct {
	PostID  string `json:"postId"`
	URL     string `json:"url"`
	Caption string `json:"caption,omitempty"`
	Order   int    `json:"order,omitempty"`
}

// UpdatePictureRequest represents picture update data
type UpdatePictureRequest struct {
	URL     string `json:"url,omitempty"`
	Caption string `json:"caption,omitempty"`
	Order   int    `json:"order,omitempty"`
}
