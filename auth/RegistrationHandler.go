package auth

import (
	"fmt"
	"log"
	"net/http"
	"socialnetwork/models"
	"socialnetwork/repo"
	"socialnetwork/utils"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type RegistrationHandler struct {
	Repo repo.IRepository
}

func NewRegistrationHandler(r repo.IRepository) *RegistrationHandler {
	return &RegistrationHandler{Repo: r}
}

func (h *RegistrationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.post(w, r)
		return
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *RegistrationHandler) post(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(cookieName)
	if err == nil {
		_, exists := sessionMap[cookie.Value]
		if exists {
			utils.HandleError("Login failed - user already logged in:", err)
			http.Error(w, "user already logged in", http.StatusBadRequest)
			return
		}
	}
	contentType := r.Header.Get("Content-Type")
	var user models.User
	switch contentType {
	case "application/x-www-form-urlencoded":
		err := r.ParseForm()
		if err != nil {
			utils.HandleError("Failed to parse form:", err)
			http.Error(w, "Failed to parse form", http.StatusInternalServerError)
			return
		}
		ctime := time.Now().UTC().UnixMilli()
		user.Bio = r.PostFormValue("bio")
		user.CreatedAt = ctime
		t := fmt.Sprintf("%s%s", r.PostFormValue("dob"), "T00:00:00Z")
		dobtime, err := time.Parse(time.RFC3339, t)
		if err != nil {
			utils.HandleError("Failed to parse date-time data", err)
			http.Error(w, "Failed to parse date-time", http.StatusInternalServerError)
			return
		}
		user.DOB = dobtime.UTC().UnixMilli()
		user.Email = r.PostFormValue("email")
		password := r.PostFormValue("password")
		encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			utils.HandleError("Error with password encryption", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		user.EncryptedPassword = string(encryptedPassword)
		user.FirstName = r.PostFormValue("first-name")
		user.ImageURL = ""
		user.IsPublic = true
		user.LastName = r.PostFormValue("last-name")
		user.UpdatedAt = ctime
		user.Username = r.PostFormValue("username")
	}
	err = user.Validate()

	if err != nil {
		utils.HandleError("User invalid", err)
		http.Error(w, "validation failed for user registration", http.StatusBadRequest)
		return
	}
	log.Println("Received user:", user)
	// _, err = sqlite.CreateUser(db, user)
	if err != nil {
		utils.HandleError("Unable to register a new user in AddUserHandler", err)
		http.Error(w, "Unable to register a new user", http.StatusBadRequest)
		return
	}
	user, err = h.Repo.CreateUser(user)
	if err != nil {
		utils.HandleError("Unable to register a new user in AddUserHandler", err)
		http.Error(w, "Unable to register a new user", http.StatusBadRequest)
		return
	}
	cookieValue = GenerateNewUUID()
	sessionMap[cookieValue] = &user
	followers, err := h.Repo.GetUserUsersBySubjectId(user.UserId)
	if err != nil {
		utils.HandleError("Failed to retrieve followers", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	var followerIds []int
	for _, f := range followers {
		followerIds = append(followerIds, f.FollowerId)
	}
	followersMap[user.UserId] = followerIds
	following, err := h.Repo.GetUserUsersByFollowerId(user.UserId)
	if err != nil {
		utils.HandleError("Failed to retrieve following", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	var followingIds []int
	for _, f := range following {
		followingIds = append(followingIds, f.SubjectId)
	}
	followingMap[user.UserId] = followingIds
	groups, err := h.Repo.GetGroupUsersByUserId(user.UserId)
	if err != nil {
		utils.HandleError("Failed to retrieve groups", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	var groupIds []int
	for _, g := range groups {
		groupIds = append(groupIds, g.GroupId)
	}
	groupsMap[user.UserId] = groupIds
	cookie = &http.Cookie{
		Name:     cookieName,
		Value:    cookieValue,
		Path:     "/",
		Expires:  time.Now().Add(timeout),
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
