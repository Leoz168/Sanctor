package community

// Repository defines the interface for group data access
type Repository interface {
	Create(group *Community) error
	FindByID(id string) (*Community, error)
	FindAll() []*Community
	Update(group *Community) error
	Delete(id string) error
	AddUserToGroup(userGroup *UserCommunity) error
	AddGroupToInstitution(groupInstitution *CommunityInstitution) error
	RemoveUserFromGroup(userID, groupID string) error
	GetGroupMembers(groupID string) ([]*UserCommunity, error)
	GetUserGroups(userID string) []*UserCommunity
	IsUserInGroup(userID, groupID string) bool
	GetMemberCount(groupID string) int
	GetUserRole(userID, groupID string) (string, error)
}
