package pubsub

import "testing"

func TestPublishNoSubscribers(t *testing.T) {
	ps := NewPubSub()
	ps.Publish("missing", "msg")
}
