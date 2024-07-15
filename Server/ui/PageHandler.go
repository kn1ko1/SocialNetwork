package ui

import (
	"log"
	"net/http"
)

type PageHandler struct {
}

func NewPageHandler() *PageHandler {
	return &PageHandler{}
}

func (h *PageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		cookie, err := r.Cookie("SessionID")
		if err != nil {
			log.Println(err.Error())
		} else {
			log.Println("[ui/PageHandler.go]", cookie.Value)
		}
		h.get(w)
		return
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *PageHandler) get(w http.ResponseWriter) {
	err := Template.ExecuteTemplate(w, "Index", nil)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
