package dm

import (
	"errors"
	"sync"

	"github.com/google/uuid"
)

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
	r.groups[group.ID.String()] = group
	return nil
}

func (r *InMemoryRepository) AddUserToGroup(groupUser *DMGroupUser) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	groupID := groupUser.GroupID.String()
	userID := groupUser.UserID.String()

	if _, ok := r.groups[groupID]; !ok {
		return ErrDMGroupNotFound
	}

	for _, existing := range r.groupUsers[groupID] {
		if existing.UserID == groupUser.UserID {
			return nil
		}
	}

	r.groupUsers[groupID] = append(r.groupUsers[groupID], groupUser)
	r.userGroups[userID] = append(r.userGroups[userID], groupID)
	return nil
}

func (r *InMemoryRepository) GetGroupUsers(groupID uuid.UUID) ([]*DMGroupUser, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	groupKey := groupID.String()
	if _, ok := r.groups[groupKey]; !ok {
		return nil, ErrDMGroupNotFound
	}

	users := make([]*DMGroupUser, len(r.groupUsers[groupKey]))
	copy(users, r.groupUsers[groupKey])
	return users, nil
}

func (r *InMemoryRepository) GetUserGroups(userID uuid.UUID) ([]*DMGroup, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	groupIDs := r.userGroups[userID.String()]
	groups := make([]*DMGroup, 0, len(groupIDs))
	for _, groupID := range groupIDs {
		if group, ok := r.groups[groupID]; ok {
			groups = append(groups, group)
		}
	}
	return groups, nil
}

func (r *InMemoryRepository) FindDirectGroupByUsers(userA, userB uuid.UUID) (*DMGroup, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	groupIDs := r.userGroups[userA.String()]
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

func (r *InMemoryRepository) IsUserInGroup(userID, groupID uuid.UUID) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, member := range r.groupUsers[groupID.String()] {
		if member.UserID == userID {
			return true
		}
	}

	return false
}

func (r *InMemoryRepository) SaveMessage(message *DMMessage) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	groupID := message.GroupID.String()

	if _, ok := r.groups[groupID]; !ok {
		return ErrDMGroupNotFound
	}

	r.groupMessage[groupID] = append(r.groupMessage[groupID], message)
	return nil
}

func (r *InMemoryRepository) GetGroupMessages(groupID uuid.UUID, limit int) ([]*DMMessage, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	groupKey := groupID.String()
	if _, ok := r.groups[groupKey]; !ok {
		return nil, ErrDMGroupNotFound
	}

	messages := r.groupMessage[groupKey]
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

func (r *InMemoryRepository) GetLatestGroupMessage(groupID uuid.UUID) (*DMMessage, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	groupKey := groupID.String()
	if _, ok := r.groups[groupKey]; !ok {
		return nil, ErrDMGroupNotFound
	}

	messages := r.groupMessage[groupKey]
	if len(messages) == 0 {
		return nil, nil
	}

	return messages[len(messages)-1], nil
}

func (r *InMemoryRepository) FindMessageByID(messageID uuid.UUID) (*DMMessage, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, messages := range r.groupMessage {
		for _, message := range messages {
			if message.ID == messageID {
				return message, nil
			}
		}
	}

	return nil, errors.New("message not found")
}

func (r *InMemoryRepository) UpdateMessage(message *DMMessage) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	groupKey := message.GroupID.String()
	messages := r.groupMessage[groupKey]
	for index, existing := range messages {
		if existing.ID == message.ID {
			r.groupMessage[groupKey][index] = message
			return nil
		}
	}

	return errors.New("message not found")
}

func (r *InMemoryRepository) DeleteMessage(messageID uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for groupKey, messages := range r.groupMessage {
		for index, message := range messages {
			if message.ID == messageID {
				r.groupMessage[groupKey] = append(messages[:index], messages[index+1:]...)
				return nil
			}
		}
	}

	return errors.New("message not found")
}
