package community

import "github.com/google/uuid"

// Repository defines the interface for community data access
type Repository interface {
	Create(community *Community) error
	FindByID(id uuid.UUID) (*Community, error)
	FindAll() []*Community
	Update(community *Community) error
	Delete(id uuid.UUID) error
	AddUserToCommunity(userCommunity *UserCommunity) error
	AddCommunityToInstitution(communityInstitution *CommunityInstitution) error
	RemoveUserFromCommunity(userID, communityID uuid.UUID) error
	GetCommunityMembers(communityID uuid.UUID) ([]*UserCommunity, error)
	GetUserCommunities(userID uuid.UUID) []*UserCommunity
	IsUserInCommunity(userID, communityID uuid.UUID) bool
	GetMemberCount(communityID uuid.UUID) int
	GetUserRole(userID, communityID uuid.UUID) (string, error)
}
