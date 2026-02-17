package group

import "time"

// Group represents a group in the system
type Group struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	IsPrivate   bool      `json:"isPrivate"`
	CreatedBy   string    `json:"createdBy"` // User ID of creator
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// UserGroup represents the many-to-many relationship between users and groups
type UserGroup struct {
	UserID    string    `json:"userId"`
	GroupID   string    `json:"groupId"`
	Role      string    `json:"role"` // "member", "admin", "owner"
	JoinedAt  time.Time `json:"joinedAt"`
}

// CreateGroupRequest represents the data needed to create a new group
type CreateGroupRequest struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	IsPrivate   bool   `json:"isPrivate"`
	CreatedBy   string `json:"createdBy"` // User ID
}

// UpdateGroupRequest represents the data that can be updated
type UpdateGroupRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	IsPrivate   *bool  `json:"isPrivate,omitempty"`
}

// AddUserToGroupRequest represents adding a user to a group
type AddUserToGroupRequest struct {
	UserID  string `json:"userId"`
	GroupID string `json:"groupId"`
	Role    string `json:"role,omitempty"` // defaults to "member"
}

// GroupWithMembers includes group data and member count
type GroupWithMembers struct {
	*Group
	MemberCount int `json:"memberCount"`
}

// UserGroupInfo includes user info in a group context
type UserGroupInfo struct {
	UserID   string    `json:"userId"`
	Role     string    `json:"role"`
	JoinedAt time.Time `json:"joinedAt"`
}
