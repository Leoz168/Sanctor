package institution

import (
	"encoding/json"
	"net/http"

	"sanctor/internal/database"

	"github.com/google/uuid"
)

// Initialize repository and service (defaults to in-memory)
var (
	repo    Repository = NewRepository()
	service            = NewService(repo)
)

// InitWithDatabase initializes the institution module with a database connection
func InitWithDatabase(db *database.DB) {
	repo = NewPostgresRepository(db)
	service = NewService(repo)
}

// GetInstitutions returns all institutions
func GetInstitutions(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	institutions, err := service.GetAllInstitutions()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(institutions)
}

// GetInstitution returns a single institution by ID
func GetInstitution(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Institution ID is required", http.StatusBadRequest)
		return
	}

	institutionID, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "Invalid institution ID format", http.StatusBadRequest)
		return
	}
	institution, err := service.GetInstitution(institutionID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(institution)
}

// CreateInstitution creates a new institution
func CreateInstitution(w http.ResponseWriter, r *http.Request) {
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

	var req CreateInstitutionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	institution, err := service.CreateInstitution(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(institution)
}

// UpdateInstitution updates an existing institution
func UpdateInstitution(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, "Institution ID is required", http.StatusBadRequest)
		return
	}

	var req UpdateInstitutionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	institutionID, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "Invalid institution ID format", http.StatusBadRequest)
		return
	}
	institution, err := service.UpdateInstitution(institutionID, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(institution)
}

// DeleteInstitution deletes an institution
func DeleteInstitution(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, "Institution ID is required", http.StatusBadRequest)
		return
	}

	institutionID, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "Invalid institution ID format", http.StatusBadRequest)
		return
	}
	if err := service.DeleteInstitution(institutionID); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}
