package dm

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Hub struct {
	groupConns map[string]map[*websocket.Conn]bool
	mu         sync.RWMutex
}

func NewHub() *Hub {
	return &Hub{
		groupConns: make(map[string]map[*websocket.Conn]bool),
	}
}

func (h *Hub) AddConn(groupID string, conn *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if _, ok := h.groupConns[groupID]; !ok {
		h.groupConns[groupID] = make(map[*websocket.Conn]bool)
	}

	h.groupConns[groupID][conn] = true
}

func (h *Hub) RemoveConn(groupID string, conn *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if group, ok := h.groupConns[groupID]; ok {
		delete(group, conn)
		if len(group) == 0 {
			delete(h.groupConns, groupID)
		}
	}
}

func (h *Hub) Broadcast(groupID string, payload interface{}) {
	h.mu.RLock()
	conns := make([]*websocket.Conn, 0)
	for conn := range h.groupConns[groupID] {
		conns = append(conns, conn)
	}
	h.mu.RUnlock()

	for _, conn := range conns {
		if err := conn.WriteJSON(payload); err != nil {
			h.RemoveConn(groupID, conn)
			_ = conn.Close()
		}
	}
}
