package picture

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"sanctor/internal/config"
	"sanctor/internal/database"

	"github.com/google/uuid"
)

var service *Service

// InitWithDatabase initializes the picture module with database and Supabase Storage clients.
func InitWithDatabase(db *database.DB, cfg config.SupabaseConfig) {
	repo := NewGormRepository(db)
	storage, err := NewSupabaseStorageClient(cfg.URL, cfg.ServiceRoleKey, cfg.StorageBucket)
	if err != nil {
		log.Printf("picture storage not fully configured: %v", err)
		service = NewService(repo, nil)
		return
	}
	service = NewService(repo, storage)
}

// GetPictures returns all pictures for a post.
func GetPictures(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	pictureService, ok := requireService(w)
	if !ok {
		return
	}

	postID, err := uuid.Parse(r.URL.Query().Get("postId"))
	if err != nil {
		http.Error(w, "invalid post ID format", http.StatusBadRequest)
		return
	}

	pictures, err := pictureService.GetPicturesByPost(postID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(pictures)
}

// UploadPicture uploads an image to Supabase Storage and links it to a post.
func UploadPicture(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	pictureService, ok := requireService(w)
	if !ok {
		return
	}

	postID, err := uuid.Parse(r.URL.Query().Get("postId"))
	if err != nil {
		http.Error(w, "invalid post ID format", http.StatusBadRequest)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, MaxUploadBytes)
	if err := r.ParseMultipartForm(MaxUploadBytes); err != nil {
		http.Error(w, "image upload exceeds 10MB limit or is not multipart/form-data", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "image file is required", http.StatusBadRequest)
		return
	}
	defer file.Close()

	contentType := header.Header.Get("Content-Type")
	if contentType == "" {
		buffer := make([]byte, 512)
		n, readErr := file.Read(buffer)
		if readErr != nil && readErr != io.EOF {
			http.Error(w, "failed to read image", http.StatusBadRequest)
			return
		}
		contentType = http.DetectContentType(buffer[:n])
		if _, err := file.Seek(0, io.SeekStart); err != nil {
			http.Error(w, "failed to read image", http.StatusBadRequest)
			return
		}
	}

	order := 0
	if rawOrder := r.FormValue("order"); rawOrder != "" {
		orderValue, err := strconv.Atoi(rawOrder)
		if err != nil {
			http.Error(w, "order must be an integer", http.StatusBadRequest)
			return
		}
		order = orderValue
	}

	picture, err := pictureService.UploadPicture(r.Context(), postID, file, contentType, header.Filename, r.FormValue("caption"), order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(picture)
}

// DeletePicture deletes picture metadata and its Supabase Storage object.
func DeletePicture(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	pictureService, ok := requireService(w)
	if !ok {
		return
	}

	id, err := uuid.Parse(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "invalid picture ID format", http.StatusBadRequest)
		return
	}

	if err := pictureService.DeletePicture(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func requireService(w http.ResponseWriter) (*Service, bool) {
	if service == nil {
		http.Error(w, "picture service not initialized", http.StatusServiceUnavailable)
		return nil, false
	}
	return service, true
}

func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}
