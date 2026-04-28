package post

import (
	"time"

	sharedtypes "sanctor/pkg/types"

	"github.com/google/uuid"
)

// Model represents a post in the system
type Post struct {
	ID              uuid.UUID          `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`    // uuid
	UserID          uuid.UUID          `json:"user_id" gorm:"type:uuid;index:idx_posts_user_id"`            // uuid
	Address         string             `json:"address" gorm:"index:idx_posts_address"`                      // varchar
	IsSublet        bool               `json:"is_sublet"`                                                   // bool
	Price           int64              `json:"price" gorm:"index:idx_posts_price"`                          // int8
	Rooms           int64              `json:"rooms"`                                                       // int8
	RoomsOccupied   int64              `json:"rooms_occupied"`                                              // int8
	Bathrooms       int64              `json:"bathrooms"`                                                   // int8
	Description     string             `json:"description"`                                                 // text
	Gender          sharedtypes.Gender `json:"gender" gorm:"index:idx_posts_gender"`                        // varchar
	PropertyType    string             `json:"property_type"`                                               // varchar
	Term            sharedtypes.Term   `json:"term"`                                                        // varchar
	Title           string             `json:"title"`                                                       // varchar
	Content         string             `json:"content"`                                                     // text
	CreatedAt       time.Time          `json:"created_at"`                                                  // timestamptz
	UpdatedAt       time.Time          `json:"updated_at"`                                                  // timestamptz
	UpdatedByUserID uuid.UUID          `json:"updated_by_user_id" gorm:"column:updated_by;type:uuid;index"` // uuid
	CreatedByUserID uuid.UUID          `json:"created_by_user_id" gorm:"column:created_by;type:uuid;index"` // uuid
}

// PostCommunity represents the many-to-many relationship between posts and groups
type PostCommunity struct {
	PostID      uuid.UUID `json:"postId" gorm:"type:uuid;primaryKey"`
	CommunityID uuid.UUID `json:"communityId" gorm:"column:group_id;type:uuid;primaryKey"`
	LinkedAt    time.Time `json:"linkedAt" gorm:"autoCreateTime"`
}

func (PostCommunity) TableName() string {
	return "post_groups"
}

// PostInstitution represents the many-to-many relationship between posts and institutions
type PostInstitution struct {
	PostID        uuid.UUID `json:"postId" gorm:"type:uuid;primaryKey"`
	InstitutionID uuid.UUID `json:"institutionId" gorm:"type:uuid;primaryKey"`
	LinkedAt      time.Time `json:"linkedAt" gorm:"autoCreateTime"`
}

// CreatePostRequest represents post creation data
type CreatePostRequest struct {
	UserID         string              `json:"userId"`
	Address        *string             `json:"address"`
	IsSublet       *bool               `json:"isSublet"`
	Price          *int64              `json:"price"`
	Rooms          *int64              `json:"bedrooms"`
	RoomsOccupied  *int64              `json:"roomsOccupied"`
	Bathrooms      *int64              `json:"bathrooms"`
	Description    *string             `json:"description"`
	Gender         *sharedtypes.Gender `json:"gender"`
	PropertyType   *string             `json:"propertyType"`
	Term           *sharedtypes.Term   `json:"terms"`
	CommunityIDs   []string            `json:"communityIds,omitempty"`
	InstitutionIDs []string            `json:"institutionIds,omitempty"`
}

// UpdatePostRequest represents post update data
type UpdatePostRequest struct {
	Address       *string             `json:"address"`
	IsSublet      *bool               `json:"is_sublet"`
	Price         *int64              `json:"price"`
	Rooms         *int64              `json:"rooms"`
	RoomsOccupied *int64              `json:"rooms_occupied"`
	Bathrooms     *int64              `json:"bathrooms"`
	Description   *string             `json:"description"`
	Gender        *sharedtypes.Gender `json:"gender"`
	PropertyType  *string             `json:"property_type"`
	Term          *sharedtypes.Term   `json:"term"`
}

// PostSearchFilters represents backend filter/sort/pagination options for post search.
type PostSearchFilters struct {
	Query         string
	MinPrice      *int64
	MaxPrice      *int64
	MinRooms      *int64
	MinBathrooms  *int64
	IsSublet      *bool
	Gender        *sharedtypes.Gender
	PropertyType  string
	Term          *sharedtypes.Term
	CommunityID   string
	InstitutionID string
	SortBy        string
	SortOrder     string
	Limit         int
	Offset        int
}
