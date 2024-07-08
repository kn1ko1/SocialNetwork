package api

import (
	"encoding/json"
	"log"
	"net/http"
	"socialnetwork/Server/repo"
	"socialnetwork/Server/transport"
	"socialnetwork/Server/utils"
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

	// Map the existing messages to the transport type
	var transportMessages []transport.MessageTransport
	for _, message := range messages {

		// Fetch the sender's username using the senderId
		sender, err := h.Repo.GetUserById(message.SenderId)
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
