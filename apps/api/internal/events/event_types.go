package events

import (
	sharedtypes "sanctor/pkg/types"
	"time"

	"github.com/google/uuid"
)

type BaseEvent struct {
	EventID   uuid.UUID `json:"event_id"`
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
	PostID         uuid.UUID          `json:"post_id"`
	AuthorID       uuid.UUID          `json:"author_id"`
	InstitutionIDs []uuid.UUID        `json:"institution_ids"`
	CommunityID    uuid.UUID          `json:"community_id"`
	Price          int64              `json:"price"`
	Gender         sharedtypes.Gender `json:"gender"`
}

type PostUpdatedEvent struct {
	BaseEvent
	PostID        uuid.UUID              `json:"post_id"`
	AuthorID      uuid.UUID              `json:"author_id"`
	UpdatedFields map[string]FieldChange `json:"updated_fields"`
}

type PostDeletedEvent struct {
	BaseEvent
	PostID   uuid.UUID `json:"post_id"`
	AuthorID uuid.UUID `json:"author_id"`
}

type PostViewedEvent struct {
	BaseEvent
	PostID uuid.UUID `json:"post_id"`
}
