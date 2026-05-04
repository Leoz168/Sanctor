package pubsub

import (
	"testing"
	"time"
)

func TestPublishAndUnsubscribe(t *testing.T) {
	ps := NewPubSub()
	ch := ps.Subscribe("topic")

	ps.Publish("topic", "hello")
	select {
	case msg := <-ch:
		if msg != "hello" {
			t.Fatalf("expected message hello, got %v", msg)
		}
	case <-time.After(100 * time.Millisecond):
		t.Fatalf("expected message to be delivered")
	}

	ps.Unsubscribe("topic", ch)
	_, ok := <-ch
	if ok {
		t.Fatalf("expected channel to be closed")
	}
}
