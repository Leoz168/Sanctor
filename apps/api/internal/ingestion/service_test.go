package ingestion

import "testing"

func TestIngestionServiceNoOps(t *testing.T) {
	svc := NewService()
	if svc == nil {
		t.Fatal("expected service instance")
	}
	if err := svc.IngestData(map[string]string{"key": "value"}); err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if err := svc.ProcessBatch([]interface{}{"one", 2}); err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
}
