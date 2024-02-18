package ui

import (
	"fmt"
	"log"
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
		cookie, err := r.Cookie("SessionID")
		if err != nil {
			log.Println(err.Error())
		} else {
			fmt.Println(cookie.Value)
		}
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
