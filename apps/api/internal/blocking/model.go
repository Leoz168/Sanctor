package blocking

import (
	"time"

	"github.com/google/uuid"
)

// UserBlock represents a directed block relationship between two users.
type UserBlock struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	LinkedAt time.Time `json:"linkedAt" gorm:"column:linked_at;not null;index"`
	Blocker  uuid.UUID `json:"blocker" gorm:"type:uuid;not null;index:idx_user_blocking_blocker"`
	Blockee  uuid.UUID `json:"blockee" gorm:"type:uuid;not null;index:idx_user_blocking_blockee"`
}

func (UserBlock) TableName() string {
	return "user_blocking"
}

type CreateBlockRequest struct {
	BlockerID uuid.UUID `json:"blockerId"`
	BlockeeID uuid.UUID `json:"blockeeId"`
}

type BlockStatusResponse struct {
	BlockerID uuid.UUID `json:"blockerId"`
	BlockeeID uuid.UUID `json:"blockeeId"`
	IsBlocked bool      `json:"isBlocked"`
}
