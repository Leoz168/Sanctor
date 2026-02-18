package post

import "time"

// Term represents the lease term season
type Term string

const (
	TermWinter Term = "Winter"
	TermSpring Term = "Spring"
	TermSummer Term = "Summer"
	TermFall   Term = "Fall"
)

// Model represents a post in the system
type Model struct {
	ID            string    `json:"id"`
	UserID        string    `json:"userId"`
	Address       string    `json:"address"`
	IsSublet      bool      `json:"isSublet"`
	Price         string    `json:"price"`
	Rooms         string    `json:"bedrooms"`
	RoomsOccupied int       `json:"roomsOccupied"`
	Bathrooms     string    `json:"bathrooms"`
	Description   string    `json:"description"`
	Gender        string    `json:"gender"`
	PropertyType  string    `json:"propertyType"`
	Term          Term      `json:"terms"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
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
