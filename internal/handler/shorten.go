package handler

import (
	"fmt"
	"github.com/mpioski/goshort/internal/service"
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

var tpl = template.Must(template.ParseFiles("templates/index.html"))

type TemplateData struct {
	ShortURL string
}

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
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad form", http.StatusBadRequest)
		return
	}
	originalURL := r.FormValue("url")
	if originalURL == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	shortCode := generateShortCode(6)
	service.SetURL(shortCode, originalURL)

	data := TemplateData{
		ShortURL: fmt.Sprintf("http://localhost:8080/%s", shortCode),
	}

	tpl.ExecuteTemplate(w, "index", data)

}

func generateShortCode(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seed.Intn(charsetLength)]
	}
	return string(b)
}
