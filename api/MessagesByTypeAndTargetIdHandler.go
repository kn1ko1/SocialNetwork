package api

import (
	"encoding/json"
	"log"
	"net/http"
	"socialnetwork/repo"
	"socialnetwork/utils"
	"strconv"
	"strings"
)

// Allowed methods: GET, PUT, DELETE

type MessagesByTypeAndTargetIdHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewMessagesByTypeAndTargetIdHandler(r repo.IRepository) *MessagesByTypeAndTargetIdHandler {
	return &MessagesByTypeAndTargetIdHandler{Repo: r}
}

// A MessagesHandler instance implements the ServeHTTP interface, and thus
// itself becomes an HTTPHandler
func (h *MessagesByTypeAndTargetIdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Switch on the Request method, call the correct subroutine...
	switch r.Method {
	case http.MethodGet:
		h.get(w, r)
		return
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *MessagesByTypeAndTargetIdHandler) get(w http.ResponseWriter, r *http.Request) {
	fields := strings.Split(r.URL.Path, "/")
	messageType := fields[len(fields)-3]
	targetIdStr := fields[len(fields)-2]
	targetId, err := strconv.Atoi(targetIdStr)
	if err != nil {
		utils.HandleError("Invalid message ID. ", err)
		http.Error(w, "internal server errror", http.StatusInternalServerError)
		return
	}
	log.Println("[api/MessagesByTypeAndTargetIdHandler]:", messageType, targetId)
	messages, err := h.Repo.GetMessagesByMessageTypeandTargetId(messageType, targetId)
	if err != nil {
		utils.HandleError("Failed to get message in GetMessagesByTypeAndTargetIdHandler. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(messages)
	if err != nil {
		utils.HandleError("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
