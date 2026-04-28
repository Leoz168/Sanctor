package events

import (
	sharedtypes "sanctor/pkg/types"
	"time"
)

type BaseEvent struct {
	EventID   string    `json:"event_id"`
	Timestamp time.Time `json:"timestamp"`
	EventType string    `json:"event_type"`
}

type FieldChange struct {
	Old interface{} `json:"old"`
	New interface{} `json:"new"`
}

const (
	EventTypePostCreated = "PostCreated"
	EventTypePostUpdated = "PostUpdated"
	EventTypePostDeleted = "PostDeleted"
	EventTypePostViewed  = "PostViewed"
)

type PostCreatedEvent struct {
	BaseEvent
	PostID        string             `json:"post_id"`
	AuthorID      string             `json:"author_id"`
	InstitutionID string             `json:"institution_id"`
	CommunityID   string             `json:"community_id"`
	Price         float64            `json:"price"`
	Gender        sharedtypes.Gender `json:"gender"`
}

type PostUpdatedEvent struct {
	BaseEvent
	PostID        string                 `json:"post_id"`
	AuthorID      string                 `json:"author_id"`
	UpdatedFields map[string]FieldChange `json:"updated_fields"`
}

type PostDeletedEvent struct {
	BaseEvent
	PostID   string `json:"post_id"`
	AuthorID string `json:"author_id"`
}

type PostViewedEvent struct {
	BaseEvent
	PostID string `json:"post_id"`
}
