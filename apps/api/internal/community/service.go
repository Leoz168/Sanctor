package community

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// Service handles business logic for community operations
type Service struct {
	repo Repository
}

// NewService creates a new community service
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// CreateCommunity creates a new community with validation
func (s *Service) CreateCommunity(req CreateCommunityRequest) (*Community, error) {
	// Validate input
	if req.Name == "" {
		return nil, errors.New("community name is required")
	}

	if req.CreatedBy == "" {
		return nil, errors.New("creator user ID is required")
	}

	if req.InstitutionID == "" {
		return nil, errors.New("institution ID is required")
	}

	creatorUUID, err := uuid.Parse(req.CreatedBy)
	if err != nil {
		return nil, errors.New("invalid creator user ID format")
	}
	institutionUUID, err := uuid.Parse(req.InstitutionID)
	if err != nil {
		return nil, errors.New("invalid institution ID format")
	}

	// Create community
	community := &Community{
		ID:          uuid.New(),
		Name:        req.Name,
		Description: req.Description,
		IsPrivate:   req.IsPrivate,
		CreatedBy:   creatorUUID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := s.repo.Create(community); err != nil {
		return nil, err
	}

	// Automatically add creator as owner
	userCommunity := &UserCommunity{
		UserID:      creatorUUID,
		CommunityID: community.ID,
		Role:        "owner",
		JoinedAt:    time.Now(),
	}

	if err := s.repo.AddUserToCommunity(userCommunity); err != nil {
		// Rollback: delete the community if adding creator fails
		s.repo.Delete(community.ID)
		return nil, errors.New("failed to add creator to community")
	}

	communityInstitution := &CommunityInstitution{
		CommunityID:   community.ID,
		InstitutionID: institutionUUID,
		LinkedAt:      time.Now(),
	}

	if err := s.repo.AddCommunityToInstitution(communityInstitution); err != nil {
		// Rollback: delete the community and owner membership if institution linking fails
		s.repo.Delete(community.ID)
		return nil, errors.New("failed to link community to institution")
	}

	return community, nil
}

// GetCommunity retrieves a community by ID
func (s *Service) GetCommunity(id uuid.UUID) (*Community, error) {
	if id == uuid.Nil {
		return nil, errors.New("community ID is required")
	}

	return s.repo.FindByID(id)
}

// GetCommunityWithMembers retrieves a community with member count
func (s *Service) GetCommunityWithMembers(id uuid.UUID) (*CommunityWithMembers, error) {
	community, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	memberCount := s.repo.GetMemberCount(id)

	return &CommunityWithMembers{
		Community:   community,
		MemberCount: memberCount,
	}, nil
}

// GetAllCommunities retrieves all communities
func (s *Service) GetAllCommunities() ([]*Community, error) {
	return s.repo.FindAll(), nil
}

// UpdateCommunity updates an existing community
func (s *Service) UpdateCommunity(requestingUserID, id uuid.UUID, req UpdateCommunityRequest) (*Community, error) {
	if requestingUserID == uuid.Nil {
		return nil, errors.New("requesting user ID is required")
	}

	community, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("community not found")
	}

	if err := s.ensureCommunityMutationAccess(requestingUserID, community); err != nil {
		return nil, err
	}

	// Update fields if provided
	if req.Name != "" {
		community.Name = req.Name
	}
	if req.Description != "" {
		community.Description = req.Description
	}
	if req.IsPrivate != nil {
		community.IsPrivate = *req.IsPrivate
	}
	community.UpdatedAt = time.Now()

	if err := s.repo.Update(community); err != nil {
		return nil, err
	}

	return community, nil
}

// DeleteCommunity deletes a community by ID
func (s *Service) DeleteCommunity(requestingUserID, id uuid.UUID) error {
	if requestingUserID == uuid.Nil {
		return errors.New("requesting user ID is required")
	}

	if id == uuid.Nil {
		return errors.New("community ID is required")
	}

	community, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("community not found")
	}

	if err := s.ensureCommunityMutationAccess(requestingUserID, community); err != nil {
		return err
	}

	return s.repo.Delete(id)
}

// AddUserToCommunity adds a user to a community
func (s *Service) AddUserToCommunity(req AddUserToCommunityRequest) error {
	if req.UserID == "" || req.CommunityID == "" {
		return errors.New("user ID and community ID are required")
	}

	userUUID, err := uuid.Parse(req.UserID)
	if err != nil {
		return errors.New("invalid user ID format")
	}
	communityUUID, err := uuid.Parse(req.CommunityID)
	if err != nil {
		return errors.New("invalid community ID format")
	}

	// Validate community exists
	if _, err := s.repo.FindByID(communityUUID); err != nil {
		return errors.New("community not found")
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

	userCommunity := &UserCommunity{
		UserID:      userUUID,
		CommunityID: communityUUID,
		Role:        role,
		JoinedAt:    time.Now(),
	}

	return s.repo.AddUserToCommunity(userCommunity)
}

// RemoveUserFromCommunity removes a user from a community
func (s *Service) RemoveUserFromCommunity(userID, communityID uuid.UUID) error {
	if userID == uuid.Nil || communityID == uuid.Nil {
		return errors.New("user ID and community ID are required")
	}

	// Check if user is the owner
	role, err := s.repo.GetUserRole(userID, communityID)
	if err != nil {
		return err
	}

	if role == "owner" {
		// Check if there are other members
		members, _ := s.repo.GetCommunityMembers(communityID)
		if len(members) > 1 {
			return errors.New("owner cannot leave community with other members. Transfer ownership first or delete the community")
		}
	}

	return s.repo.RemoveUserFromCommunity(userID, communityID)
}

// GetCommunityMembers returns all members of a community
func (s *Service) GetCommunityMembers(communityID uuid.UUID) ([]*UserCommunityInfo, error) {
	if communityID == uuid.Nil {
		return nil, errors.New("community ID is required")
	}

	userCommunities, err := s.repo.GetCommunityMembers(communityID)
	if err != nil {
		return nil, err
	}

	members := make([]*UserCommunityInfo, len(userCommunities))
	for i, uc := range userCommunities {
		members[i] = &UserCommunityInfo{
			UserID:   uc.UserID,
			Role:     uc.Role,
			JoinedAt: uc.JoinedAt,
		}
	}

	return members, nil
}

// GetUserCommunities returns all communities a user belongs to
func (s *Service) GetUserCommunities(userID uuid.UUID) ([]*UserCommunity, error) {
	if userID == uuid.Nil {
		return nil, errors.New("user ID is required")
	}

	return s.repo.GetUserCommunities(userID), nil
}

// IsUserInCommunity checks if a user is a member of a community
func (s *Service) IsUserInCommunity(userID, communityID uuid.UUID) bool {
	return s.repo.IsUserInCommunity(userID, communityID)
}

// GetUserRole returns a user's role in a community
func (s *Service) GetUserRole(userID, communityID uuid.UUID) (string, error) {
	return s.repo.GetUserRole(userID, communityID)
}

func (s *Service) ensureCommunityMutationAccess(requestingUserID uuid.UUID, community *Community) error {
	if community == nil {
		return errors.New("community not found")
	}

	if community.CreatedBy == requestingUserID {
		return nil
	}

	role, err := s.repo.GetUserRole(requestingUserID, community.ID)
	if err == nil && role == "owner" {
		return nil
	}

	return errors.New("forbidden: only the creator or owner can modify this community")
}
