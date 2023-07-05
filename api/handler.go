package api

import (
	"encoding/json"
	"github.com/Rehtang/ozonTest/storage"
	"net/http"

	"github.com/Rehtang/ozonTest/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	storage storage.URLStorage
}

func NewHandler(storage storage.URLStorage) *Handler {
	return &Handler{
		storage: storage,
	}
}

func (h *Handler) CreateShortURL(w http.ResponseWriter, r *http.Request) {
	var req struct {
		URL string `json:"url"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "failed to decode request body", http.StatusBadRequest)
		return
	}

	if !utils.ValidateURL(req.URL) {
		http.Error(w, "invalid URL", http.StatusBadRequest)
		return
	}

	shortURL := utils.GenerateShortURL()

	err = h.storage.SaveURL(shortURL, req.URL)
	if err != nil {
		http.Error(w, "failed to save URL", http.StatusInternalServerError)
		return
	}

	resp := struct {
		ShortURL string `json:"short_url"`
	}{
		ShortURL: shortURL,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, "failed to send response", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetOriginalURL(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	shortURL := params["shortURL"]

	originalURL, found, err := h.storage.GetURL(shortURL)
	if err != nil {
		http.Error(w, "failed to get URL", http.StatusInternalServerError)
		return
	}

	if !found {
		http.Error(w, "short URL not found", http.StatusNotFound)
		return
	}

	resp := struct {
		OriginalURL string `json:"original_url"`
	}{
		OriginalURL: originalURL,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, "failed to send response", http.StatusInternalServerError)
		return
	}
}
