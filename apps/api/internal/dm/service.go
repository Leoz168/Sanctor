package dm

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateOrGetDirectGroup(userID, peerUserID uuid.UUID) (*DMGroup, error) {
	if userID == uuid.Nil || peerUserID == uuid.Nil {
		return nil, errors.New("userId and peerUserId are required")
	}
	if userID == peerUserID {
		return nil, errors.New("direct message requires two different users")
	}

	group, err := s.repo.FindDirectGroupByUsers(userID, peerUserID)
	if err == nil {
		return group, nil
	}
	if !errors.Is(err, ErrDMGroupNotFound) {
		return nil, err
	}

	now := time.Now()
	newGroup := &DMGroup{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := s.repo.CreateGroup(newGroup); err != nil {
		return nil, err
	}

	if err := s.repo.AddUserToGroup(&DMGroupUser{GroupID: newGroup.ID, UserID: userID, JoinedAt: now}); err != nil {
		return nil, err
	}

	if err := s.repo.AddUserToGroup(&DMGroupUser{GroupID: newGroup.ID, UserID: peerUserID, JoinedAt: now}); err != nil {
		return nil, err
	}

	return newGroup, nil
}

func (s *Service) GetUserGroups(userID uuid.UUID) ([]*DMGroup, error) {
	if userID == uuid.Nil {
		return nil, errors.New("userId is required")
	}
	return s.repo.GetUserGroups(userID)
}

func (s *Service) SendMessage(req SendMessageRequest) (*DMMessage, error) {
	if req.GroupID == "" || req.UserID == "" {
		return nil, errors.New("groupId and userId are required")
	}

	groupUUID, err := uuid.Parse(req.GroupID)
	if err != nil {
		return nil, errors.New("invalid groupId format")
	}
	userUUID, err := uuid.Parse(req.UserID)
	if err != nil {
		return nil, errors.New("invalid userId format")
	}

	content := strings.TrimSpace(req.Content)
	if content == "" {
		return nil, errors.New("content is required")
	}

	if !s.repo.IsUserInGroup(userUUID, groupUUID) {
		return nil, ErrNotDMMember
	}

	now := time.Now()
	message := &DMMessage{
		ID:          uuid.New(),
		GroupID:     groupUUID,
		UserID:      userUUID,
		Content:     content,
		MessageTime: now,
		CreatedAt:   now,
	}

	if err := s.repo.SaveMessage(message); err != nil {
		return nil, err
	}

	return message, nil
}

func (s *Service) GetMessages(userID, groupID uuid.UUID, limit int) ([]*DMMessage, error) {
	if groupID == uuid.Nil || userID == uuid.Nil {
		return nil, errors.New("groupId and userId are required")
	}

	if !s.repo.IsUserInGroup(userID, groupID) {
		return nil, ErrNotDMMember
	}

	return s.repo.GetGroupMessages(groupID, limit)
}

func (s *Service) IsUserInGroup(userID, groupID uuid.UUID) bool {
	return s.repo.IsUserInGroup(userID, groupID)
}
