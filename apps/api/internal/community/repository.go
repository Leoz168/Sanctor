package community

import (
	"errors"
	"sync"

	"github.com/google/uuid"
)

// InMemoryRepository handles data access for communities in memory
type InMemoryRepository struct {
	communities          map[string]*Community              // communityID -> Community
	userCommunities      map[string][]*UserCommunity        // userID -> []UserCommunity
	communityUsers       map[string][]*UserCommunity        // communityID -> []UserCommunity
	communityInstitutions map[string][]*CommunityInstitution // communityID -> []CommunityInstitution
	mu                sync.RWMutex
}

// NewRepository creates a new in-memory community repository
func NewRepository() Repository {
	return &InMemoryRepository{
		communities:          make(map[string]*Community),
		userCommunities:      make(map[string][]*UserCommunity),
		communityUsers:       make(map[string][]*UserCommunity),
		communityInstitutions: make(map[string][]*CommunityInstitution),
	}
}

// Create creates a new community
func (r *InMemoryRepository) Create(community *Community) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.communities[community.ID.String()] = community
	return nil
}

// FindByID finds a community by ID
func (r *InMemoryRepository) FindByID(id uuid.UUID) (*Community, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	community, exists := r.communities[id.String()]
	if !exists {
		return nil, errors.New("community not found")
	}
	return community, nil
}

// FindAll returns all communities
func (r *InMemoryRepository) FindAll() []*Community {
	r.mu.RLock()
	defer r.mu.RUnlock()

	communities := make([]*Community, 0, len(r.communities))
	for _, community := range r.communities {
		communities = append(communities, community)
	}
	return communities
}

// Update updates an existing community
func (r *InMemoryRepository) Update(community *Community) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.communities[community.ID.String()]; !exists {
		return errors.New("community not found")
	}
	r.communities[community.ID.String()] = community
	return nil
}

// Delete deletes a community and its memberships
func (r *InMemoryRepository) Delete(id uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	communityKey := id.String()
	if _, exists := r.communities[communityKey]; !exists {
		return errors.New("community not found")
	}

	// Remove community
	delete(r.communities, communityKey)

	// Remove all memberships for this community
	delete(r.communityUsers, communityKey)
	delete(r.communityInstitutions, communityKey)

	// Remove from user's community lists
	for userID, userCommunities := range r.userCommunities {
		newList := make([]*UserCommunity, 0)
		for _, uc := range userCommunities {
			if uc.CommunityID != id {
				newList = append(newList, uc)
			}
		}
		r.userCommunities[userID] = newList
	}

	return nil
}

// AddUserToCommunity adds a user to a community
func (r *InMemoryRepository) AddUserToCommunity(userCommunity *UserCommunity) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Check if community exists
	if _, exists := r.communities[userCommunity.CommunityID.String()]; !exists {
		return errors.New("community not found")
	}

	// Check if user is already in community
	if r.isUserInCommunity(userCommunity.UserID, userCommunity.CommunityID) {
		return errors.New("user already in community")
	}

	// Add to communityUsers
	r.communityUsers[userCommunity.CommunityID.String()] = append(r.communityUsers[userCommunity.CommunityID.String()], userCommunity)

	// Add to userCommunities
	r.userCommunities[userCommunity.UserID.String()] = append(r.userCommunities[userCommunity.UserID.String()], userCommunity)

	return nil
}

// AddCommunityToInstitution links a community to an institution
func (r *InMemoryRepository) AddCommunityToInstitution(communityInstitution *CommunityInstitution) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.communities[communityInstitution.CommunityID.String()]; !exists {
		return errors.New("community not found")
	}

	for _, ci := range r.communityInstitutions[communityInstitution.CommunityID.String()] {
		if ci.InstitutionID == communityInstitution.InstitutionID {
			return errors.New("community already linked to institution")
		}
	}

	r.communityInstitutions[communityInstitution.CommunityID.String()] = append(r.communityInstitutions[communityInstitution.CommunityID.String()], communityInstitution)
	return nil
}

// RemoveUserFromCommunity removes a user from a community
func (r *InMemoryRepository) RemoveUserFromCommunity(userID, communityID uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if !r.isUserInCommunity(userID, communityID) {
		return errors.New("user not in community")
	}

	// Remove from communityUsers
	newCommunityUsers := make([]*UserCommunity, 0)
	communityKey := communityID.String()
	for _, uc := range r.communityUsers[communityKey] {
		if uc.UserID != userID {
			newCommunityUsers = append(newCommunityUsers, uc)
		}
	}
	r.communityUsers[communityKey] = newCommunityUsers

	// Remove from userCommunities
	newUserCommunities := make([]*UserCommunity, 0)
	userKey := userID.String()
	for _, uc := range r.userCommunities[userKey] {
		if uc.CommunityID != communityID {
			newUserCommunities = append(newUserCommunities, uc)
		}
	}
	r.userCommunities[userKey] = newUserCommunities

	return nil
}

// GetCommunityMembers returns all users in a community
func (r *InMemoryRepository) GetCommunityMembers(communityID uuid.UUID) ([]*UserCommunity, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	communityKey := communityID.String()
	if _, exists := r.communities[communityKey]; !exists {
		return nil, errors.New("community not found")
	}

	return r.communityUsers[communityKey], nil
}

// GetUserCommunities returns all communities a user belongs to
func (r *InMemoryRepository) GetUserCommunities(userID uuid.UUID) []*UserCommunity {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.userCommunities[userID.String()]
}

// IsUserInCommunity checks if a user is in a community (exported version)
func (r *InMemoryRepository) IsUserInCommunity(userID, communityID uuid.UUID) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.isUserInCommunity(userID, communityID)
}

// isUserInCommunity checks if a user is in a community (internal, no lock)
func (r *InMemoryRepository) isUserInCommunity(userID, communityID uuid.UUID) bool {
	for _, uc := range r.userCommunities[userID.String()] {
		if uc.CommunityID == communityID {
			return true
		}
	}
	return false
}

// GetMemberCount returns the number of members in a community
func (r *InMemoryRepository) GetMemberCount(communityID uuid.UUID) int {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return len(r.communityUsers[communityID.String()])
}

// GetUserRole returns the role of a user in a community
func (r *InMemoryRepository) GetUserRole(userID, communityID uuid.UUID) (string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, uc := range r.userCommunities[userID.String()] {
		if uc.CommunityID == communityID {
			return uc.Role, nil
		}
	}
	return "", errors.New("user not in community")
}
