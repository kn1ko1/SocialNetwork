package auth

import (
	"encoding/json"
	"log"
	"net/http"
	"socialnetwork/repo"
	"socialnetwork/transport"
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
	var loginInfo transport.LoginInfo

	cookie, err := r.Cookie(cookieName)
	if err == nil {
		log.Println("[auth/LoginHandler] here")
		_, exists := sessionMap[cookie.Value]
		if exists {
			utils.HandleError("Login failed - user already logged in:", err)
			http.Error(w, "user already logged in", http.StatusBadRequest)
			return
		}
	}

	json.NewDecoder(r.Body).Decode(&loginInfo)

	user, err := h.Repo.GetUserByUsernameOrEmail(loginInfo.UsernameOrEmail)
	if err != nil {
		utils.HandleError("Failed to retrieve user", err)
		http.Error(w, "user with specified username or email does not exist", http.StatusUnauthorized)
		return
	}

	log.Println("[auth/LoginHandler] User: ", user)
	log.Println("[auth/LoginHandler] Email: ", loginInfo.UsernameOrEmail)
	log.Println("[auth/LoginHandler] EncryptedPassword: ", user.EncryptedPassword, ". password: ", loginInfo.Password)
	err = bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(loginInfo.Password))
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
