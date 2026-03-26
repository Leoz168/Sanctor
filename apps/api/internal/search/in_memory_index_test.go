package search

import "testing"

func TestInMemoryIndex_UpsertAndSearch(t *testing.T) {
	idx := NewInvertedIndex(NewSimpleTokenizer(TokenizerConfig{}))

	err := idx.UpsertEntity(SearchEntity{
		ID:       "post:1",
		Type:     "post",
		EntityID: 1,
		Text:     "Go search basics Learn tokenizer and inverted index",
	})
	if err != nil {
		t.Fatalf("upsert failed: %v", err)
	}

	err = idx.UpsertEntity(SearchEntity{
		ID:       "group:7",
		Type:     "group",
		EntityID: 7,
		Text:     "Backend builders Discuss Go services and APIs",
	})
	if err != nil {
		t.Fatalf("upsert failed: %v", err)
	}

	results := idx.Search("tokenizer", 10)
	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}
	if results[0].Entity.ID != "post:1" {
		t.Fatalf("expected post:1, got %s", results[0].Entity.ID)
	}
}

func TestInMemoryIndex_UpdateReplacesTerms(t *testing.T) {
	idx := NewInvertedIndex(NewSimpleTokenizer(TokenizerConfig{}))

	err := idx.UpsertEntity(SearchEntity{
		ID:   "post:2",
		Text: "first title alpha beta",
	})
	if err != nil {
		t.Fatalf("upsert failed: %v", err)
	}

	err = idx.UpsertEntity(SearchEntity{
		ID:   "post:2",
		Text: "updated title gamma delta",
	})
	if err != nil {
		t.Fatalf("upsert failed: %v", err)
	}

	oldResults := idx.Search("alpha", 10)
	if len(oldResults) != 0 {
		t.Fatalf("expected 0 results for alpha, got %d", len(oldResults))
	}

	newResults := idx.Search("gamma", 10)
	if len(newResults) != 1 {
		t.Fatalf("expected 1 result for gamma, got %d", len(newResults))
	}
}

func TestInMemoryIndex_DeleteRemovesDocument(t *testing.T) {
	idx := NewInvertedIndex(NewSimpleTokenizer(TokenizerConfig{}))

	err := idx.UpsertEntity(SearchEntity{
		ID:   "group:9",
		Text: "Search Club we love relevance tuning",
	})
	if err != nil {
		t.Fatalf("upsert failed: %v", err)
	}

	idx.DeleteEntity("group:9")

	results := idx.Search("relevance", 10)
	if len(results) != 0 {
		t.Fatalf("expected 0 results after delete, got %d", len(results))
	}
}

func TestInMemoryIndex_ValidatesRequiredFields(t *testing.T) {
	idx := NewInvertedIndex(NewSimpleTokenizer(TokenizerConfig{}))

	err := idx.UpsertEntity(SearchEntity{Text: "missing id"})
	if err == nil {
		t.Fatal("expected validation error for missing document id")
	}
}
