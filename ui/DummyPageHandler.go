package ui

import (
	"net/http"
)

type DummyPageHandler struct {
}

func NewDummyPageHandler() *DummyPageHandler {
	return &DummyPageHandler{}
}

func (h *DummyPageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.get(w, r)
		return
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *DummyPageHandler) get(w http.ResponseWriter, r *http.Request) {
	err := Template.ExecuteTemplate(w, "Index", nil)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
