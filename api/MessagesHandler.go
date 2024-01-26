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

// A MessagesHandler instance implements the ServeHTTP interface, and thus
// itself becomes an HTTPHandler
func (h *MessagesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Switch on the Request method, call the correct subroutine...
	switch r.Method {
	// HTTP GET logic
	case http.MethodPost:
		h.post(w, r)
		return

	// All unimplemented methods default to a "method not allowed" error
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

// POST /api/messages
func (h *MessagesHandler) post(w http.ResponseWriter, r *http.Request) {

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

// // GET /api/messages?senderId=1&targetId=2
// func (h *MessagesHandler) get(w http.ResponseWriter, r *http.Request) {

// 	queryParams := r.URL.Query()
// 	senderIdString := queryParams.Get("senderId")
// 	targetIdString := queryParams.Get("targetId")
// 	senderId, senderIdErr := strconv.Atoi(senderIdString)
// 	if senderIdErr != nil {
// 		log.Println("Problem with AtoI senderId. ", senderIdErr)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}
// 	targetId, targetIdErr := strconv.Atoi(targetIdString)
// 	if targetIdErr != nil {
// 		log.Println("Problem with AtoI targetId. ", targetIdErr)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}
// 	chatMessages, err := h.Repo.GetMessagesBySenderAndTargetIDs(senderId, targetId)
// 	if err != nil {
// 		log.Println("Failed to get messages in Messages. ", err)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}

// 	err = json.NewEncoder(w).Encode(chatMessages)
// 	if err != nil {
// 		log.Println("Failed to encode and write JSON response. ", err)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("Here are your messages"))
// }

// // PUT /api/messages/{messageId}
// func (h *MessagesHandler) put(w http.ResponseWriter, r *http.Request) {

// 	var message models.Message
// 	err := json.NewDecoder(r.Body).Decode(&message)
// 	if err != nil {
// 		log.Println("Failed to decode request body:", err)
// 		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
// 		return
// 	}
// 	log.Println("Received message:", message.Body)

// 	// Example message to test function
// 	// message := models.Message{
// 	// 	Body:        "updated example message",
// 	// 	CreatedAt:   111111111,
// 	// 	MessageType: "example",
// 	// 	SenderID:    1,
// 	// 	TargetID:    2,
// 	// 	UpdatedAt:   333333333333}

// 	// Validate the event
// 	if validationErr := message.Validate(); validationErr != nil {
// 		log.Println("Validation failed:", validationErr)
// 		http.Error(w, "Validation failed", http.StatusBadRequest)
// 		return
// 	}

// 	// Create event in the repository
// 	result, createErr := h.Repo.UpdateMessage(message)
// 	if createErr != nil {
// 		log.Println("Failed to update message in the repository:", createErr)
// 		http.Error(w, "Failed to update message", http.StatusInternalServerError)
// 		return
// 	}

// 	// Encode and write the response
// 	err = json.NewEncoder(w).Encode(result)
// 	if err != nil {
// 		log.Println("Failed to encode and write JSON response. ", err)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}
// 	// Correct HTTP header for a newly created resource:
// 	w.WriteHeader(http.StatusCreated)
// 	w.Write([]byte("Post updated successfully!"))
// }

// // DELETE /api/messages/{messageId}
// func (h *MessagesHandler) delete(w http.ResponseWriter, r *http.Request) {

// 	// figure out eventId
// 	var messageId int
// 	err := json.NewDecoder(r.Body).Decode(&messageId)
// 	if err != nil {
// 		log.Println("Failed to decode request body:", err)
// 		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
// 		return
// 	}
// 	log.Println("Received delete request for messageId:", messageId)

// 	// example eventId for testing
// 	// eventId := 1

// 	err = h.Repo.DeleteMessageById(messageId)
// 	if err != nil {
// 		log.Println("Failed to delete messageId. ", err)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("messageId was deleted"))
// }
