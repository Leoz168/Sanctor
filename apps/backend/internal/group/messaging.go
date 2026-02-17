package group

import (
	"sanctor/internal/pubsub"
	"time"
)

// Message represents a message in a group
type Message struct {
	ID        string    `json:"id"`
	GroupID   string    `json:"groupId"`
	UserID    string    `json:"userId"`
	Content   string    `json:"content"`
	Type      string    `json:"type"` // "text", "notification", "system"
	Timestamp time.Time `json:"timestamp"`
}

// GroupEvent represents events that happen in groups
type GroupEvent struct {
	Type      string    `json:"type"` // "user_joined", "user_left", "message", "group_updated", "group_deleted"
	GroupID   string    `json:"groupId"`
	UserID    string    `json:"userId,omitempty"`
	Data      interface{} `json:"data,omitempty"`
	Timestamp time.Time `json:"timestamp"`
}

// Messaging handles group messaging and notifications
type Messaging struct {
	pubsub  *pubsub.PubSub
	service *Service
}

// NewMessaging creates a new group messaging instance
func NewMessaging(ps *pubsub.PubSub, svc *Service) *Messaging {
	return &Messaging{
		pubsub:  ps,
		service: svc,
	}
}

// SendMessage sends a message to a group
func (m *Messaging) SendMessage(msg *Message) error {
	// Verify user is in group
	if !m.service.IsUserInGroup(msg.UserID, msg.GroupID) {
		return ErrNotMember
	}

	// Set timestamp if not provided
	if msg.Timestamp.IsZero() {
		msg.Timestamp = time.Now()
	}

	// Publish to group topic
	topic := "group:" + msg.GroupID
	m.pubsub.Publish(topic, msg)

	return nil
}

// PublishEvent publishes a group event
func (m *Messaging) PublishEvent(event *GroupEvent) {
	if event.Timestamp.IsZero() {
		event.Timestamp = time.Now()
	}

	// Publish to group-specific topic
	groupTopic := "group:" + event.GroupID
	m.pubsub.Publish(groupTopic, event)

	// Also publish to global group events topic
	m.pubsub.Publish("group:events", event)
}

// SubscribeToGroup subscribes to all messages in a group
func (m *Messaging) SubscribeToGroup(groupID string) (<-chan interface{}, error) {
	// Could add permission check here
	topic := "group:" + groupID
	return m.pubsub.Subscribe(topic), nil
}

// SubscribeToAllGroupEvents subscribes to all group events
func (m *Messaging) SubscribeToAllGroupEvents() <-chan interface{} {
	return m.pubsub.Subscribe("group:events")
}

// UnsubscribeFromGroup unsubscribes from a group
func (m *Messaging) UnsubscribeFromGroup(groupID string, ch <-chan interface{}) {
	topic := "group:" + groupID
	m.pubsub.Unsubscribe(topic, ch)
}

// NotifyUserJoined sends a notification when a user joins a group
func (m *Messaging) NotifyUserJoined(groupID, userID string) {
	event := &GroupEvent{
		Type:    "user_joined",
		GroupID: groupID,
		UserID:  userID,
	}
	m.PublishEvent(event)
}

// NotifyUserLeft sends a notification when a user leaves a group
func (m *Messaging) NotifyUserLeft(groupID, userID string) {
	event := &GroupEvent{
		Type:    "user_left",
		GroupID: groupID,
		UserID:  userID,
	}
	m.PublishEvent(event)
}

// NotifyGroupUpdated sends a notification when a group is updated
func (m *Messaging) NotifyGroupUpdated(group *Group) {
	event := &GroupEvent{
		Type:    "group_updated",
		GroupID: group.ID,
		Data:    group,
	}
	m.PublishEvent(event)
}

// NotifyGroupDeleted sends a notification when a group is deleted
func (m *Messaging) NotifyGroupDeleted(groupID string) {
	event := &GroupEvent{
		Type:    "group_deleted",
		GroupID: groupID,
	}
	m.PublishEvent(event)
}
