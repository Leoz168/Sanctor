package group

// Repository defines the interface for group data access
type Repository interface {
	Create(group *Group) error
	FindByID(id string) (*Group, error)
	FindAll() []*Group
	Update(group *Group) error
	Delete(id string) error
	AddUserToGroup(userGroup *UserGroup) error
	RemoveUserFromGroup(userID, groupID string) error
	GetGroupMembers(groupID string) ([]*UserGroup, error)
	GetUserGroups(userID string) []*UserGroup
	IsUserInGroup(userID, groupID string) bool
	GetMemberCount(groupID string) int
	GetUserRole(userID, groupID string) (string, error)
}
