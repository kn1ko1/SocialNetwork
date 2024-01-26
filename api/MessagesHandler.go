package api

import (
	"encoding/json"
	"log"
	"net/http"
	"socialnetwork/models"
	"socialnetwork/repo"
)

// Endpoint: /api/messages
// Allowed methods: POST

type MessagesHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewMessagesHandler(r repo.IRepository) *MessagesHandler {
	return &MessagesHandler{Repo: r}
}

func (h *MessagesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodPost:
		h.post(w, r)
		return

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *MessagesHandler) post(w http.ResponseWriter, r *http.Request) {

	var message models.Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		log.Println("Failed to decode request body:", err)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	log.Println("Received message:", message.Body)

	// Validate the event
	if validationErr := message.Validate(); validationErr != nil {
		log.Println("Validation failed:", validationErr)
		http.Error(w, "Validation failed", http.StatusBadRequest)
		return
	}

	// Create event in the repository
	result, createErr := h.Repo.CreateMessage(message)
	if createErr != nil {
		log.Println("Failed to create message in the repository:", createErr)
		http.Error(w, "Failed to create message", http.StatusInternalServerError)
		return
	}

	// Encode and write the response
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		log.Println("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// Correct HTTP header for a newly created resource:
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Message created successfully!"))
}
