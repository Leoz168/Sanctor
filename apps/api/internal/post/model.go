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
	ID            string    `json:"id"`             // uuid
	UserID        string    `json:"user_id"`        // uuid
	Address       string    `json:"address"`        // varchar
	IsSublet      bool      `json:"is_sublet"`      // bool
	Price         string    `json:"price"`          // varchar
	Rooms         string    `json:"rooms"`          // varchar
	RoomsOccupied int64     `json:"rooms_occupied"` // int8
	Bathrooms     string    `json:"bathrooms"`      // varchar
	Description   string    `json:"description"`    // text
	Gender        string    `json:"gender"`         // varchar
	PropertyType  string    `json:"property_type"`  // varchar
	Term          string    `json:"term"`           // varchar
	Title         string    `json:"title"`          // varchar
	Content       string    `json:"content"`        // text
	CreatedAt     time.Time `json:"created_at"`     // timestamptz
	UpdatedAt     time.Time `json:"updated_at"`     // timestamptz
	UpdatedBy     string    `json:"updated_by"`     // text
	CreatedBy     string    `json:"created_by"`     // uuid
}

// CreatePostRequest represents post creation data
type CreatePostRequest struct {
	UserID        string  `json:"userId"`
	Address       *string `json:"address"`
	IsSublet      *bool   `json:"isSublet"`
	Price         *string `json:"price"`
	Rooms         *string `json:"bedrooms"`
	RoomsOccupied *int64  `json:"roomsOccupied"`
	Bathrooms     *string `json:"bathrooms"`
	Description   *string `json:"description"`
	Gender        *string `json:"gender"`
	PropertyType  *string `json:"propertyType"`
	Term          *Term   `json:"terms"`
}

// UpdatePostRequest represents post update data
type UpdatePostRequest struct {
	Address       *string `json:"address"`
	IsSublet      *bool   `json:"is_sublet"`
	Price         *string `json:"price"`
	Rooms         *string `json:"rooms"`
	RoomsOccupied *int64  `json:"rooms_occupied"`
	Bathrooms     *string `json:"bathrooms"`
	Description   *string `json:"description"`
	Gender        *string `json:"gender"`
	PropertyType  *string `json:"property_type"`
	Term          *Term   `json:"term"`
}
