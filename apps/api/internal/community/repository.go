package community

import (
	"errors"
	"sync"
)

// InMemoryRepository handles data access for groups in memory
type InMemoryRepository struct {
	groups            map[string]*Community              // groupID -> Group
	userGroups        map[string][]*UserCommunity        // userID -> []UserGroup
	groupUsers        map[string][]*UserCommunity        // groupID -> []UserGroup
	groupInstitutions map[string][]*CommunityInstitution // groupID -> []GroupInstitution
	mu                sync.RWMutex
}

// NewRepository creates a new in-memory group repository
func NewRepository() Repository {
	return &InMemoryRepository{
		groups:            make(map[string]*Community),
		userGroups:        make(map[string][]*UserCommunity),
		groupUsers:        make(map[string][]*UserCommunity),
		groupInstitutions: make(map[string][]*CommunityInstitution),
	}
}

// Create creates a new group
func (r *InMemoryRepository) Create(group *Community) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.groups[group.ID.String()] = group
	return nil
}

// FindByID finds a group by ID
func (r *InMemoryRepository) FindByID(id string) (*Community, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	group, exists := r.groups[id]
	if !exists {
		return nil, errors.New("group not found")
	}
	return group, nil
}

// FindAll returns all groups
func (r *InMemoryRepository) FindAll() []*Community {
	r.mu.RLock()
	defer r.mu.RUnlock()

	groups := make([]*Community, 0, len(r.groups))
	for _, group := range r.groups {
		groups = append(groups, group)
	}
	return groups
}

// Update updates an existing group
func (r *InMemoryRepository) Update(group *Community) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.groups[group.ID.String()]; !exists {
		return errors.New("group not found")
	}
	r.groups[group.ID.String()] = group
	return nil
}

// Delete deletes a group and its memberships
func (r *InMemoryRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.groups[id]; !exists {
		return errors.New("group not found")
	}

	// Remove group
	delete(r.groups, id)

	// Remove all memberships for this group
	delete(r.groupUsers, id)
	delete(r.groupInstitutions, id)

	// Remove from user's group lists
	for userID, userGroupsList := range r.userGroups {
		newList := make([]*UserCommunity, 0)
		for _, ug := range userGroupsList {
			if ug.CommunityID.String() != id {
				newList = append(newList, ug)
			}
		}
		r.userGroups[userID] = newList
	}

	return nil
}

// AddUserToGroup adds a user to a group
func (r *InMemoryRepository) AddUserToGroup(userGroup *UserCommunity) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Check if group exists
	if _, exists := r.groups[userGroup.CommunityID.String()]; !exists {
		return errors.New("group not found")
	}

	// Check if user is already in group
	if r.isUserInGroup(userGroup.UserID.String(), userGroup.CommunityID.String()) {
		return errors.New("user already in group")
	}

	// Add to groupUsers
	r.groupUsers[userGroup.CommunityID.String()] = append(r.groupUsers[userGroup.CommunityID.String()], userGroup)

	// Add to userGroups
	r.userGroups[userGroup.UserID.String()] = append(r.userGroups[userGroup.UserID.String()], userGroup)

	return nil
}

// AddGroupToInstitution links a group to an institution
func (r *InMemoryRepository) AddGroupToInstitution(groupInstitution *CommunityInstitution) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.groups[groupInstitution.CommunityID.String()]; !exists {
		return errors.New("group not found")
	}

	for _, gi := range r.groupInstitutions[groupInstitution.CommunityID.String()] {
		if gi.InstitutionID == groupInstitution.InstitutionID {
			return errors.New("group already linked to institution")
		}
	}

	r.groupInstitutions[groupInstitution.CommunityID.String()] = append(r.groupInstitutions[groupInstitution.CommunityID.String()], groupInstitution)
	return nil
}

// RemoveUserFromGroup removes a user from a group
func (r *InMemoryRepository) RemoveUserFromGroup(userID, groupID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if !r.isUserInGroup(userID, groupID) {
		return errors.New("user not in group")
	}

	// Remove from groupUsers
	newGroupUsers := make([]*UserCommunity, 0)
	for _, ug := range r.groupUsers[groupID] {
		if ug.UserID.String() != userID {
			newGroupUsers = append(newGroupUsers, ug)
		}
	}
	r.groupUsers[groupID] = newGroupUsers

	// Remove from userGroups
	newUserGroups := make([]*UserCommunity, 0)
	for _, ug := range r.userGroups[userID] {
		if ug.CommunityID.String() != groupID {
			newUserGroups = append(newUserGroups, ug)
		}
	}
	r.userGroups[userID] = newUserGroups

	return nil
}

// GetGroupMembers returns all users in a group
func (r *InMemoryRepository) GetGroupMembers(groupID string) ([]*UserCommunity, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if _, exists := r.groups[groupID]; !exists {
		return nil, errors.New("group not found")
	}

	return r.groupUsers[groupID], nil
}

// GetUserGroups returns all groups a user belongs to
func (r *InMemoryRepository) GetUserGroups(userID string) []*UserCommunity {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.userGroups[userID]
}

// IsUserInGroup checks if a user is in a group (exported version)
func (r *InMemoryRepository) IsUserInGroup(userID, groupID string) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.isUserInGroup(userID, groupID)
}

// isUserInGroup checks if a user is in a group (internal, no lock)
func (r *InMemoryRepository) isUserInGroup(userID, groupID string) bool {
	for _, ug := range r.userGroups[userID] {
		if ug.CommunityID.String() == groupID {
			return true
		}
	}
	return false
}

// GetMemberCount returns the number of members in a group
func (r *InMemoryRepository) GetMemberCount(groupID string) int {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return len(r.groupUsers[groupID])
}

// GetUserRole returns the role of a user in a group
func (r *InMemoryRepository) GetUserRole(userID, groupID string) (string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, ug := range r.userGroups[userID] {
		if ug.CommunityID.String() == groupID {
			return ug.Role, nil
		}
	}
	return "", errors.New("user not in group")
}
