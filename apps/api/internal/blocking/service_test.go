package blocking

import "testing"

func TestCreateBlockRejectsSelfBlock(t *testing.T) {
	repo := NewRepository()
	service := NewService(repo)

	req := CreateBlockRequest{
		BlockerID: "3c05d8d4-9cbf-44a7-82bc-dfd0d6912a2a",
		BlockeeID: "3c05d8d4-9cbf-44a7-82bc-dfd0d6912a2a",
	}

	_, err := service.CreateBlock(req)
	if err == nil || err.Error() != "user cannot block themselves" {
		t.Fatalf("expected self-block validation error, got %v", err)
	}
}

func TestCreateBlockRejectsDuplicate(t *testing.T) {
	repo := NewRepository()
	service := NewService(repo)

	req := CreateBlockRequest{
		BlockerID: "3c05d8d4-9cbf-44a7-82bc-dfd0d6912a2a",
		BlockeeID: "8bbca9d5-e6f1-40da-84a2-41f0289e55ef",
	}

	if _, err := service.CreateBlock(req); err != nil {
		t.Fatalf("expected first create to succeed, got %v", err)
	}

	_, err := service.CreateBlock(req)
	if err == nil || err.Error() != "user is already blocked" {
		t.Fatalf("expected duplicate validation error, got %v", err)
	}
}
