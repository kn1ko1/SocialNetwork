package api

import (
	"log"
	"net/http"
	"socialnetwork/repo"
	"strconv"
)

// Allowed methods: GET, PUT, DELETE

type EventUsersByEventIdAnduserIdHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewEventUserByEventIdAndUserIdHandler(r repo.IRepository) *EventUsersByEventIdAnduserIdHandler {
	return &EventUsersByEventIdAnduserIdHandler{Repo: r}
}

// A EventsHandler instance implements the ServeHTTP interface, and thus
// itself becomes an HTTPHandler
func (h *EventUsersByEventIdAnduserIdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Switch on the Request method, call the correct subroutine...
	switch r.Method {
	// case http.MethodGet:
	// 	h.get(w, r)
	// 	return
	case http.MethodDelete:
		h.delete(w, r)
		return
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
func (h *EventUsersByEventIdAnduserIdHandler) delete(w http.ResponseWriter, r *http.Request) {

	// figure out userId
	queryParams := r.URL.Query()
	eventIdString := queryParams.Get("eventId")
	eventId, eventIdErr := strconv.Atoi(eventIdString)
	if eventIdErr != nil {
		log.Println("Problem with AtoI userId. ", eventIdErr)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	userIdString := queryParams.Get("userId")
	userId, userIdErr := strconv.Atoi(userIdString)
	if userIdErr != nil {
		log.Println("Problem with AtoI userId. ", userIdErr)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	log.Println("Received delete request for eventId", eventId, ", userId", userId)

	err := h.Repo.DeleteEventUserByEventIdAndUserId(eventId, userId)
	if err != nil {
		log.Println("Failed to delete Events. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("posts were deleted"))
}

// func (h *EventUsersByEventIdAnduserIdHandler) get(w http.ResponseWriter, r *http.Request) {
// 	// figure out userId
// 	queryParams := r.URL.Query()
// 	eventIdString := queryParams.Get("eventId")
// 	eventId, eventIdErr := strconv.Atoi(eventIdString)
// 	if eventIdErr != nil {
// 		log.Println("Problem with AtoI userId. ", eventIdErr)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}
// 	userIdString := queryParams.Get("userId")
// 	userId, userIdErr := strconv.Atoi(userIdString)
// 	if userIdErr != nil {
// 		log.Println("Problem with AtoI userId. ", userIdErr)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}
// 	log.Println("Received delete request for eventId", eventId, ", userId", userId)

// 	eventUser, err := h.Repo.GetEventUserByEventIdanduserId(eventId, userId)
// 	if err != nil {
// 		log.Println("Failed to get posts in GetEventsByGroupId. ", err)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}

// 	err = json.NewEncoder(w).Encode(eventUser)
// 	if err != nil {
// 		log.Println("Failed to encode and write JSON response. ", err)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte("Here are your posts"))
// }
