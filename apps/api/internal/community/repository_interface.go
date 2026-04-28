package community

import "github.com/google/uuid"

// Repository defines the interface for group data access
type Repository interface {
	Create(group *Community) error
	FindByID(id uuid.UUID) (*Community, error)
	FindAll() []*Community
	Update(group *Community) error
	Delete(id uuid.UUID) error
	AddUserToGroup(userGroup *UserCommunity) error
	AddGroupToInstitution(groupInstitution *CommunityInstitution) error
	RemoveUserFromGroup(userID, groupID uuid.UUID) error
	GetGroupMembers(groupID uuid.UUID) ([]*UserCommunity, error)
	GetUserGroups(userID uuid.UUID) []*UserCommunity
	IsUserInGroup(userID, groupID uuid.UUID) bool
	GetMemberCount(groupID uuid.UUID) int
	GetUserRole(userID, groupID uuid.UUID) (string, error)
}
