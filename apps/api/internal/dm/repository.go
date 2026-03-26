package dm

import "sync"

type InMemoryRepository struct {
	groups       map[string]*DMGroup
	groupUsers   map[string][]*DMGroupUser
	userGroups   map[string][]string
	groupMessage map[string][]*DMMessage
	mu           sync.RWMutex
}

func NewRepository() Repository {
	return &InMemoryRepository{
		groups:       make(map[string]*DMGroup),
		groupUsers:   make(map[string][]*DMGroupUser),
		userGroups:   make(map[string][]string),
		groupMessage: make(map[string][]*DMMessage),
	}
}

func (r *InMemoryRepository) CreateGroup(group *DMGroup) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.groups[group.ID] = group
	return nil
}

func (r *InMemoryRepository) AddUserToGroup(groupUser *DMGroupUser) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.groups[groupUser.GroupID]; !ok {
		return ErrDMGroupNotFound
	}

	for _, existing := range r.groupUsers[groupUser.GroupID] {
		if existing.UserID == groupUser.UserID {
			return nil
		}
	}

	r.groupUsers[groupUser.GroupID] = append(r.groupUsers[groupUser.GroupID], groupUser)
	r.userGroups[groupUser.UserID] = append(r.userGroups[groupUser.UserID], groupUser.GroupID)
	return nil
}

func (r *InMemoryRepository) GetGroupUsers(groupID string) ([]*DMGroupUser, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if _, ok := r.groups[groupID]; !ok {
		return nil, ErrDMGroupNotFound
	}

	users := make([]*DMGroupUser, len(r.groupUsers[groupID]))
	copy(users, r.groupUsers[groupID])
	return users, nil
}

func (r *InMemoryRepository) GetUserGroups(userID string) ([]*DMGroup, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	groupIDs := r.userGroups[userID]
	groups := make([]*DMGroup, 0, len(groupIDs))
	for _, groupID := range groupIDs {
		if group, ok := r.groups[groupID]; ok {
			groups = append(groups, group)
		}
	}
	return groups, nil
}

func (r *InMemoryRepository) FindDirectGroupByUsers(userA, userB string) (*DMGroup, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	groupIDs := r.userGroups[userA]
	for _, groupID := range groupIDs {
		users := r.groupUsers[groupID]
		if len(users) != 2 {
			continue
		}

		foundA := false
		foundB := false
		for _, member := range users {
			if member.UserID == userA {
				foundA = true
			}
			if member.UserID == userB {
				foundB = true
			}
		}

		if foundA && foundB {
			if group, ok := r.groups[groupID]; ok {
				return group, nil
			}
		}
	}

	return nil, ErrDMGroupNotFound
}

func (r *InMemoryRepository) IsUserInGroup(userID, groupID string) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, member := range r.groupUsers[groupID] {
		if member.UserID == userID {
			return true
		}
	}

	return false
}

func (r *InMemoryRepository) SaveMessage(message *DMMessage) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.groups[message.GroupID]; !ok {
		return ErrDMGroupNotFound
	}

	r.groupMessage[message.GroupID] = append(r.groupMessage[message.GroupID], message)
	return nil
}

func (r *InMemoryRepository) GetGroupMessages(groupID string, limit int) ([]*DMMessage, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if _, ok := r.groups[groupID]; !ok {
		return nil, ErrDMGroupNotFound
	}

	messages := r.groupMessage[groupID]
	if limit <= 0 || len(messages) <= limit {
		out := make([]*DMMessage, len(messages))
		copy(out, messages)
		return out, nil
	}

	start := len(messages) - limit
	out := make([]*DMMessage, len(messages[start:]))
	copy(out, messages[start:])
	return out, nil
}
