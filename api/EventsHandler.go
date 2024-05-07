package api

import (
	"encoding/json"
	"log"
	"net/http"
	"socialnetwork/auth"
	"socialnetwork/models"
	"socialnetwork/repo"
	"socialnetwork/utils"
	"strconv"
	"time"
)

// Endpoint: /api/events
// Allowed methods: GET, POST

type EventsHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewEventsHandler(r repo.IRepository) *EventsHandler {
	return &EventsHandler{Repo: r}
}

// A PostsHandler instance implements the ServeHTTP interface, and thus
// itself becomes an HTTPHandler
func (h *EventsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Switch on the Request method, call the correct subroutine...
	switch r.Method {
	case http.MethodPost:
		h.post(w, r)
		return
	// All unimplemented methods default to a "method not allowed" error
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *EventsHandler) post(w http.ResponseWriter, r *http.Request) {
	var event models.Event
	var err error
	ctime := time.Now().UTC().UnixMilli()
	event.CreatedAt = ctime

	date, err := time.Parse("2006-01-02", r.PostFormValue("dateTime"))
	if err != nil {
		utils.HandleError("Failed to parse dateTime", err)
		http.Error(w, "Failed to parse dateTime", http.StatusInternalServerError)
		return
	}
	event.DateTime = date.UTC().UnixMilli()

	event.UpdatedAt = ctime
	event.Description = r.PostFormValue("description")
	groupIdStr := r.PostFormValue("groupId")
	event.GroupId, err = strconv.Atoi(groupIdStr)
	if err != nil {
		utils.HandleError("Failed to Atoi groupIdStr", err)
		http.Error(w, "Failed to Atoi groupIdStr", http.StatusInternalServerError)
		return
	}
	event.Title = r.PostFormValue("title")
	user, err := auth.AuthenticateRequest(r)
	if err != nil {
		utils.HandleError("Error verifying cookie", err)
		http.Redirect(w, r, "auth/login", http.StatusSeeOther)
		return
	}
	event.UserId = user.UserId

	// Validate the event
	if validationErr := event.Validate(); validationErr != nil {
		utils.HandleError("Validation failed:", validationErr)
		http.Error(w, "Validation failed", http.StatusBadRequest)
		return
	}

	log.Println("Received event:", event.Title, event.Description)

	// Create event in the repository
	eventResult, createErr := h.Repo.CreateEvent(event)
	if createErr != nil {
		utils.HandleError("Failed to create event in the repository:", createErr)
		http.Error(w, "Failed to create event", http.StatusInternalServerError)
		return
	}

	newEventUser := models.EventUser{
		CreatedAt: ctime,
		EventId:   eventResult.EventId,
		IsGoing:   true,
		UpdatedAt: ctime,
		UserId:    user.UserId,
	}
	_, err = h.Repo.CreateEventUser(newEventUser)
	if err != nil {
		utils.HandleError("Problem with creatingEventUser in EventsHandler. ", err)
		http.Error(w, "Internal Server Error, Problem with creatingEventUser in EventsHandler", http.StatusInternalServerError)
		return
	}

	groupUsers, err := h.Repo.GetGroupUsersByGroupId(eventResult.GroupId)
	if err != nil {
		utils.HandleError("Problem with getting groupUsers in EventsHandler. ", err)
		http.Error(w, "Internal Server Error, Problem with getting groupUsers in EventsHandler", http.StatusInternalServerError)
		return
	}

	for i := 0; i < len(groupUsers); i++ {
		if groupUsers[i].UserId == user.UserId {
			continue
		}
		notification := models.Notification{
			CreatedAt:        ctime,
			NotificationType: "eventInvite",
			ObjectId:         eventResult.EventId,
			SenderId:         user.UserId,
			Status:           "pending",
			TargetId:         groupUsers[i].UserId,
			UpdatedAt:        ctime,
		}
		h.Repo.CreateNotification(notification)
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(eventResult)
	if err != nil {
		utils.HandleError("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
