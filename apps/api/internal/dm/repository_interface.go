package dm

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrDMGroupNotFound = errors.New("dm group not found")
	ErrNotDMMember     = errors.New("user is not a member of this dm group")
)

type Repository interface {
	CreateGroup(group *DMGroup) error
	AddUserToGroup(groupUser *DMGroupUser) error
	GetGroupUsers(groupID uuid.UUID) ([]*DMGroupUser, error)
	GetUserGroups(userID uuid.UUID) ([]*DMGroup, error)
	FindDirectGroupByUsers(userA, userB uuid.UUID) (*DMGroup, error)
	IsUserInGroup(userID, groupID uuid.UUID) bool
	SaveMessage(message *DMMessage) error
	GetGroupMessages(groupID uuid.UUID, limit int) ([]*DMMessage, error)
	GetLatestGroupMessage(groupID uuid.UUID) (*DMMessage, error)
	FindMessageByID(messageID uuid.UUID) (*DMMessage, error)
	UpdateMessage(message *DMMessage) error
	DeleteMessage(messageID uuid.UUID) error
}
