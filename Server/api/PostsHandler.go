package api

import (
	"encoding/json"
	"log"
	"net/http"
	"socialnetwork/Server/auth"
	"socialnetwork/Server/imageProcessing"
	"socialnetwork/Server/models"
	"socialnetwork/Server/repo"
	"socialnetwork/utils"
	"strconv"
	"time"
)

// Endpoint: /api/posts
// Allowed methods: POST

type PostsHandler struct {
	Repo repo.IRepository
}

// Constructor with dependency injection of a repo implementation
func NewPostsHandler(r repo.IRepository) *PostsHandler {
	return &PostsHandler{Repo: r}
}

// A PostsHandler instance implements the ServeHTTP interface, and thus
// itself becomes an HTTPHandler
func (h *PostsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// http.Error(w, "unauthorized", http.StatusUnauthorized)
	switch r.Method {
	case http.MethodPost:
		h.post(w, r)
		return
	// case http.MethodGet:
	// 	h.get(w, r)
	// 	return
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *PostsHandler) post(w http.ResponseWriter, r *http.Request) {
	user, err := auth.AuthenticateRequest(r)
	if err != nil {
		utils.HandleError("Error verifying cookie", err)
		http.Redirect(w, r, "auth/login", http.StatusSeeOther)
		return
	}
	ctime := time.Now().UTC().UnixMilli()
	// Parse form data
	err = r.ParseMultipartForm(10 << 20)
	if err != nil {
		utils.HandleError("Failed to parse form data:", err)
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	// Extract form fields
	body := r.FormValue("body")
	groupIDStr := r.FormValue("groupId")
	groupID, _ := strconv.Atoi(groupIDStr)

	imageURL := ""
	privacy := r.FormValue("privacy")
	userId := user.UserId

	log.Println("[api/PostsHandler] Received Post:", body)
	log.Println("[api/PostsHandler] GroupId:", groupID)
	log.Println("[api/PostsHandler] Privacy:", privacy)
	log.Println("[api/PostsHandler] UserId:", userId)

	// Handle file upload
	file, fileHeader, _ := r.FormFile("image")
	if file != nil {

		defer file.Close()
		imageURL, err = imageProcessing.ImageProcessing(w, r, file, *fileHeader)
		if err != nil {
			utils.HandleError("Error with ImageHandler", err)
			// http.Error(w, "Failed to process image", http.StatusInternalServerError)
			return
		}
		log.Println("[api/PostsHandler] Image Stored at:", imageURL)
	}

	// Create Post object
	post := models.Post{
		Body:      body,
		CreatedAt: ctime,
		GroupId:   groupID,
		Privacy:   privacy,
		UpdatedAt: ctime,
		ImageURL:  imageURL,
		UserId:    userId,
	}

	// Validate the post
	if validationErr := post.Validate(); validationErr != nil {
		utils.HandleError("Validation failed:", validationErr)
		http.Error(w, "Validation failed", http.StatusBadRequest)
		return
	}

	// Create post in the repository
	result, createErr := h.Repo.CreatePost(post)
	if createErr != nil {
		utils.HandleError("Failed to create post in the repository:", createErr)
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
		return
	}

	// Creates postUser if post is set to "almost private"
	if result.Privacy == "almost private" {
		almostPrivatePostUsersString := r.FormValue("almostPrivatePostUsers")
		// Parse the JSON string into a slice of strings
		var userIdStrings []string
		err := json.Unmarshal([]byte(almostPrivatePostUsersString), &userIdStrings)
		if err != nil {
			utils.HandleError("Failed to unmarshal almost private user Ids:", err)
			http.Error(w, "Failed to unmarshal almost private user Ids", http.StatusInternalServerError)
		}
		// Convert each string element to an integer
		var userIds []int
		for _, str := range userIdStrings {
			userId, err := strconv.Atoi(str)
			if err != nil {
				utils.HandleError("Failed to covert almost private user Id string to int:", err)
				http.Error(w, "Failed to covert almost private user Id string to int", http.StatusInternalServerError)
			}
			userIds = append(userIds, userId)
		}

		// allows you to see your own almost private posts
		postUser := models.PostUser{
			CreatedAt: ctime,
			PostId:    result.PostId,
			UpdatedAt: ctime,
			UserId:    user.UserId,
		}
		h.Repo.CreatePostUser(postUser)
		// allows your chosen followees to see your almost private post
		for i := 0; i < len(userIds); i++ {
			postUser := models.PostUser{
				CreatedAt: ctime,
				PostId:    result.PostId,
				UpdatedAt: ctime,
				UserId:    userIds[i],
			}
			h.Repo.CreatePostUser(postUser)
		}

	}

	// Encode and write the response
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		utils.HandleError("Failed to encode and write JSON response. ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// func (h *PostsHandler) get(w http.ResponseWriter, r *http.Request) {

// 	allPosts, err := h.Repo.GetAllPosts()
// 	if err != nil {
// 		utils.HandleError("Failed to get posts in PostHandler. ", err)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}

// 	err = json.NewEncoder(w).Encode(allPosts)
// 	if err != nil {
// 		utils.HandleError("Failed to encode and write JSON response. ", err)
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}
// }
