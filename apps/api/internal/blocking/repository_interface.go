package blocking

import "github.com/google/uuid"

type Repository interface {
	Create(block *UserBlock) error
	Delete(blockerID, blockeeID uuid.UUID) error
	FindByBlockerID(blockerID uuid.UUID) ([]*UserBlock, error)
	FindByBlockeeID(blockeeID uuid.UUID) ([]*UserBlock, error)
	Exists(blockerID, blockeeID uuid.UUID) (bool, error)
	ExistsUser(userID uuid.UUID) bool
}
