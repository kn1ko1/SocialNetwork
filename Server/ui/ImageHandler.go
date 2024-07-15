package ui

import (
	"net/http"
	"os"
	"path/filepath"
)

type ImageHandler struct {
}

func NewImageHandler() *ImageHandler {
	return &ImageHandler{}
}

func (h *ImageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.get(w, r)
		return
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *ImageHandler) get(w http.ResponseWriter, r *http.Request) {
	// Construct the file path
	workDir, _ := os.Getwd()
	filePath := filepath.Join(workDir, "uploads", r.URL.Path[len("/uploads/"):])

	// Serve the file
	http.ServeFile(w, r, filePath)
}
