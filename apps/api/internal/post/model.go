package post

import (
	"time"
)

// Term represents the lease term season
type Term string

const (
	TermWinter Term = "Winter"
	TermSpring Term = "Spring"
	TermSummer Term = "Summer"
	TermFall   Term = "Fall"
)

// Model represents a post in the system
type Post struct {
	ID            string    `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID        string    `json:"userId" gorm:"type:uuid;not null;index"`
	Address       string    `json:"address" gorm:"type:varchar(500);not null"`
	IsSublet      bool      `json:"isSublet" gorm:"default:false"`
	Price         string    `json:"price" gorm:"type:varchar(50)"`
	Rooms         string    `json:"bedrooms" gorm:"type:varchar(20)"`
	RoomsOccupied int       `json:"roomsOccupied" gorm:"default:0"`
	Bathrooms     string    `json:"bathrooms" gorm:"type:varchar(20)"`
	Description   string    `json:"description" gorm:"type:text"`
	Gender        string    `json:"gender" gorm:"type:varchar(20)"`
	PropertyType  string    `json:"propertyType" gorm:"type:varchar(50)"`
	Term          Term      `json:"terms" gorm:"type:varchar(20)"`
	CreatedAt     time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt     time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

// CreatePostRequest represents post creation data
type CreatePostRequest struct {
	UserID  string `json:"userId"`
	Title   string `json:"title"`
	Content string `json:"content"`
	PropertyType string `json:"propertyType"`
}

// UpdatePostRequest represents post update data
type UpdatePostRequest struct {
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
	Status  string `json:"status,omitempty"`
}
