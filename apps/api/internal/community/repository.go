package community

import (
	"errors"
	"sync"

	"github.com/google/uuid"
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
func (r *InMemoryRepository) FindByID(id uuid.UUID) (*Community, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	group, exists := r.groups[id.String()]
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
func (r *InMemoryRepository) Delete(id uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	groupKey := id.String()
	if _, exists := r.groups[groupKey]; !exists {
		return errors.New("group not found")
	}

	// Remove group
	delete(r.groups, groupKey)

	// Remove all memberships for this group
	delete(r.groupUsers, groupKey)
	delete(r.groupInstitutions, groupKey)

	// Remove from user's group lists
	for userID, userGroupsList := range r.userGroups {
		newList := make([]*UserCommunity, 0)
		for _, ug := range userGroupsList {
			if ug.CommunityID != id {
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
	if r.isUserInGroup(userGroup.UserID, userGroup.CommunityID) {
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
func (r *InMemoryRepository) RemoveUserFromGroup(userID, groupID uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if !r.isUserInGroup(userID, groupID) {
		return errors.New("user not in group")
	}

	// Remove from groupUsers
	newGroupUsers := make([]*UserCommunity, 0)
	groupKey := groupID.String()
	for _, ug := range r.groupUsers[groupKey] {
		if ug.UserID != userID {
			newGroupUsers = append(newGroupUsers, ug)
		}
	}
	r.groupUsers[groupKey] = newGroupUsers

	// Remove from userGroups
	newUserGroups := make([]*UserCommunity, 0)
	userKey := userID.String()
	for _, ug := range r.userGroups[userKey] {
		if ug.CommunityID != groupID {
			newUserGroups = append(newUserGroups, ug)
		}
	}
	r.userGroups[userKey] = newUserGroups

	return nil
}

// GetGroupMembers returns all users in a group
func (r *InMemoryRepository) GetGroupMembers(groupID uuid.UUID) ([]*UserCommunity, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	groupKey := groupID.String()
	if _, exists := r.groups[groupKey]; !exists {
		return nil, errors.New("group not found")
	}

	return r.groupUsers[groupKey], nil
}

// GetUserGroups returns all groups a user belongs to
func (r *InMemoryRepository) GetUserGroups(userID uuid.UUID) []*UserCommunity {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.userGroups[userID.String()]
}

// IsUserInGroup checks if a user is in a group (exported version)
func (r *InMemoryRepository) IsUserInGroup(userID, groupID uuid.UUID) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.isUserInGroup(userID, groupID)
}

// isUserInGroup checks if a user is in a group (internal, no lock)
func (r *InMemoryRepository) isUserInGroup(userID, groupID uuid.UUID) bool {
	for _, ug := range r.userGroups[userID.String()] {
		if ug.CommunityID == groupID {
			return true
		}
	}
	return false
}

// GetMemberCount returns the number of members in a group
func (r *InMemoryRepository) GetMemberCount(groupID uuid.UUID) int {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return len(r.groupUsers[groupID.String()])
}

// GetUserRole returns the role of a user in a group
func (r *InMemoryRepository) GetUserRole(userID, groupID uuid.UUID) (string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, ug := range r.userGroups[userID.String()] {
		if ug.CommunityID == groupID {
			return ug.Role, nil
		}
	}
	return "", errors.New("user not in group")
}
