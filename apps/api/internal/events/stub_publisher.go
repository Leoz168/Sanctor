package events

import "context"

type StubPublisher struct{}

func NewStubPublisher() *StubPublisher {
	return &StubPublisher{}
}

func (p *StubPublisher) PublishPostCreated(ctx context.Context, event PostCreatedEvent) error {
	// Stub implementation - do nothing
	return nil
}

func (p *StubPublisher) PublishPostUpdated(ctx context.Context, event PostUpdatedEvent) error {
	// Stub implementation - do nothing
	return nil
}

func (p *StubPublisher) PublishPostDeleted(ctx context.Context, event PostDeletedEvent) error {
	// Stub implementation - do nothing
	return nil
}

func (p *StubPublisher) PublishPostViewed(ctx context.Context, event PostViewedEvent) error {
	// Stub implementation - do nothing
	return nil
}
