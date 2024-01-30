package api

import (
	"encoding/json"
	"log"
	"net/http"
	"socialnetwork/models"
	"socialnetwork/repo"
	"strconv"
)

// Allowed methods: GET, PUT, DELETE

type MessageByIdHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewMessageByIdHandler(r repo.IRepository) *MessageByIdHandler {
	return &MessageByIdHandler{Repo: r}
}

// A MessagesHandler instance implements the ServeHTTP interface, and thus
// itself becomes an HTTPHandler
func (h *MessageByIdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Switch on the Request method, call the correct subroutine...
	switch r.Method {
	case http.MethodGet:
		h.get(w, r)
		return
	case http.MethodDelete:
		h.delete(w, r)
		return
	case http.MethodPut:
		h.put(w, r)
		return
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *MessageByIdHandler) get(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	messageIdString := queryParams.Get("messageId")
	messageId, messageIdErr := strconv.Atoi(messageIdString)
	if messageIdErr != nil {
		log.Println("Problem with AtoI messageId. ", messageIdErr)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	message, err := h.Repo.GetMessageById(messageId)
	if err != nil {
		log.Println("Failed to get message in GetMessageByIdHandler. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(message)
	if err != nil {
		log.Println("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *MessageByIdHandler) put(w http.ResponseWriter, r *http.Request) {

	var Message models.Message
	err := json.NewDecoder(r.Body).Decode(&Message)
	if err != nil {
		log.Println("Failed to decode request body:", err)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	log.Println("received Message to update:", Message.Body)

	// Validate the Message
	if validationErr := Message.Validate(); validationErr != nil {
		log.Println("Validation failed:", validationErr)
		http.Error(w, "Validation failed", http.StatusBadRequest)
		return
	}

	// Update post in the repository
	result, createErr := h.Repo.UpdateMessage(Message)
	if createErr != nil {
		log.Println("Failed to update post in the repository:", createErr)
		http.Error(w, "Failed to update post", http.StatusInternalServerError)
		return
	}

	// Encode and write the response
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		log.Println("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *MessageByIdHandler) delete(w http.ResponseWriter, r *http.Request) {

	queryParams := r.URL.Query()
	messageIdString := queryParams.Get("messageId")
	messageId, messageIdErr := strconv.Atoi(messageIdString)
	if messageIdErr != nil {
		log.Println("Problem with AtoI messageId. ", messageIdErr)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	log.Println("Received delete request for messageId:", messageId)

	err := h.Repo.DeleteMessageById(messageId)
	if err != nil {
		log.Println("Failed to delete Messages. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
