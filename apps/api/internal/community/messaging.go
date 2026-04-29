package community

import (
	"sanctor/internal/pubsub"
	"time"

	"github.com/google/uuid"
)

// CommunityMessage represents a message in a community
type CommunityMessage struct {
	ID          string    `json:"id"`
	CommunityID string    `json:"groupId"`
	UserID      string    `json:"userId"`
	Content     string    `json:"content"`
	Type        string    `json:"type"` // "text", "notification", "system"
	Timestamp   time.Time `json:"timestamp"`
}

// CommunityEvent represents events that happen in communities
type CommunityEvent struct {
	Type        string      `json:"type"` // "user_joined", "user_left", "message", "group_updated", "group_deleted"
	CommunityID string      `json:"groupId"`
	UserID      string      `json:"userId,omitempty"`
	Data        interface{} `json:"data,omitempty"`
	Timestamp   time.Time   `json:"timestamp"`
}

// Messaging handles community messaging and notifications
type Messaging struct {
	pubsub  *pubsub.PubSub
	service *Service
}

// NewMessaging creates a new community messaging instance
func NewMessaging(ps *pubsub.PubSub, svc *Service) *Messaging {
	return &Messaging{
		pubsub:  ps,
		service: svc,
	}
}

// SendCommunityMessage sends a message to a community
func (m *Messaging) SendCommunityMessage(msg *CommunityMessage) error {
	// Verify user is in community
	userID, err := uuid.Parse(msg.UserID)
	if err != nil {
		return ErrNotMember
	}
	communityID, err := uuid.Parse(msg.CommunityID)
	if err != nil {
		return ErrNotMember
	}
	if !m.service.IsUserInCommunity(userID, communityID) {
		return ErrNotMember
	}

	// Set timestamp if not provided
	if msg.Timestamp.IsZero() {
		msg.Timestamp = time.Now()
	}

	// Publish to group topic (legacy naming)
	topic := "group:" + msg.CommunityID
	m.pubsub.Publish(topic, msg)

	return nil
}

// PublishEvent publishes a community event
func (m *Messaging) PublishEvent(event *CommunityEvent) {
	if event.Timestamp.IsZero() {
		event.Timestamp = time.Now()
	}

	// Publish to group-specific topic (legacy naming)
	communityTopic := "group:" + event.CommunityID
	m.pubsub.Publish(communityTopic, event)

	// Also publish to global group events topic (legacy naming)
	m.pubsub.Publish("group:events", event)
}

// SubscribeToCommunity subscribes to all messages in a community
func (m *Messaging) SubscribeToCommunity(communityID string) (<-chan interface{}, error) {
	// Could add permission check here
	topic := "group:" + communityID
	return m.pubsub.Subscribe(topic), nil
}

// SubscribeToAllCommunityEvents subscribes to all community events
func (m *Messaging) SubscribeToAllCommunityEvents() <-chan interface{} {
	return m.pubsub.Subscribe("group:events")
}

// UnsubscribeFromCommunity unsubscribes from a community
func (m *Messaging) UnsubscribeFromCommunity(communityID string, ch <-chan interface{}) {
	topic := "group:" + communityID
	m.pubsub.Unsubscribe(topic, ch)
}

// NotifyUserJoined sends a notification when a user joins a community
func (m *Messaging) NotifyUserJoined(communityID, userID string) {
	event := &CommunityEvent{
		Type:    "user_joined",
		CommunityID: communityID,
		UserID:  userID,
	}
	m.PublishEvent(event)
}

// NotifyUserLeft sends a notification when a user leaves a community
func (m *Messaging) NotifyUserLeft(communityID, userID string) {
	event := &CommunityEvent{
		Type:    "user_left",
		CommunityID: communityID,
		UserID:  userID,
	}
	m.PublishEvent(event)
}

// NotifyCommunityUpdated sends a notification when a community is updated
func (m *Messaging) NotifyCommunityUpdated(community *Community) {
	event := &CommunityEvent{
		Type:    "group_updated",
		CommunityID: community.ID.String(),
		Data:    community,
	}
	m.PublishEvent(event)
}

// NotifyCommunityDeleted sends a notification when a community is deleted
func (m *Messaging) NotifyCommunityDeleted(communityID string) {
	event := &CommunityEvent{
		Type:    "group_deleted",
		CommunityID: communityID,
	}
	m.PublishEvent(event)
}
