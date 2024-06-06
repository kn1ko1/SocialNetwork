package api

import (
	"encoding/json"
	"log"
	"net/http"
	"socialnetwork/models"
	"socialnetwork/repo"
	"socialnetwork/transport"
	"socialnetwork/utils"
	"strconv"
	"strings"
	"time"
)

// Allowed methods: GET, PUT, DELETE

type NotificationByIdHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewNotificationByIdHandler(r repo.IRepository) *NotificationByIdHandler {
	return &NotificationByIdHandler{Repo: r}
}

// A NotificationsHandler instance implements the ServeHTTP interface, and thus
// itself becomes an HTTPHandler
func (h *NotificationByIdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

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

func (h *NotificationByIdHandler) get(w http.ResponseWriter, r *http.Request) {
	fields := strings.Split(r.URL.Path, "/")
	notificationStr := fields[len(fields)-1]

	notificationId, err := strconv.Atoi(notificationStr)
	if err != nil {
		utils.HandleError("Invalid Group ID. ", err)
		http.Error(w, "internal server errror", http.StatusInternalServerError)
		return
	}
	log.Println("Received delete request for notificationId:", notificationId)

	notification, err := h.Repo.GetNotificationById(notificationId)
	if err != nil {
		utils.HandleError("Failed to get notification in GetNotificationByIdHandler. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(notification)
	if err != nil {
		utils.HandleError("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *NotificationByIdHandler) put(w http.ResponseWriter, r *http.Request) {

	var Notification models.Notification
	err := json.NewDecoder(r.Body).Decode(&Notification)
	if err != nil {
		utils.HandleError("Failed to decode request body:", err)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	log.Println("received Notification to update:", Notification.NotificationType)

	// // Validate the Notification
	if validationErr := Notification.Validate(); validationErr != nil {
		utils.HandleError("Validation failed:", validationErr)
		http.Error(w, "Validation failed", http.StatusBadRequest)
		return
	}

	// Update post in the repository
	result, createErr := h.Repo.UpdateNotification(Notification)
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
}

func (h *NotificationByIdHandler) delete(w http.ResponseWriter, r *http.Request) {

	fields := strings.Split(r.URL.Path, "/")
	notificationStr := fields[len(fields)-1]

	notificationId, notificationIdErr := strconv.Atoi(notificationStr)
	if notificationIdErr != nil {
		utils.HandleError("Problem with AtoI notificationId. ", notificationIdErr)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	log.Println("Received delete request for notificationId:", notificationId)

	ctime := time.Now().UTC().UnixMilli()

	var notificationResponse transport.NotificationResponse

	err := json.NewDecoder(r.Body).Decode(&notificationResponse)
	if err != nil {
		utils.HandleError("Failed to decode request body:", err)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	log.Println("[api/NotificationById] Response:", notificationResponse)

	if notificationResponse.Reply == "confirm" {
		switch notificationResponse.Notification.NotificationType {
		case "groupInvite":
			log.Println("groupInvite")
			groupUser := models.GroupUser{
				CreatedAt: ctime,
				GroupId:   notificationResponse.Notification.ObjectId,
				UpdatedAt: ctime,
				UserId:    notificationResponse.Notification.TargetId,
			}
			h.Repo.CreateGroupUser(groupUser)
		case "groupRequest":
			log.Println("groupRequest")
			groupUser := models.GroupUser{
				CreatedAt: ctime,
				GroupId:   notificationResponse.Notification.ObjectId,
				UpdatedAt: ctime,
				UserId:    notificationResponse.Notification.SenderId,
			}
			h.Repo.CreateGroupUser(groupUser)
		case "eventInvite":
			log.Println("eventInvite")
			eventUser := models.EventUser{
				CreatedAt: ctime,
				EventId:   notificationResponse.Notification.ObjectId,
				IsGoing:   true,
				UpdatedAt: ctime,
				UserId:    notificationResponse.Notification.TargetId,
			}
			h.Repo.CreateEventUser(eventUser)
		case "followRequest":
			log.Println("followRequest")
			userUser := models.UserUser{
				CreatedAt:  ctime,
				FollowerId: notificationResponse.Notification.SenderId,
				UpdatedAt:  ctime,
				SubjectId:  notificationResponse.Notification.TargetId,
			}
			h.Repo.CreateUserUser(userUser)
		}
	}

	err = h.Repo.DeleteNotificationById(notificationId)
	if err != nil {
		utils.HandleError("Failed to delete Notifications. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// Send a success response
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("Notification deleted successfully"))
}
