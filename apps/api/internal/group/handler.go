package group

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"sanctor/internal/database"
	"sanctor/internal/pubsub"
)

// Initialize repository, service, and messaging (defaults to in-memory)
var (
	repo      Repository = NewRepository()
	service              = NewService(repo)
	ps                   = pubsub.NewPubSub()
	messaging            = NewMessaging(ps, service)
)

// InitWithDatabase initializes the group module with a database connection
func InitWithDatabase(db *database.DB) {
	repo = NewPostgresRepository(db)
	service = NewService(repo)
	messaging = NewMessaging(ps, service)
}

// GetGroups returns all groups
func GetGroups(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	groups, err := service.GetAllGroups()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(groups)
}

// GetGroup returns a single group by ID
func GetGroup(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Group ID is required", http.StatusBadRequest)
		return
	}

	group, err := service.GetGroupWithMembers(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(group)
}

// CreateGroup creates a new group
func CreateGroup(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var req CreateGroupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	group, err := service.CreateGroup(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(group)
}

// UpdateGroup updates an existing group
func UpdateGroup(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "PUT" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Group ID is required", http.StatusBadRequest)
		return
	}

	var req UpdateGroupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	group, err := service.UpdateGroup(id, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Notify group members of the update
	messaging.NotifyGroupUpdated(group)

	json.NewEncoder(w).Encode(group)
}

// DeleteGroup deletes a group
func DeleteGroup(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "DELETE" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Group ID is required", http.StatusBadRequest)
		return
	}

	if err := service.DeleteGroup(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Notify that group was deleted
	messaging.NotifyGroupDeleted(id)

	w.WriteHeader(http.StatusNoContent)
}

// AddUserToGroup adds a user to a group
func AddUserToGroup(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var req AddUserToGroupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := service.AddUserToGroup(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Notify group members
	messaging.NotifyUserJoined(req.GroupID, req.UserID)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User added to group successfully"})
}

// RemoveUserFromGroup removes a user from a group
func RemoveUserFromGroup(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "DELETE" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID := r.URL.Query().Get("userId")
	groupID := r.URL.Query().Get("groupId")

	if userID == "" || groupID == "" {
		http.Error(w, "User ID and Group ID are required", http.StatusBadRequest)
		return
	}

	if err := service.RemoveUserFromGroup(userID, groupID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Notify group members
	messaging.NotifyUserLeft(groupID, userID)

	w.WriteHeader(http.StatusNoContent)
}

// GetGroupMembers returns all members of a group
func GetGroupMembers(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	groupID := r.URL.Query().Get("groupId")
	if groupID == "" {
		http.Error(w, "Group ID is required", http.StatusBadRequest)
		return
	}

	members, err := service.GetGroupMembers(groupID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(members)
}

// GetUserGroups returns all groups a user belongs to
func GetUserGroups(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	userID := r.URL.Query().Get("userId")
	if userID == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	groups, err := service.GetUserGroups(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(groups)
}

// SendGroupMessage sends a message to a group
func SendGroupMessage(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var req struct {
		GroupID string `json:"groupId"`
		UserID  string `json:"userId"`
		Content string `json:"content"`
		Type    string `json:"type,omitempty"` // defaults to "text"
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.GroupID == "" || req.UserID == "" || req.Content == "" {
		http.Error(w, "groupId, userId, and content are required", http.StatusBadRequest)
		return
	}

	// Default to text type
	msgType := req.Type
	if msgType == "" {
		msgType = "text"
	}

	msg := &Message{
		ID:      uuid.New().String(),
		GroupID: req.GroupID,
		UserID:  req.UserID,
		Content: req.Content,
		Type:    msgType,
	}

	if err := messaging.SendMessage(msg); err != nil {
		if err == ErrNotMember {
			http.Error(w, err.Error(), http.StatusForbidden)
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(msg)
}

func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}
