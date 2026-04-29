package blocking

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateBlock(blockerID, blockeeID uuid.UUID) (*UserBlock, error) {
	if blockerID == uuid.Nil {
		return nil, errors.New("blocker ID is required")
	}
	if blockeeID == uuid.Nil {
		return nil, errors.New("blockee ID is required")
	}

	if blockerID == blockeeID {
		return nil, errors.New("user cannot block themselves")
	}

	if !s.repo.ExistsUser(blockerID) {
		return nil, errors.New("blocker user not found")
	}
	if !s.repo.ExistsUser(blockeeID) {
		return nil, errors.New("blockee user not found")
	}

	exists, err := s.repo.Exists(blockerID, blockeeID)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("user is already blocked")
	}

	block := &UserBlock{
		ID:       uuid.New(),
		LinkedAt: time.Now(),
		Blocker:  blockerID,
		Blockee:  blockeeID,
	}

	if err := s.repo.Create(block); err != nil {
		return nil, err
	}

	return block, nil
}

func (s *Service) DeleteBlock(blockerID, blockeeID uuid.UUID) error {
	if blockerID == uuid.Nil {
		return errors.New("blocker ID is required")
	}
	if blockeeID == uuid.Nil {
		return errors.New("blockee ID is required")
	}
	return s.repo.Delete(blockerID, blockeeID)
}

func (s *Service) GetBlocksByBlocker(blockerID uuid.UUID) ([]*UserBlock, error) {
	if blockerID == uuid.Nil {
		return nil, errors.New("blocker ID is required")
	}
	return s.repo.FindByBlockerID(blockerID)
}

func (s *Service) GetBlocksByBlockee(blockeeID uuid.UUID) ([]*UserBlock, error) {
	if blockeeID == uuid.Nil {
		return nil, errors.New("blockee ID is required")
	}
	return s.repo.FindByBlockeeID(blockeeID)
}

func (s *Service) IsBlocked(blockerID, blockeeID uuid.UUID) (bool, error) {
	if blockerID == uuid.Nil {
		return false, errors.New("blocker ID is required")
	}
	if blockeeID == uuid.Nil {
		return false, errors.New("blockee ID is required")
	}
	return s.repo.Exists(blockerID, blockeeID)
}
