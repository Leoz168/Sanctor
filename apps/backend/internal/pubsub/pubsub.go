package pubsub

import "log"

// PubSub handles publish/subscribe messaging
type PubSub struct {
	subscribers map[string][]chan interface{}
}

// NewPubSub creates a new pub/sub instance
func NewPubSub() *PubSub {
	return &PubSub{
		subscribers: make(map[string][]chan interface{}),
	}
}

// Subscribe registers a subscriber for a topic
func (ps *PubSub) Subscribe(topic string) <-chan interface{} {
	ch := make(chan interface{}, 10)
	ps.subscribers[topic] = append(ps.subscribers[topic], ch)
	return ch
}

// Publish sends a message to all subscribers of a topic
func (ps *PubSub) Publish(topic string, message interface{}) {
	if subs, exists := ps.subscribers[topic]; exists {
		for _, ch := range subs {
			select {
			case ch <- message:
			default:
				log.Printf("Subscriber channel full for topic: %s", topic)
			}
		}
	}
}

// Unsubscribe removes a subscriber from a topic
func (ps *PubSub) Unsubscribe(topic string, ch <-chan interface{}) {
	if subs, exists := ps.subscribers[topic]; exists {
		for i, subscriber := range subs {
			if subscriber == ch {
				ps.subscribers[topic] = append(subs[:i], subs[i+1:]...)
				close(subscriber)
				break
			}
		}
	}
}
