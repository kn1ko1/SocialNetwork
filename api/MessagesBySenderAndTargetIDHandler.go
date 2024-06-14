package api

import (
	"encoding/json"
	"log"
	"net/http"
	"socialnetwork/repo"
	"socialnetwork/transport"
	"socialnetwork/utils"
	"strconv"
	"strings"
)

// Allowed methods: GET, PUT, DELETE

type MessagesBySenderAndTargetIDHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewMessagesBySenderAndTargetIDHandler(r repo.IRepository) *MessagesBySenderAndTargetIDHandler {
	return &MessagesBySenderAndTargetIDHandler{Repo: r}
}

func (h *MessagesBySenderAndTargetIDHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

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

func (h *MessagesBySenderAndTargetIDHandler) get(w http.ResponseWriter, r *http.Request) {
	fields := strings.Split(r.URL.Path, "/")

	senderIdStr := fields[len(fields)-1]
	senderId, err := strconv.Atoi(senderIdStr)
	if err != nil {
		utils.HandleError("Invalid senderId. ", err)
		http.Error(w, "internal server errror", http.StatusInternalServerError)
		return
	}
	targetIdStr := fields[len(fields)-3]
	targetId, err := strconv.Atoi(targetIdStr)
	if err != nil {
		utils.HandleError("Invalid targetId. ", err)
		http.Error(w, "internal server errror", http.StatusInternalServerError)
		return
	}

	log.Println("[api/MessagesBySenderAndTargetIDHandler]:", senderId, targetId)
	messages, err := h.Repo.GetMessagesBySenderAndTargetIDs(senderId, targetId)
	if err != nil {
		utils.HandleError("Failed to get message in GetMessagesBySenderAndTargetIDs. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Map the existing messages to the new type
	var transportMessages []transport.MessageTransport
	for _, message := range messages {

		// Fetch the sender's username using the senderId
		sender, err := h.Repo.GetUserById(senderId)
		if err != nil {
			utils.HandleError("Failed to get sender in GetUserById. ", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			sender.Username = "Error Fetching Username"
		}
		transportMessage := transport.MessageTransport{
			MessageId:      message.MessageId,
			Body:           message.Body,
			CreatedAt:      message.CreatedAt,
			MessageType:    message.MessageType,
			SenderUsername: sender.Username,
			TargetId:       message.TargetId,
			UpdatedAt:      message.UpdatedAt,
		}
		transportMessages = append(transportMessages, transportMessage)
	}

	err = json.NewEncoder(w).Encode(transportMessages)
	if err != nil {
		utils.HandleError("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
