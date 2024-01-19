package api

import (
	"encoding/json"
	"log"
	"net/http"
	"socialnetwork/auth"
	"socialnetwork/models"
	"socialnetwork/repo"
)

// Endpoint: /api/messages
// Allowed methods: GET, POST, PUT, DELETE

type MessagesHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewMessagesHandler(r repo.IRepository) *MessagesHandler {
	return &MessagesHandler{Repo: r}
}

// A MessagesHandler instance implements the ServeHTTP interface, and thus
// itself becomes an HTTPHandler
func (h *MessagesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Get Session Cookie
	c, err := r.Cookie("Session")
	if err != nil {
		// Log Error
		log.Println(err.Error())
		// Return HTTP Status Unauthorized
		//
		// N.B. for simplicity of the example, we are simply returning
		// an HTTP error. In the actual project, probably a JSON payload.
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	// Authenticate Session Cookie - user variable discarded because user struct not used here...
	_, err = auth.AuthenticateSessionCookie(c)
	if err != nil {
		// Same error as above - maker of request is unauthorized
		log.Println(err.Error())
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	// Switch on the Request method, call the correct subroutine...
	switch r.Method {
	// HTTP GET logic
	case http.MethodPost:
		h.post(w, r)
		return
	case http.MethodGet:
		h.get(w, r)
		return
	case http.MethodPut:
		h.put(w, r)
		return
	case http.MethodDelete:
		h.delete(w, r)
		return
	// All unimplemented methods default to a "method not allowed" error
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

// Separate out HTTP methods for clean separation of concerns
// N.B. Use lowercase names, i.e. "post", "get", etc. for correct encapsulation
func (h *MessagesHandler) post(w http.ResponseWriter, r *http.Request) {

	// Enable CORS headers for this handler
	SetupCORS(&w, r)

	var message models.Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		log.Println("Failed to decode request body:", err)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	log.Println("Received message:", message.Body)

	// Example message to test function
	// message := models.Message{
	// 	Body:        "example message",
	// 	CreatedAt:   111111111,
	// 	MessageType: "example",
	// 	SenderID:    1,
	// 	TargetID:    2,
	// 	UpdatedAt:   111111}

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
	w.Write([]byte("Event created successfully!"))
}

func (h *MessagesHandler) get(w http.ResponseWriter, r *http.Request) {

	// Enable CORS headers for this handler
	SetupCORS(&w, r)

	allMessages, err := h.Repo.GetAllMessages()
	if err != nil {
		log.Println("Failed to get messages in Messages. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(allMessages)
	if err != nil {
		log.Println("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Here are your messages"))
}

func (h *MessagesHandler) put(w http.ResponseWriter, r *http.Request) {

	// Enable CORS headers for this handler
	SetupCORS(&w, r)

	var message models.Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		log.Println("Failed to decode request body:", err)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	log.Println("Received message:", message.Body)

	// Example message to test function
	// message := models.Message{
	// 	Body:        "updated example message",
	// 	CreatedAt:   111111111,
	// 	MessageType: "example",
	// 	SenderID:    1,
	// 	TargetID:    2,
	// 	UpdatedAt:   333333333333}

	// Validate the event
	if validationErr := message.Validate(); validationErr != nil {
		log.Println("Validation failed:", validationErr)
		http.Error(w, "Validation failed", http.StatusBadRequest)
		return
	}

	// Create event in the repository
	result, createErr := h.Repo.UpdateMessage(message)
	if createErr != nil {
		log.Println("Failed to update message in the repository:", createErr)
		http.Error(w, "Failed to update message", http.StatusInternalServerError)
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
	w.Write([]byte("Post updated successfully!"))
}

func (h *MessagesHandler) delete(w http.ResponseWriter, r *http.Request) {

	// Enable CORS headers for this handler
	SetupCORS(&w, r)

	// figure out eventId
	var messageId int
	err := json.NewDecoder(r.Body).Decode(&messageId)
	if err != nil {
		log.Println("Failed to decode request body:", err)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	log.Println("Received delete request for messageId:", messageId)

	// example eventId for testing
	// eventId := 1

	err = h.Repo.DeleteMessageById(messageId)
	if err != nil {
		log.Println("Failed to delete messageId. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("messageId was deleted"))
}
