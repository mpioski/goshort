package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/mpioski/goshort/internal/service"
)

func Redirect(w http.ResponseWriter, r *http.Request) {
	shortCode := strings.Trim(r.URL.Path, "/")

	if shortCode == "" {
		http.ServeFile(w, r, "./static/index.html")
		return
	}

	originalURL, err := service.GetURL(shortCode)
	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusNotFound)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusFound)
}
