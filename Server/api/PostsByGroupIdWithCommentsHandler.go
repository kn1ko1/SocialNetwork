package api

import (
	"encoding/json"
	"log"
	"net/http"
	"socialnetwork/Server/models"
	"socialnetwork/Server/repo"
	"socialnetwork/Server/transport"
	"socialnetwork/utils"
	"strconv"
	"strings"
)

// Endpoint: /api/posts/privacy/public
// Allowed methods: GET

type PostsByGroupIdWithCommentsHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewPostsByGroupIdWithCommentsHandler(r repo.IRepository) *PostsByGroupIdWithCommentsHandler {
	return &PostsByGroupIdWithCommentsHandler{Repo: r}
}

// A PostsHandler instance implements the ServeHTTP interface, and thus
// itself becomes an HTTPHandler
func (h *PostsByGroupIdWithCommentsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		h.get(w, r)
		return
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *PostsByGroupIdWithCommentsHandler) get(w http.ResponseWriter, r *http.Request) {
	fields := strings.Split(r.URL.Path, "/")
	groupIdStr := fields[len(fields)-3]

	groupId, err := strconv.Atoi(groupIdStr)
	if err != nil {
		utils.HandleError("Invalid group Id. ", err)
		http.Error(w, "internal server errror", http.StatusInternalServerError)
		return
	}
	log.Println("Received get request for group Id:", groupId)
	groupPosts, err := h.Repo.GetPostsByGroupId(groupId)
	if err != nil {
		utils.HandleError("Failed to get posts in GetPostsByGroupIdWithCommentsHandler.", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	var groupPostsWithComments []transport.PostWithComments
	userCache := make(map[int]models.User)

	for _, post := range groupPosts {
		// Fetch and cache the post author's user details
		user, exists := userCache[post.UserId]
		if !exists {
			user, err = h.Repo.GetUserById(post.UserId)
			if err != nil {
				utils.HandleError("Failed to get user in GetPostsByGroupIdWithCommentsHandler.", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			userCache[post.UserId] = user
		}

		postTransport := transport.PostTransport{
			Post: post,
			User: user,
		}

		// Fetch comments for the post
		postComments, err := h.Repo.GetCommentsByPostId(post.PostId)
		if err != nil {
			utils.HandleError("Failed to get comments in GetPostsByGroupIdWithCommentsHandler.", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		var transportComments []transport.CommentTransport
		for _, comment := range postComments {
			// Fetch and cache the comment author's user details
			user, exists := userCache[comment.UserId]
			if !exists {
				user, err := h.Repo.GetUserById(comment.UserId)
				if err != nil {
					utils.HandleError("Failed to get user in GetPostsByGroupIdWithCommentsHandler.", err)
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
					return
				}
				userCache[comment.UserId] = user
			}

			commentTransport := transport.CommentTransport{
				Comment: comment,
				User:    user,
			}
			transportComments = append(transportComments, commentTransport)
		}

		postWithComments := transport.PostWithComments{
			Post:     postTransport,
			Comments: transportComments,
		}
		groupPostsWithComments = append(groupPostsWithComments, postWithComments)
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(groupPostsWithComments)
	if err != nil {
		utils.HandleError("Failed to encode and write JSON response.", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
