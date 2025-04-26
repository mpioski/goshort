package handler

import (
	"encoding/json"
	"fmt"
	"github.com/mpioski/goshort/internal/service"
	"math/rand"
	"net/http"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const charsetLength = len(charset)

var seed = rand.New(rand.NewSource(time.Now().UnixNano()))

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	ShortURL string `json:"shortURL"`
}

func Shorten(w http.ResponseWriter, r *http.Request) {
	var request ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if request.URL == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}
	shortCode := generateShortCode(6)
	service.SetURL(shortCode, request.URL)
	resp := ShortenResponse{
		ShortURL: fmt.Sprintf("localhost:8080/%s", shortCode),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(resp)

}

func generateShortCode(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seed.Intn(charsetLength)]
	}
	return string(b)
}
