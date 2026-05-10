package dm

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"sanctor/internal/auth"
	"sanctor/internal/database"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

func currentUserIDFromContext(r *http.Request) (uuid.UUID, error) {
	userID, ok := r.Context().Value("userId").(string)
	if !ok || userID == "" {
		return uuid.Nil, http.ErrNoCookie
	}

	return uuid.Parse(userID)
}

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

	userID, err := currentUserIDFromContext(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	peerUserID, err := uuid.Parse(strings.TrimSpace(req.PeerUserID))
	if err != nil {
		http.Error(w, "invalid peerUserId format", http.StatusBadRequest)
		return
	}

	group, err := service.CreateOrGetDirectGroup(userID, peerUserID)
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

	userUUID, err := currentUserIDFromContext(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	groups, err := service.GetUserConversations(userUUID)
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

	userID, err := currentUserIDFromContext(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	message, err := service.SendMessage(userID, req)
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
	limitStr := r.URL.Query().Get("limit")

	limit := 50
	if limitStr != "" {
		if parsed, err := strconv.Atoi(limitStr); err == nil {
			limit = parsed
		}
	}

	userUUID, err := currentUserIDFromContext(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	groupUUID, err := uuid.Parse(strings.TrimSpace(groupID))
	if err != nil {
		http.Error(w, "invalid groupId format", http.StatusBadRequest)
		return
	}

	messages, err := service.GetMessages(userUUID, groupUUID, limit)
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

func UpdateMessage(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	messageID, err := uuid.Parse(strings.TrimSpace(r.URL.Query().Get("id")))
	if err != nil {
		http.Error(w, "invalid message id format", http.StatusBadRequest)
		return
	}

	userID, err := currentUserIDFromContext(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var req UpdateMessageRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	message, err := service.UpdateMessage(userID, messageID, req.Content)
	if err != nil {
		if errors.Is(err, ErrNotDMMember) {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}

func DeleteMessage(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	messageID, err := uuid.Parse(strings.TrimSpace(r.URL.Query().Get("id")))
	if err != nil {
		http.Error(w, "invalid message id format", http.StatusBadRequest)
		return
	}

	userID, err := currentUserIDFromContext(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if err := service.DeleteMessage(userID, messageID); err != nil {
		if errors.Is(err, ErrNotDMMember) {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
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

	userUUID, err := uuid.Parse(strings.TrimSpace(userID))
	if err != nil {
		http.Error(w, "invalid token user ID", http.StatusUnauthorized)
		return
	}
	groupUUID, err := uuid.Parse(strings.TrimSpace(groupID))
	if err != nil {
		http.Error(w, "invalid groupId format", http.StatusBadRequest)
		return
	}

	if !service.IsUserInGroup(userUUID, groupUUID) {
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

		message, err := service.SendMessage(userUUID, SendMessageRequest{
			GroupID: groupID,
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
