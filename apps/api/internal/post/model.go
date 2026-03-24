package post

import (
	sharedtypes "sanctor/pkg/types"
	"time"
)

// Model represents a post in the system
type Post struct {
	ID              string    `json:"id"`             // uuid
	UserID          string    `json:"user_id"`        // uuid
	Address         string    `json:"address" gorm:"index:idx_posts_address"`        // varchar
	IsSublet        bool      `json:"is_sublet"`      // bool
	Price           int64     `json:"price" gorm:"index:idx_posts_price"`          // int8
	Rooms           int64     `json:"rooms"`          // int8
	RoomsOccupied   int64     `json:"rooms_occupied"` // int8
	Bathrooms       int64     `json:"bathrooms"`      // int8
	Description     string    `json:"description"`    // text
	Gender          sharedtypes.Gender `json:"gender" gorm:"index:idx_posts_gender"`         // varchar
	PropertyType    string    `json:"property_type"`  // varchar
	Term            sharedtypes.Term `json:"term"`           // varchar
	Title           string    `json:"title"`          // varchar
	Content         string    `json:"content"`        // text
	CreatedAt       time.Time `json:"created_at"`     // timestamptz
	UpdatedAt       time.Time `json:"updated_at"`     // timestamptz
	UpdatedByUserID string    `json:"updated_by_user_id" gorm:"column:updated_by"` // uuid
	CreatedByUserID string    `json:"created_by_user_id" gorm:"column:created_by"` // uuid
}

// PostGroup represents the many-to-many relationship between posts and groups
type PostGroup struct {
	PostID   string    `json:"postId" gorm:"type:uuid;primaryKey"`
	GroupID  string    `json:"groupId" gorm:"type:uuid;primaryKey"`
	LinkedAt time.Time `json:"linkedAt" gorm:"autoCreateTime"`
}

// PostInstitution represents the many-to-many relationship between posts and institutions
type PostInstitution struct {
	PostID        string    `json:"postId" gorm:"type:uuid;primaryKey"`
	InstitutionID string    `json:"institutionId" gorm:"type:uuid;primaryKey"`
	LinkedAt      time.Time `json:"linkedAt" gorm:"autoCreateTime"`
}

// CreatePostRequest represents post creation data
type CreatePostRequest struct {
	UserID         string   `json:"userId"`
	Address        *string  `json:"address"`
	IsSublet       *bool    `json:"isSublet"`
	Price          *int64   `json:"price"`
	Rooms          *int64   `json:"bedrooms"`
	RoomsOccupied  *int64   `json:"roomsOccupied"`
	Bathrooms      *int64   `json:"bathrooms"`
	Description    *string  `json:"description"`
	Gender         *sharedtypes.Gender  `json:"gender"`
	PropertyType   *string  `json:"propertyType"`
	Term           *sharedtypes.Term    `json:"terms"`
	GroupIDs       []string `json:"groupIds,omitempty"`
	InstitutionIDs []string `json:"institutionIds,omitempty"`
}

// UpdatePostRequest represents post update data
type UpdatePostRequest struct {
	Address       *string `json:"address"`
	IsSublet      *bool   `json:"is_sublet"`
	Price         *int64  `json:"price"`
	Rooms         *int64  `json:"rooms"`
	RoomsOccupied *int64  `json:"rooms_occupied"`
	Bathrooms     *int64  `json:"bathrooms"`
	Description   *string `json:"description"`
	Gender        *sharedtypes.Gender `json:"gender"`
	PropertyType  *string `json:"property_type"`
	Term          *sharedtypes.Term   `json:"term"`
}
