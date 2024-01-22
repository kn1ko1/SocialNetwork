package ui

import (
	"html/template"
	"log"
	"net/http"
)

type DummyPageHandler struct {
	Template *template.Template
}

func NewDummyPageHandler() *DummyPageHandler {
	tmpl, err := template.ParseFiles(layoutPath, indexPath)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &DummyPageHandler{Template: tmpl}
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
	err := h.Template.ExecuteTemplate(w, layout, nil)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
