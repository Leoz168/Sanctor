package community

import (
	"encoding/json"
	"net/http"

	"sanctor/internal/database"
	"sanctor/internal/pubsub"

	"github.com/google/uuid"
)

// Initialize repository, service, and messaging (defaults to in-memory)
var (
	repo      Repository = NewRepository()
	service              = NewService(repo)
	ps                   = pubsub.NewPubSub()
	messaging            = NewMessaging(ps, service)
)

// InitWithDatabase initializes the community module with a database connection
func InitWithDatabase(db *database.DB) {
	repo = NewPostgresRepository(db)
	service = NewService(repo)
	messaging = NewMessaging(ps, service)
}

// GetCommunities returns all communities
func GetCommunities(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	communities, err := service.GetAllCommunities()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(communities)
}

// GetCommunity returns a single community by ID
func GetCommunity(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Community ID is required", http.StatusBadRequest)
		return
	}

	communityID, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "Invalid community ID format", http.StatusBadRequest)
		return
	}

	community, err := service.GetCommunityWithMembers(communityID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(community)
}

// CreateCommunity creates a new community
func CreateCommunity(w http.ResponseWriter, r *http.Request) {
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

	var req CreateCommunityRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	community, err := service.CreateCommunity(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(community)
}

// UpdateCommunity updates an existing community
func UpdateCommunity(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, "Community ID is required", http.StatusBadRequest)
		return
	}

	var req UpdateCommunityRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	communityID, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "Invalid community ID format", http.StatusBadRequest)
		return
	}

	community, err := service.UpdateCommunity(communityID, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Notify community members of the update
	messaging.NotifyCommunityUpdated(community)

	json.NewEncoder(w).Encode(community)
}

// DeleteCommunity deletes a community
func DeleteCommunity(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, "Community ID is required", http.StatusBadRequest)
		return
	}

	communityID, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "Invalid community ID format", http.StatusBadRequest)
		return
	}

	if err := service.DeleteCommunity(communityID); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Notify that community was deleted
	messaging.NotifyCommunityDeleted(id)

	w.WriteHeader(http.StatusNoContent)
}

// AddUserToCommunity adds a user to a community
func AddUserToCommunity(w http.ResponseWriter, r *http.Request) {
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

	var req AddUserToCommunityRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := service.AddUserToCommunity(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Notify community members
	messaging.NotifyUserJoined(req.CommunityID, req.UserID)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User added to community successfully"})
}

// RemoveUserFromCommunity removes a user from a community
func RemoveUserFromCommunity(w http.ResponseWriter, r *http.Request) {
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
	communityID := r.URL.Query().Get("groupId")

	if userID == "" || communityID == "" {
		http.Error(w, "userId and groupId are required", http.StatusBadRequest)
		return
	}

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		http.Error(w, "Invalid user ID format", http.StatusBadRequest)
		return
	}
	communityUUID, err := uuid.Parse(communityID)
	if err != nil {
		http.Error(w, "Invalid groupId format", http.StatusBadRequest)
		return
	}

	if err := service.RemoveUserFromCommunity(userUUID, communityUUID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Notify community members
	messaging.NotifyUserLeft(communityID, userID)

	w.WriteHeader(http.StatusNoContent)
}

// GetCommunityMembers returns all members of a community
func GetCommunityMembers(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	communityID := r.URL.Query().Get("groupId")
	if communityID == "" {
		http.Error(w, "groupId is required", http.StatusBadRequest)
		return
	}

	communityUUID, err := uuid.Parse(communityID)
	if err != nil {
		http.Error(w, "Invalid groupId format", http.StatusBadRequest)
		return
	}

	members, err := service.GetCommunityMembers(communityUUID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(members)
}

// GetUserCommunities returns all communities a user belongs to
func GetUserCommunities(w http.ResponseWriter, r *http.Request) {
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

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		http.Error(w, "Invalid user ID format", http.StatusBadRequest)
		return
	}

	communities, err := service.GetUserCommunities(userUUID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(communities)
}

// SendCommunityMessage sends a message to a community
func SendCommunityMessage(w http.ResponseWriter, r *http.Request) {
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
		CommunityID string `json:"groupId"`
		UserID  string `json:"userId"`
		Content string `json:"content"`
		Type    string `json:"type,omitempty"` // defaults to "text"
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.CommunityID == "" || req.UserID == "" || req.Content == "" {
		http.Error(w, "groupId, userId, and content are required", http.StatusBadRequest)
		return
	}

	// Default to text type
	msgType := req.Type
	if msgType == "" {
		msgType = "text"
	}

	msg := &CommunityMessage{
		ID:          uuid.New().String(),
		CommunityID: req.CommunityID,
		UserID:      req.UserID,
		Content:     req.Content,
		Type:        msgType,
	}

	if err := messaging.SendCommunityMessage(msg); err != nil {
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
