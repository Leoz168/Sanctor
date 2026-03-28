package dm

import "time"

type DMGroup struct {
	ID        string    `json:"id" gorm:"type:uuid;primaryKey"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

type DMGroupUser struct {
	GroupID  string    `json:"groupId" gorm:"type:uuid;primaryKey;index"`
	UserID   string    `json:"userId" gorm:"type:uuid;primaryKey;index"`
	JoinedAt time.Time `json:"joinedAt" gorm:"autoCreateTime"`
}

type DMMessage struct {
	ID          string    `json:"id" gorm:"type:uuid;primaryKey"`
	GroupID     string    `json:"groupId" gorm:"type:uuid;index;not null"`
	UserID      string    `json:"userId" gorm:"type:uuid;index;not null"`
	Content     string    `json:"content" gorm:"type:text;not null"`
	MessageTime time.Time `json:"messageTime" gorm:"index;not null"`
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime"`
}

type CreateDirectGroupRequest struct {
	UserID     string `json:"userId"`
	PeerUserID string `json:"peerUserId"`
}

type SendMessageRequest struct {
	GroupID string `json:"groupId"`
	UserID  string `json:"userId"`
	Content string `json:"content"`
}
