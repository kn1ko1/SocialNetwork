package api

import (
	"encoding/json"
	"net/http"
	"socialnetwork/auth"
	"socialnetwork/models"
	"socialnetwork/repo"
	"socialnetwork/transport"
	"socialnetwork/utils"
)

type PostsAlmostPrivateWithCommentsHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewPostsAlmostPrivateWithCommentsHandler(r repo.IRepository) *PostsAlmostPrivateWithCommentsHandler {
	return &PostsAlmostPrivateWithCommentsHandler{Repo: r}
}

// A PostsHandler instance implements the ServeHTTP interface, and thus
// itself becomes an HTTPHandler
func (h *PostsAlmostPrivateWithCommentsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		h.get(w, r)
		return
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *PostsAlmostPrivateWithCommentsHandler) get(w http.ResponseWriter, r *http.Request) {
	user, err := auth.AuthenticateRequest(r)
	if err != nil {
		utils.HandleError("Error verifying cookie", err)
		http.Redirect(w, r, "auth/login", http.StatusSeeOther)
		return
	}
	privatePosts, err := h.Repo.GetPostsAlmostPrivateForUserId(user.UserId)
	if err != nil {
		utils.HandleError("Error getting private posts in GetPostsPrivateForUserId.", err)
		return
	}

	var PrivatePostsWithComments []transport.PostWithComments
	userCache := make(map[int]models.User)

	for _, post := range privatePosts {
		// Fetch and cache the post author's user details
		user, exists := userCache[post.UserId]
		if !exists {
			user, err = h.Repo.GetUserById(post.UserId)
			if err != nil {
				utils.HandleError("Failed to get user in GetPostsAlmostPrivateWithCommentsHandler.", err)
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
			utils.HandleError("Failed to get comments in GetPostsAlmostPrivateWithCommentsHandler.", err)
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
					utils.HandleError("Failed to get user in GetPostsAlmostPrivateWithCommentsHandler.", err)
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
		PrivatePostsWithComments = append(PrivatePostsWithComments, postWithComments)
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(PrivatePostsWithComments)
	if err != nil {
		utils.HandleError("Failed to encode and write JSON response.", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
