package picture

import (
	"time"

	"github.com/google/uuid"
)

// Picture represents a picture in the system
type Picture struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	PostID    uuid.UUID `json:"postId" gorm:"type:uuid;not null;index"` // Foreign key to Post
	URL       string    `json:"url" gorm:"type:varchar(500);not null"`
	Caption   string    `json:"caption" gorm:"type:text"`
	Order     int       `json:"order" gorm:"default:0"` // Order of picture in the post
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
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
