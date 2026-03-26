package search

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"
)

type EntityType string

const (
	PostEntity  EntityType = "post"
	GroupEntity EntityType = "group"
	UserEntity  EntityType = "user"
)

// SearchEntity is the normalized entity shape stored in the in-memory index.
type SearchEntity struct {
	ID       string
	Type     EntityType //e.g. post, group, user
	EntityID int64
	Text     string
}

// TermData stores term statistics for one entity.
type TermData struct {
	TermFrequency int
	Positions     []int
}

// SearchResult is one ranked match returned by the index.
type SearchResult struct {
	Entity SearchEntity
	Score  float64
}

// InvertedIndex keeps a mutable, RAM-backed inverted index.
type InvertedIndex struct {
	mu          sync.RWMutex
	tokenizer   Tokenizer
	entities    map[string]SearchEntity
	inverted    map[string]map[string]TermData // tokens -> entityID -> termData
	entityTerms map[string]map[string]struct{}
}

// NewInvertedIndex creates an empty in-memory index.
func NewInvertedIndex(tokenizer Tokenizer) *InvertedIndex {
	return &InvertedIndex{
		tokenizer:   tokenizer,
		entities:    make(map[string]SearchEntity),
		inverted:    make(map[string]map[string]TermData),
		entityTerms: make(map[string]map[string]struct{}),
	}
}

// UpsertEntity inserts or updates one entity in the index.
func (idx *InvertedIndex) UpsertEntity(ent SearchEntity) error {
	if idx.tokenizer == nil {
		return fmt.Errorf("tokenizer is required")
	}
	if strings.TrimSpace(ent.ID) == "" {
		return fmt.Errorf("entity id is required")
	}

	tokens := idx.tokenizer.Tokenize(strings.TrimSpace(ent.Text))
	termDataMap := make(map[string]TermData)
	termSet := make(map[string]struct{})

	for _, token := range tokens {
		termData := termDataMap[token.Term]
		termData.TermFrequency++
		termData.Positions = append(termData.Positions, token.Pos)
		termDataMap[token.Term] = termData
		termSet[token.Term] = struct{}{}
	}

	idx.mu.Lock()
	defer idx.mu.Unlock()

	idx.removeEntityLocked(ent.ID)
	idx.entities[ent.ID] = ent
	idx.entityTerms[ent.ID] = termSet

	for term, data := range termDataMap {
		entityMap, ok := idx.inverted[term]
		if !ok {
			entityMap = make(map[string]TermData)
			idx.inverted[term] = entityMap
		}
		entityMap[ent.ID] = data
	}

	return nil
}

// DeleteEntity removes one entity and all of its postings.
func (idx *InvertedIndex) DeleteEntity(entityID string) {
	idx.mu.Lock()
	defer idx.mu.Unlock()
	idx.removeEntityLocked(entityID)
}

func (idx *InvertedIndex) removeEntityLocked(entityID string) {
	if entityID == "" {
		return
	}

	terms, ok := idx.entityTerms[entityID]
	if ok {
		for term := range terms {
			entityMap, ok := idx.inverted[term]
			if ok {
				delete(entityMap, entityID)
				if len(entityMap) == 0 {
					delete(idx.inverted, term)
				}
			}
		}
	}

	delete(idx.entityTerms, entityID)
	delete(idx.entities, entityID)
}

// Search tokenizes the query, calculates simple TF-IDF style scores, and returns ranked matches.
func (idx *InvertedIndex) Search(query string, limit int) []SearchResult {
	if idx.tokenizer == nil {
		return nil
	}

	queryTokens := idx.tokenizer.Tokenize(query)
	if len(queryTokens) == 0 {
		return nil
	}

	queryTerms := make(map[string]struct{}, len(queryTokens))
	for _, token := range queryTokens {
		queryTerms[token.Term] = struct{}{}
	}

	idx.mu.RLock()
	defer idx.mu.RUnlock()

	totalDocs := len(idx.entities)
	if totalDocs == 0 {
		return nil
	}

	scores := make(map[string]float64)
	for term := range queryTerms {
		docPostings, ok := idx.inverted[term]
		if !ok || len(docPostings) == 0 {
			continue
		}

		df := float64(len(docPostings))
		idf := math.Log(1.0 + float64(totalDocs)/(1.0+df))
		for docID, posting := range docPostings {
			scores[docID] += float64(posting.TermFrequency) * idf
		}
	}

	results := make([]SearchResult, 0, len(scores))
	for docID, score := range scores {
		if score <= 0 {
			continue
		}
		doc, ok := idx.entities[docID]
		if !ok {
			continue
		}
		results = append(results, SearchResult{Entity: doc, Score: score})
	}

	sort.Slice(results, func(i, j int) bool {
		if results[i].Score == results[j].Score {
			return results[i].Entity.ID < results[j].Entity.ID
		}
		return results[i].Score > results[j].Score
	})

	if limit > 0 && len(results) > limit {
		return results[:limit]
	}

	return results
}
