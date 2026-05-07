package picture

import (
	sharedtypes "sanctor/pkg/types"
	"time"

	"github.com/google/uuid"
)

// Picture represents a picture in the system
type Picture struct {
	ID         uuid.UUID             `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	PostID     *uuid.UUID            `json:"-" gorm:"type:uuid;index"`
	OwnerType  sharedtypes.OwnerType `json:"ownerType" gorm:"type:varchar(40);not null;index:idx_pictures_owner"`
	OwnerID    uuid.UUID             `json:"ownerId" gorm:"type:uuid;not null;index:idx_pictures_owner"`
	URL        string                `json:"url" gorm:"type:varchar(1000);not null"`
	StorageKey string                `json:"storageKey" gorm:"type:varchar(1000);not null;uniqueIndex"`
	Caption    string                `json:"caption" gorm:"type:text"`
	Order      int                   `json:"order" gorm:"column:display_order;default:0"`
	CreatedAt  time.Time             `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt  time.Time             `json:"updatedAt" gorm:"autoUpdateTime"`
}

// CreatePictureRequest represents picture creation data
type CreatePictureRequest struct {
	OwnerType  sharedtypes.OwnerType `json:"ownerType"`
	OwnerID    string                `json:"ownerId"`
	URL        string                `json:"url"`
	StorageKey string                `json:"storageKey"`
	Caption    string                `json:"caption,omitempty"`
	Order      int                   `json:"order,omitempty"`
}

// UpdatePictureRequest represents picture update data
type UpdatePictureRequest struct {
	URL     string `json:"url,omitempty"`
	Caption string `json:"caption,omitempty"`
	Order   int    `json:"order,omitempty"`
}
