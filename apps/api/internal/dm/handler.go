package dm

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"sanctor/internal/auth"
	"sanctor/internal/database"

	"github.com/gorilla/websocket"
)

var (
	repo     Repository = NewRepository()
	service             = NewService(repo)
	hub                 = NewHub()
	upgrader            = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func InitWithDatabase(db *database.DB) {
	repo = NewGormRepository(db)
	service = NewService(repo)
}

func CreateDirectGroup(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req CreateDirectGroupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	group, err := service.CreateOrGetDirectGroup(req.UserID, req.PeerUserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(group)
}

func GetUserGroups(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	userID := r.URL.Query().Get("userId")
	groups, err := service.GetUserGroups(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(groups)
}

func SendMessage(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req SendMessageRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	message, err := service.SendMessage(req)
	if err != nil {
		if errors.Is(err, ErrNotDMMember) {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hub.Broadcast(req.GroupID, message)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(message)
}

func GetMessages(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	groupID := r.URL.Query().Get("groupId")
	userID := r.URL.Query().Get("userId")
	limitStr := r.URL.Query().Get("limit")

	limit := 50
	if limitStr != "" {
		if parsed, err := strconv.Atoi(limitStr); err == nil {
			limit = parsed
		}
	}

	messages, err := service.GetMessages(userID, groupID, limit)
	if err != nil {
		if errors.Is(err, ErrNotDMMember) {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)

	groupID := r.URL.Query().Get("groupId")
	if groupID == "" {
		http.Error(w, "groupId is required", http.StatusBadRequest)
		return
	}

	token := strings.TrimSpace(r.URL.Query().Get("token"))
	if token == "" {
		http.Error(w, "token is required", http.StatusUnauthorized)
		return
	}

	userID, err := auth.ValidateJWT(token)
	if err != nil {
		http.Error(w, "invalid token", http.StatusUnauthorized)
		return
	}

	if !service.IsUserInGroup(userID, groupID) {
		http.Error(w, ErrNotDMMember.Error(), http.StatusForbidden)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	hub.AddConn(groupID, conn)
	defer func() {
		hub.RemoveConn(groupID, conn)
		_ = conn.Close()
	}()

	for {
		var payload struct {
			Content string `json:"content"`
		}

		if err := conn.ReadJSON(&payload); err != nil {
			break
		}

		message, err := service.SendMessage(SendMessageRequest{
			GroupID: groupID,
			UserID:  userID,
			Content: payload.Content,
		})
		if err != nil {
			_ = conn.WriteJSON(map[string]string{"error": err.Error()})
			continue
		}

		hub.Broadcast(groupID, message)
	}
}

func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}
