package dm

import (
	"time"

	"github.com/google/uuid"
)

type DMGroup struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

type DMGroupUser struct {
	GroupID  uuid.UUID `json:"groupId" gorm:"type:uuid;primaryKey;index"`
	UserID   uuid.UUID `json:"userId" gorm:"type:uuid;primaryKey;index"`
	JoinedAt time.Time `json:"joinedAt" gorm:"autoCreateTime"`
}

type DMMessage struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	GroupID     uuid.UUID `json:"groupId" gorm:"type:uuid;index;not null"`
	UserID      uuid.UUID `json:"userId" gorm:"type:uuid;index;not null"`
	Content     string    `json:"content" gorm:"type:text;not null"`
	MessageTime time.Time `json:"messageTime" gorm:"index;not null"`
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

type CreateDirectGroupRequest struct {
	PeerUserID string `json:"peerUserId"`
}

type SendMessageRequest struct {
	GroupID string `json:"groupId"`
	Content string `json:"content"`
}

type UpdateMessageRequest struct {
	Content string `json:"content"`
}

type ConversationSummary struct {
	GroupID         uuid.UUID  `json:"groupId"`
	PeerUserID      uuid.UUID  `json:"peerUserId"`
	LastMessage     string     `json:"lastMessage,omitempty"`
	LastMessageTime *time.Time `json:"lastMessageTime,omitempty"`
}
