package community

import (
	"time"

	"github.com/google/uuid"
)

// Community represents a community in the system
type Community struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name        string    `json:"name" gorm:"type:varchar(200);not null"`
	Description string    `json:"description,omitempty" gorm:"type:text"`
	IsPrivate   bool      `json:"isPrivate" gorm:"default:false"`
	CreatedBy   uuid.UUID `json:"createdBy" gorm:"type:uuid;not null;index"` // User ID of creator
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

func (Community) TableName() string {
	return "communities"
}

// UserCommunity represents the many-to-many relationship between users and communities
type UserCommunity struct {
	UserID      uuid.UUID `json:"userId" gorm:"type:uuid;primaryKey"`
	CommunityID uuid.UUID `json:"communityId" gorm:"column:community_id;type:uuid;primaryKey"`
	Role        string    `json:"role" gorm:"type:varchar(20);default:'member'"` // "member", "admin", "owner"
	JoinedAt    time.Time `json:"joinedAt" gorm:"autoCreateTime"`
}

func (UserCommunity) TableName() string {
	return "user_communities"
}

// CommunityInstitution represents the many-to-many relationship between communities and institutions
type CommunityInstitution struct {
	CommunityID   uuid.UUID `json:"communityId" gorm:"column:community_id;type:uuid;primaryKey"`
	InstitutionID uuid.UUID `json:"institutionId" gorm:"type:uuid;primaryKey"`
	LinkedAt      time.Time `json:"linkedAt" gorm:"autoCreateTime"`
}

func (CommunityInstitution) TableName() string {
	return "community_institutions"
}

// CreateCommunityRequest represents the data needed to create a new community
type CreateCommunityRequest struct {
	Name          string `json:"name"`
	Description   string `json:"description,omitempty"`
	IsPrivate     bool   `json:"isPrivate"`
	InstitutionID string `json:"institutionId"` // institution association
	CreatedBy     string `json:"createdBy"`     // User ID
}

// UpdateCommunityRequest represents the data that can be updated
type UpdateCommunityRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	IsPrivate   *bool  `json:"isPrivate,omitempty"`
}

// AddUserToCommunityRequest represents adding a user to a community
type AddUserToCommunityRequest struct {
	UserID      string `json:"userId"`
	CommunityID string `json:"communityId"`
	Role        string `json:"role,omitempty"` // defaults to "member"
}

// CommunityWithMembers includes community data and member count
type CommunityWithMembers struct {
	*Community
	MemberCount int `json:"memberCount"`
}

// UserCommunityInfo includes user info in a community context
type UserCommunityInfo struct {
	UserID   uuid.UUID `json:"userId"`
	Role     string    `json:"role"`
	JoinedAt time.Time `json:"joinedAt"`
}
