package blocking

import (
	"encoding/json"
	"net/http"

	"sanctor/internal/database"

	"github.com/google/uuid"
)

var (
	repo    Repository = NewRepository()
	service            = NewService(repo)
)

func InitWithDatabase(db *database.DB) {
	repo = NewPostgresRepository(db)
	service = NewService(repo)
}

func GetBlocks(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	userID := r.URL.Query().Get("userId")
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		http.Error(w, "invalid user ID format", http.StatusBadRequest)
		return
	}

	blocks, err := service.GetBlocksByBlocker(userUUID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(blocks)
}

func GetBlockedBy(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	userID := r.URL.Query().Get("userId")
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		http.Error(w, "invalid user ID format", http.StatusBadRequest)
		return
	}

	blocks, err := service.GetBlocksByBlockee(userUUID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(blocks)
}

func CheckBlock(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	blockerID := r.URL.Query().Get("blockerId")
	blockeeID := r.URL.Query().Get("blockeeId")

	blockerUUID, err := uuid.Parse(blockerID)
	if err != nil {
		http.Error(w, "invalid blocker ID format", http.StatusBadRequest)
		return
	}
	blockeeUUID, err := uuid.Parse(blockeeID)
	if err != nil {
		http.Error(w, "invalid blockee ID format", http.StatusBadRequest)
		return
	}

	isBlocked, err := service.IsBlocked(blockerUUID, blockeeUUID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(BlockStatusResponse{
		BlockerID: blockerID,
		BlockeeID: blockeeID,
		IsBlocked: isBlocked,
	})
}

func CreateBlock(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var req CreateBlockRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	block, err := service.CreateBlock(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(block)
}

func DeleteBlock(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	blockerID := r.URL.Query().Get("blockerId")
	blockeeID := r.URL.Query().Get("blockeeId")

	blockerUUID, err := uuid.Parse(blockerID)
	if err != nil {
		http.Error(w, "invalid blocker ID format", http.StatusBadRequest)
		return
	}
	blockeeUUID, err := uuid.Parse(blockeeID)
	if err != nil {
		http.Error(w, "invalid blockee ID format", http.StatusBadRequest)
		return
	}

	if err := service.DeleteBlock(blockerUUID, blockeeUUID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}
