package dm

import "errors"

var (
	ErrDMGroupNotFound = errors.New("dm group not found")
	ErrNotDMMember     = errors.New("user is not a member of this dm group")
)

type Repository interface {
	CreateGroup(group *DMGroup) error
	AddUserToGroup(groupUser *DMGroupUser) error
	GetGroupUsers(groupID string) ([]*DMGroupUser, error)
	GetUserGroups(userID string) ([]*DMGroup, error)
	FindDirectGroupByUsers(userA, userB string) (*DMGroup, error)
	IsUserInGroup(userID, groupID string) bool
	SaveMessage(message *DMMessage) error
	GetGroupMessages(groupID string, limit int) ([]*DMMessage, error)
}
