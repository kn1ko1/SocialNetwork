package api

import (
	"encoding/json"
	"net/http"
	"socialnetwork/models"
	"socialnetwork/repo"
	"socialnetwork/transport"
	"socialnetwork/utils"
)

// Endpoint: /api/posts/privacy/public
// Allowed methods: GET

type PostsPublicWithCommentsHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewPostsPublicWithCommentsHandler(r repo.IRepository) *PostsPublicWithCommentsHandler {
	return &PostsPublicWithCommentsHandler{Repo: r}
}

// A PostsHandler instance implements the ServeHTTP interface, and thus
// itself becomes an HTTPHandler
func (h *PostsPublicWithCommentsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		h.get(w)
		return
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *PostsPublicWithCommentsHandler) get(w http.ResponseWriter) {

	publicPosts, err := h.Repo.GetPostsByPrivacy("public")
	if err != nil {
		utils.HandleError("Failed to get posts in GetPostsPublicWithCommentsHandler.", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	var publicPostsWithComments []transport.PostWithComments
	userCache := make(map[int]models.User)

	for _, post := range publicPosts {
		// Fetch and cache the post author's user details
		user, exists := userCache[post.UserId]
		if !exists {
			user, err := h.Repo.GetUserById(post.UserId)
			if err != nil {
				utils.HandleError("Failed to get user in GetPostsPublicWithCommentsHandler.", err)
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
			utils.HandleError("Failed to get comments in GetPostsPublicWithCommentsHandler.", err)
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
					utils.HandleError("Failed to get user in GetPostsPublicWithCommentsHandler.", err)
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
		publicPostsWithComments = append(publicPostsWithComments, postWithComments)
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(publicPostsWithComments)
	if err != nil {
		utils.HandleError("Failed to encode and write JSON response.", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
