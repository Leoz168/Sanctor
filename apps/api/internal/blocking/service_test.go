package blocking

import (
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
)

var (
	blockerID = uuid.MustParse("3c05d8d4-9cbf-44a7-82bc-dfd0d6912a2a")
	blockeeID = uuid.MustParse("8bbca9d5-e6f1-40da-84a2-41f0289e55ef")
)

type stubRepository struct {
	createErr       error
	deleteErr       error
	exists          bool
	existsErr       error
	existsUser      map[uuid.UUID]bool
	blocksByBlocker []*UserBlock
	blocksByBlockee []*UserBlock
	findByBlockerErr error
	findByBlockeeErr error
	createdBlock    *UserBlock
}

func (s *stubRepository) Create(block *UserBlock) error {
	s.createdBlock = block
	return s.createErr
}

func (s *stubRepository) Delete(blockerID, blockeeID uuid.UUID) error {
	return s.deleteErr
}

func (s *stubRepository) FindByBlockerID(blockerID uuid.UUID) ([]*UserBlock, error) {
	return s.blocksByBlocker, s.findByBlockerErr
}

func (s *stubRepository) FindByBlockeeID(blockeeID uuid.UUID) ([]*UserBlock, error) {
	return s.blocksByBlockee, s.findByBlockeeErr
}

func (s *stubRepository) Exists(blockerID, blockeeID uuid.UUID) (bool, error) {
	return s.exists, s.existsErr
}

func (s *stubRepository) ExistsUser(userID uuid.UUID) bool {
	if s.existsUser == nil {
		return userID != uuid.Nil
	}
	return s.existsUser[userID]
}

func TestCreateBlockRejectsNilIDs(t *testing.T) {
	cases := []struct {
		name      string
		blockerID uuid.UUID
		blockeeID uuid.UUID
		wantErr   string
	}{
		{name: "nil-blocker", blockerID: uuid.Nil, blockeeID: blockeeID, wantErr: "blocker ID is required"},
		{name: "nil-blockee", blockerID: blockerID, blockeeID: uuid.Nil, wantErr: "blockee ID is required"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			service := NewService(&stubRepository{})
			_, err := service.CreateBlock(tc.blockerID, tc.blockeeID)
			if err == nil || err.Error() != tc.wantErr {
				t.Fatalf("expected %q, got %v", tc.wantErr, err)
			}
		})
	}
}

func TestCreateBlockRejectsSelfBlock(t *testing.T) {
	repo := NewRepository()
	service := NewService(repo)

	_, err := service.CreateBlock(blockerID, blockerID)
	if err == nil || err.Error() != "user cannot block themselves" {
		t.Fatalf("expected self-block validation error, got %v", err)
	}
}

func TestCreateBlockRejectsMissingUsers(t *testing.T) {
	cases := []struct {
		name      string
		existsUser map[uuid.UUID]bool
		wantErr   string
	}{
		{
			name:      "missing-blocker",
			existsUser: map[uuid.UUID]bool{blockerID: false, blockeeID: true},
			wantErr:   "blocker user not found",
		},
		{
			name:      "missing-blockee",
			existsUser: map[uuid.UUID]bool{blockerID: true, blockeeID: false},
			wantErr:   "blockee user not found",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			repo := &stubRepository{existsUser: tc.existsUser}
			service := NewService(repo)

			_, err := service.CreateBlock(blockerID, blockeeID)
			if err == nil || err.Error() != tc.wantErr {
				t.Fatalf("expected %q, got %v", tc.wantErr, err)
			}
		})
	}
}

func TestCreateBlockRejectsWhenExistsFails(t *testing.T) {
	repo := &stubRepository{
		existsErr:  errors.New("db down"),
		existsUser: map[uuid.UUID]bool{blockerID: true, blockeeID: true},
	}
	service := NewService(repo)

	_, err := service.CreateBlock(blockerID, blockeeID)
	if err == nil || err.Error() != "db down" {
		t.Fatalf("expected repository error, got %v", err)
	}
}

func TestCreateBlockRejectsWhenCreateFails(t *testing.T) {
	repo := &stubRepository{
		createErr:  errors.New("insert failed"),
		existsUser: map[uuid.UUID]bool{blockerID: true, blockeeID: true},
	}
	service := NewService(repo)

	_, err := service.CreateBlock(blockerID, blockeeID)
	if err == nil || err.Error() != "insert failed" {
		t.Fatalf("expected repository error, got %v", err)
	}
}

func TestCreateBlockRejectsDuplicate(t *testing.T) {
	repo := NewRepository()
	service := NewService(repo)

	if _, err := service.CreateBlock(blockerID, blockeeID); err != nil {
		t.Fatalf("expected first create to succeed, got %v", err)
	}

	_, err := service.CreateBlock(blockerID, blockeeID)
	if err == nil || err.Error() != "user is already blocked" {
		t.Fatalf("expected duplicate validation error, got %v", err)
	}
}

