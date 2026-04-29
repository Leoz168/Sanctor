package blocking

import "github.com/google/uuid"

// UserBlock represents a directed block relationship between two users.
type UserBlock struct {
	ID       int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	LinkedAt int64     `json:"linkedAt" gorm:"column:linked_at;not null;index"`
	Blocker  uuid.UUID `json:"blocker" gorm:"type:uuid;not null;index:idx_user_blocking_blocker"`
	Blockee  uuid.UUID `json:"blockee" gorm:"type:uuid;not null;index:idx_user_blocking_blockee"`
}

func (UserBlock) TableName() string {
	return "user_blocking"
}

type CreateBlockRequest struct {
	BlockerID string `json:"blockerId"`
	BlockeeID string `json:"blockeeId"`
}

type BlockStatusResponse struct {
	BlockerID string `json:"blockerId"`
	BlockeeID string `json:"blockeeId"`
	IsBlocked bool   `json:"isBlocked"`
}
