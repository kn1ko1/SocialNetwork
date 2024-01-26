package api

import (
	"encoding/json"
	"log"
	"net/http"
	"socialnetwork/models"
	"socialnetwork/repo"
	"strconv"
)

// Endpoint: /api/Messa s/Message/{MessageId}
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
	messageIdString := queryParams.Get("postId")
	messageId, messageIdErr := strconv.Atoi(messageIdString)
	if messageIdErr != nil {
		log.Println("Problem with AtoI messageId. ", messageIdErr)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	userMessages, err := h.Repo.GetMessageById(messageId)
	if err != nil {
		log.Println("Failed to get posts in GetMessageByIdHandler. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(userMessages)
	if err != nil {
		log.Println("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Here are your posts"))
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

	// Example Post to test function
	// post := models.Post{
	// 	Body: "Updated Example",
	// 	CreatedAt: 111111,
	// 	UpdatedAt: 333333,
	// 	UserId: 2}

	// Validate the Message
	if validationErr := Message.Validate(); validationErr != nil {
		log.Println("Validation failed:", validationErr)
		http.Error(w, "Validation failed", http.StatusBadRequest)
		return
	}

	// Update post in the repository
	result, createErr := h.Repo.UpdateMessageById(Message)
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
	// Correct HTTP header for a newly created resource:
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Message updated successfully!"))
}

func (h *MessageByIdHandler) delete(w http.ResponseWriter, r *http.Request) {

	// look at penultimate id for userId

	// figure out userID
	queryParams := r.URL.Query()
	userIDString := queryParams.Get("userID")
	userID, userIDErr := strconv.Atoi(userIDString)
	if userIDErr != nil {
		log.Println("Problem with AtoI userID. ", userIDErr)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	log.Println("Received delete request for userID:", userID)

	// example postId for testing
	// postId := 1

	err := h.Repo.DeleteMessageById(userID)
	if err != nil {
		log.Println("Failed to delete Messages. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("posts were deleted"))
}
