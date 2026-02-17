package group

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// Service handles business logic for group operations
type Service struct {
	repo *Repository
}

// NewService creates a new group service
func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

// CreateGroup creates a new group with validation
func (s *Service) CreateGroup(req CreateGroupRequest) (*Group, error) {
	// Validate input
	if req.Name == "" {
		return nil, errors.New("group name is required")
	}

	if req.CreatedBy == "" {
		return nil, errors.New("creator user ID is required")
	}

	// Create group
	group := &Group{
		ID:          uuid.New().String(),
		Name:        req.Name,
		Description: req.Description,
		IsPrivate:   req.IsPrivate,
		CreatedBy:   req.CreatedBy,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := s.repo.Create(group); err != nil {
		return nil, err
	}

	// Automatically add creator as owner
	userGroup := &UserGroup{
		UserID:   req.CreatedBy,
		GroupID:  group.ID,
		Role:     "owner",
		JoinedAt: time.Now(),
	}

	if err := s.repo.AddUserToGroup(userGroup); err != nil {
		// Rollback: delete the group if adding creator fails
		s.repo.Delete(group.ID)
		return nil, errors.New("failed to add creator to group")
	}

	return group, nil
}

// GetGroup retrieves a group by ID
func (s *Service) GetGroup(id string) (*Group, error) {
	if id == "" {
		return nil, errors.New("group ID is required")
	}

	return s.repo.FindByID(id)
}

// GetGroupWithMembers retrieves a group with member count
func (s *Service) GetGroupWithMembers(id string) (*GroupWithMembers, error) {
	group, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	memberCount := s.repo.GetMemberCount(id)

	return &GroupWithMembers{
		Group:       group,
		MemberCount: memberCount,
	}, nil
}

// GetAllGroups retrieves all groups
func (s *Service) GetAllGroups() ([]*Group, error) {
	return s.repo.FindAll(), nil
}

// UpdateGroup updates an existing group
func (s *Service) UpdateGroup(id string, req UpdateGroupRequest) (*Group, error) {
	group, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("group not found")
	}

	// Update fields if provided
	if req.Name != "" {
		group.Name = req.Name
	}
	if req.Description != "" {
		group.Description = req.Description
	}
	if req.IsPrivate != nil {
		group.IsPrivate = *req.IsPrivate
	}
	group.UpdatedAt = time.Now()

	if err := s.repo.Update(group); err != nil {
		return nil, err
	}

	return group, nil
}

// DeleteGroup deletes a group by ID
func (s *Service) DeleteGroup(id string) error {
	if id == "" {
		return errors.New("group ID is required")
	}

	return s.repo.Delete(id)
}

// AddUserToGroup adds a user to a group
func (s *Service) AddUserToGroup(req AddUserToGroupRequest) error {
	if req.UserID == "" || req.GroupID == "" {
		return errors.New("user ID and group ID are required")
	}

	// Validate group exists
	if _, err := s.repo.FindByID(req.GroupID); err != nil {
		return errors.New("group not found")
	}

	// Default role to "member"
	role := req.Role
	if role == "" {
		role = "member"
	}

	// Validate role
	if role != "member" && role != "admin" && role != "owner" {
		return errors.New("invalid role: must be member, admin, or owner")
	}

	userGroup := &UserGroup{
		UserID:   req.UserID,
		GroupID:  req.GroupID,
		Role:     role,
		JoinedAt: time.Now(),
	}

	return s.repo.AddUserToGroup(userGroup)
}

// RemoveUserFromGroup removes a user from a group
func (s *Service) RemoveUserFromGroup(userID, groupID string) error {
	if userID == "" || groupID == "" {
		return errors.New("user ID and group ID are required")
	}

	// Check if user is the owner
	role, err := s.repo.GetUserRole(userID, groupID)
	if err != nil {
		return err
	}

	if role == "owner" {
		// Check if there are other members
		members, _ := s.repo.GetGroupMembers(groupID)
		if len(members) > 1 {
			return errors.New("owner cannot leave group with other members. Transfer ownership first or delete the group")
		}
	}

	return s.repo.RemoveUserFromGroup(userID, groupID)
}

// GetGroupMembers returns all members of a group
func (s *Service) GetGroupMembers(groupID string) ([]*UserGroupInfo, error) {
	if groupID == "" {
		return nil, errors.New("group ID is required")
	}

	userGroups, err := s.repo.GetGroupMembers(groupID)
	if err != nil {
		return nil, err
	}

	members := make([]*UserGroupInfo, len(userGroups))
	for i, ug := range userGroups {
		members[i] = &UserGroupInfo{
			UserID:   ug.UserID,
			Role:     ug.Role,
			JoinedAt: ug.JoinedAt,
		}
	}

	return members, nil
}

// GetUserGroups returns all groups a user belongs to
func (s *Service) GetUserGroups(userID string) ([]*UserGroup, error) {
	if userID == "" {
		return nil, errors.New("user ID is required")
	}

	return s.repo.GetUserGroups(userID), nil
}

// IsUserInGroup checks if a user is a member of a group
func (s *Service) IsUserInGroup(userID, groupID string) bool {
	return s.repo.IsUserInGroup(userID, groupID)
}

// GetUserRole returns a user's role in a group
func (s *Service) GetUserRole(userID, groupID string) (string, error) {
	return s.repo.GetUserRole(userID, groupID)
}