func TestCreateBlockSuccess(t *testing.T) {
	repo := &stubRepository{existsUser: map[uuid.UUID]bool{blockerID: true, blockeeID: true}}
	service := NewService(repo)
	start := time.Now()

	block, err := service.CreateBlock(blockerID, blockeeID)
	if err != nil {
		t.Fatalf("expected create to succeed, got %v", err)
	}
	if block.ID == uuid.Nil {
		t.Fatal("expected block ID to be set")
	}
	if block.Blocker != blockerID || block.Blockee != blockeeID {
		t.Fatalf("expected blocker/blockee IDs to match inputs")
	}
	if block.LinkedAt.IsZero() || block.LinkedAt.Before(start) {
		t.Fatalf("expected linked time to be set")
	}
	if repo.createdBlock != block {
		t.Fatalf("expected created block to be persisted")
	}
}

func TestDeleteBlockRejectsNilIDs(t *testing.T) {
	cases := []struct {
		name      string
		blockerID uuid.UUID
		blockeeID uuid.UUID
		wantErr   string
	}{
		{name: "nil-blocker", blockerID: uuid.Nil, blockeeID: blockeeID, wantErr: "blocker ID is required"},
		{name: "nil-blockee", blockerID: blockerID, blockeeID: uuid.Nil, wantErr: "blockee ID is required"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			service := NewService(&stubRepository{})
			err := service.DeleteBlock(tc.blockerID, tc.blockeeID)
			if err == nil || err.Error() != tc.wantErr {
				t.Fatalf("expected %q, got %v", tc.wantErr, err)
			}
		})
	}
}

func TestGetBlocksByBlockerRejectsNil(t *testing.T) {
	service := NewService(&stubRepository{})

	_, err := service.GetBlocksByBlocker(uuid.Nil)
	if err == nil || err.Error() != "blocker ID is required" {
		t.Fatalf("expected blocker ID validation error, got %v", err)
	}
}

func TestGetBlocksByBlockerReturnsRepoResults(t *testing.T) {
	blocks := []*UserBlock{{ID: uuid.New(), Blocker: blockerID, Blockee: blockeeID}}
	repo := &stubRepository{blocksByBlocker: blocks}
	service := NewService(repo)

	result, err := service.GetBlocksByBlocker(blockerID)
	if err != nil {
		t.Fatalf("expected fetch to succeed, got %v", err)
	}
	if len(result) != 1 || result[0] != blocks[0] {
		t.Fatalf("expected repository results to be returned")
	}
}

func TestGetBlocksByBlockeeRejectsNil(t *testing.T) {
	service := NewService(&stubRepository{})

	_, err := service.GetBlocksByBlockee(uuid.Nil)
	if err == nil || err.Error() != "blockee ID is required" {
		t.Fatalf("expected blockee ID validation error, got %v", err)
	}
}

func TestGetBlocksByBlockeeReturnsRepoResults(t *testing.T) {
	blocks := []*UserBlock{{ID: uuid.New(), Blocker: blockerID, Blockee: blockeeID}}
	repo := &stubRepository{blocksByBlockee: blocks}
	service := NewService(repo)

	result, err := service.GetBlocksByBlockee(blockeeID)
	if err != nil {
		t.Fatalf("expected fetch to succeed, got %v", err)
	}
	if len(result) != 1 || result[0] != blocks[0] {
		t.Fatalf("expected repository results to be returned")
	}
}

func TestIsBlockedRejectsNilIDs(t *testing.T) {
	cases := []struct {
		name      string
		blockerID uuid.UUID
		blockeeID uuid.UUID
		wantErr   string
	}{
		{name: "nil-blocker", blockerID: uuid.Nil, blockeeID: blockeeID, wantErr: "blocker ID is required"},
		{name: "nil-blockee", blockerID: blockerID, blockeeID: uuid.Nil, wantErr: "blockee ID is required"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			service := NewService(&stubRepository{})
			_, err := service.IsBlocked(tc.blockerID, tc.blockeeID)
			if err == nil || err.Error() != tc.wantErr {
				t.Fatalf("expected %q, got %v", tc.wantErr, err)
			}
		})
	}
}

func TestIsBlockedReturnsRepoValue(t *testing.T) {
	repo := &stubRepository{exists: true}
	service := NewService(repo)

	blocked, err := service.IsBlocked(blockerID, blockeeID)
	if err != nil {
		t.Fatalf("expected check to succeed, got %v", err)
	}
	if !blocked {
		t.Fatalf("expected block to be true")
	}
}
