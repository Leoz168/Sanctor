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

func (s *Service) GetUserConversations(userID uuid.UUID) ([]*ConversationSummary, error) {
	if userID == uuid.Nil {
		return nil, errors.New("userId is required")
	}

	groups, err := s.repo.GetUserGroups(userID)
	if err != nil {
		return nil, err
	}

	summaries := make([]*ConversationSummary, 0, len(groups))
	for _, group := range groups {
		members, err := s.repo.GetGroupUsers(group.ID)
		if err != nil {
			return nil, err
		}

		peerUserID := uuid.Nil
		for _, member := range members {
			if member.UserID != userID {
				peerUserID = member.UserID
				break
			}
		}

		latestMessage, err := s.repo.GetLatestGroupMessage(group.ID)
		if err != nil {
			return nil, err
		}

		summary := &ConversationSummary{
			GroupID:    group.ID,
			PeerUserID: peerUserID,
		}
		if latestMessage != nil {
			summary.LastMessage = latestMessage.Content
			summary.LastMessageTime = &latestMessage.MessageTime
		}

		summaries = append(summaries, summary)
	}

	return summaries, nil
}

func (s *Service) SendMessage(userID uuid.UUID, req SendMessageRequest) (*DMMessage, error) {
	if req.GroupID == "" || userID == uuid.Nil {
		return nil, errors.New("groupId and userId are required")
	}

	groupUUID, err := uuid.Parse(req.GroupID)
	if err != nil {
		return nil, errors.New("invalid groupId format")
	}

	content := strings.TrimSpace(req.Content)
	if content == "" {
		return nil, errors.New("content is required")
	}

	if !s.repo.IsUserInGroup(userID, groupUUID) {
		return nil, ErrNotDMMember
	}

	now := time.Now()
	message := &DMMessage{
		ID:          uuid.New(),
		GroupID:     groupUUID,
		UserID:      userID,
		Content:     content,
		MessageTime: now,
		CreatedAt:   now,
		UpdatedAt:   now,
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

func (s *Service) UpdateMessage(userID, messageID uuid.UUID, content string) (*DMMessage, error) {
	if userID == uuid.Nil || messageID == uuid.Nil {
		return nil, errors.New("userId and messageId are required")
	}

	trimmedContent := strings.TrimSpace(content)
	if trimmedContent == "" {
		return nil, errors.New("content is required")
	}

	message, err := s.repo.FindMessageByID(messageID)
	if err != nil {
		return nil, errors.New("message not found")
	}

	if message.UserID != userID {
		return nil, ErrNotDMMember
	}

	message.Content = trimmedContent
	message.UpdatedAt = time.Now()

	if err := s.repo.UpdateMessage(message); err != nil {
		return nil, err
	}

	return message, nil
}

func (s *Service) DeleteMessage(userID, messageID uuid.UUID) error {
	if userID == uuid.Nil || messageID == uuid.Nil {
		return errors.New("userId and messageId are required")
	}

	message, err := s.repo.FindMessageByID(messageID)
	if err != nil {
		return errors.New("message not found")
	}

	if message.UserID != userID {
		return ErrNotDMMember
	}

	return s.repo.DeleteMessage(messageID)
}

func (s *Service) IsUserInGroup(userID, groupID uuid.UUID) bool {
	return s.repo.IsUserInGroup(userID, groupID)
}
