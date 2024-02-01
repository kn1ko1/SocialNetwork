package api

import (
	"encoding/json"
	"log"
	"net/http"
	"socialnetwork/models"
	"socialnetwork/repo"
	"socialnetwork/utils"
	"strconv"
)

// Endpoint: /api/groupUsers
// Allowed methods: GET, PUT, DELETE

type GroupUsersHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewGroupUsersHandler(r repo.IRepository) *GroupUsersHandler {
	return &GroupUsersHandler{Repo: r}
}

func (h *GroupUsersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

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

func (h *GroupUsersHandler) get(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	MessageIdString := queryParams.Get("groupUsersId")
	MessageId, MessageIdErr := strconv.Atoi(MessageIdString)
	if MessageIdErr != nil {
		utils.HandleError("Problem with AtoI MessageId. ", MessageIdErr)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	userMessages, err := h.Repo.GetMessageById(MessageId)
	if err != nil {
		utils.HandleError("Failed to get posts in GetMessageByIdHandler. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(userMessages)
	if err != nil {
		utils.HandleError("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Here are your posts"))
}

func (h *GroupUsersHandler) put(w http.ResponseWriter, r *http.Request) {

	var Message models.Message
	err := json.NewDecoder(r.Body).Decode(&Message)
	if err != nil {
		utils.HandleError("Failed to decode request body:", err)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	log.Println("received Message to update:", Message.Body)

	// Validate the Message
	if validationErr := Message.Validate(); validationErr != nil {
		utils.HandleError("Validation failed:", validationErr)
		http.Error(w, "Validation failed", http.StatusBadRequest)
		return
	}

	// Update post in the repository
	result, createErr := h.Repo.UpdateMessage(Message)
	if createErr != nil {
		utils.HandleError("Failed to update post in the repository:", createErr)
		http.Error(w, "Failed to update post", http.StatusInternalServerError)
		return
	}

	// Encode and write the response
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		utils.HandleError("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// Correct HTTP header for a newly created resource:
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Message updated successfully!"))
}

func (h *GroupUsersHandler) delete(w http.ResponseWriter, r *http.Request) {

	// look at penultimate id for userId

	// figure out userID
	queryParams := r.URL.Query()
	userIDString := queryParams.Get("userID")
	userID, userIDErr := strconv.Atoi(userIDString)
	if userIDErr != nil {
		utils.HandleError("Problem with AtoI userID. ", userIDErr)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	log.Println("Received delete request for userID:", userID)

	// example postId for testing
	// postId := 1

	err := h.Repo.DeleteMessageById(userID)
	if err != nil {
		utils.HandleError("Failed to delete Messages. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("posts were deleted"))
}
