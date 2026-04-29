package events

import "context"

type EventPublisher interface {
	PublishPostCreated(ctx context.Context, event PostCreatedEvent) error
	PublishPostUpdated(ctx context.Context, event PostUpdatedEvent) error
	PublishPostDeleted(ctx context.Context, event PostDeletedEvent) error
	PublishPostViewed(ctx context.Context, event PostViewedEvent) error
}
