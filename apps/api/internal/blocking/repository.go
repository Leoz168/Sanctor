package blocking

import (
	"fmt"
	"sort"

	"github.com/google/uuid"
)

type InMemoryRepository struct {
	blocks map[string]*UserBlock
}

func NewRepository() Repository {
	return &InMemoryRepository{blocks: make(map[string]*UserBlock)}
}

func (r *InMemoryRepository) Create(block *UserBlock) error {
	r.blocks[blockKey(block.Blocker, block.Blockee)] = block
	return nil
}

func (r *InMemoryRepository) Delete(blockerID, blockeeID uuid.UUID) error {
	delete(r.blocks, blockKey(blockerID, blockeeID))
	return nil
}

func (r *InMemoryRepository) FindByBlockerID(blockerID uuid.UUID) ([]*UserBlock, error) {
	blocks := make([]*UserBlock, 0)
	for _, block := range r.blocks {
		if block.Blocker == blockerID {
			blocks = append(blocks, block)
		}
	}
	sort.Slice(blocks, func(i, j int) bool {
		return blocks[i].LinkedAt > blocks[j].LinkedAt
	})
	return blocks, nil
}

func (r *InMemoryRepository) FindByBlockeeID(blockeeID uuid.UUID) ([]*UserBlock, error) {
	blocks := make([]*UserBlock, 0)
	for _, block := range r.blocks {
		if block.Blockee == blockeeID {
			blocks = append(blocks, block)
		}
	}
	sort.Slice(blocks, func(i, j int) bool {
		return blocks[i].LinkedAt > blocks[j].LinkedAt
	})
	return blocks, nil
}

func (r *InMemoryRepository) Exists(blockerID, blockeeID uuid.UUID) (bool, error) {
	_, exists := r.blocks[blockKey(blockerID, blockeeID)]
	return exists, nil
}

func (r *InMemoryRepository) ExistsUser(userID uuid.UUID) bool {
	return userID != uuid.Nil
}

func blockKey(blockerID, blockeeID uuid.UUID) string {
	return fmt.Sprintf("%s:%s", blockerID, blockeeID)
}
