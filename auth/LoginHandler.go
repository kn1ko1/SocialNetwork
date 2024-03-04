package auth

import (
	"log"
	"net/http"
	"socialnetwork/repo"
	"socialnetwork/utils"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type LoginHandler struct {
	Repo repo.IRepository
}

func NewLoginHandler(r repo.IRepository) *LoginHandler {
	return &LoginHandler{Repo: r}
}

func (h *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.post(w, r)
		return
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *LoginHandler) post(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(cookieName)
	if err == nil {
		log.Println("here")
		_, exists := sessionMap[cookie.Value]
		if exists {
			utils.HandleError("Login failed - user already logged in:", err)
			http.Error(w, "user already logged in", http.StatusBadRequest)
			return
		}
	}
	contentType := r.Header.Get("Content-Type")
	var usernameOrEmail string
	var password string
	switch contentType {
	case "application/x-www-form-urlencoded":
		err := r.ParseForm()
		if err != nil {
			utils.HandleError("Failed to parse form:", err)
			http.Error(w, "Failed to parse form", http.StatusInternalServerError)
			return
		}
		usernameOrEmail = r.PostFormValue("username")
		password = r.PostFormValue("password")
	}

	user, err := h.Repo.GetUserByUsernameOrEmail(usernameOrEmail)
	if err != nil {
		utils.HandleError("Failed to retrieve user", err)
		http.Error(w, "user with specified username or email does not exist", http.StatusUnauthorized)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(password))
	if err != nil {
		utils.HandleError("Failed to retrieve user", err)
		http.Error(w, "invalid username or password", http.StatusUnauthorized)
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
