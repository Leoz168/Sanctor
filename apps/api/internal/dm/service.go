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

func (s *Service) CreateOrGetDirectGroup(userID, peerUserID string) (*DMGroup, error) {
	if userID == "" || peerUserID == "" {
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
		ID:        uuid.New().String(),
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

func (s *Service) GetUserGroups(userID string) ([]*DMGroup, error) {
	if userID == "" {
		return nil, errors.New("userId is required")
	}
	return s.repo.GetUserGroups(userID)
}

func (s *Service) SendMessage(req SendMessageRequest) (*DMMessage, error) {
	if req.GroupID == "" || req.UserID == "" {
		return nil, errors.New("groupId and userId are required")
	}

	content := strings.TrimSpace(req.Content)
	if content == "" {
		return nil, errors.New("content is required")
	}

	if !s.repo.IsUserInGroup(req.UserID, req.GroupID) {
		return nil, ErrNotDMMember
	}

	message := &DMMessage{
		ID:          uuid.New().String(),
		GroupID:     req.GroupID,
		UserID:      req.UserID,
		Content:     content,
		MessageTime: time.Now(),
		CreatedAt:   time.Now(),
	}

	if err := s.repo.SaveMessage(message); err != nil {
		return nil, err
	}

	return message, nil
}

func (s *Service) GetMessages(userID, groupID string, limit int) ([]*DMMessage, error) {
	if groupID == "" || userID == "" {
		return nil, errors.New("groupId and userId are required")
	}

	if !s.repo.IsUserInGroup(userID, groupID) {
		return nil, ErrNotDMMember
	}

	return s.repo.GetGroupMessages(groupID, limit)
}

func (s *Service) IsUserInGroup(userID, groupID string) bool {
	return s.repo.IsUserInGroup(userID, groupID)
}
