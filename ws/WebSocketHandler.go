package ws

import (
	"log"
	"net/http"
	"socialnetwork/Server/auth"
	"socialnetwork/Server/repo"
	"socialnetwork/utils"

	"github.com/gorilla/websocket"
)

const (
	bufferSize = 8192
)

// WebSocketHandler handles WebSocket connections.
type WebSocketHandler struct {
	Repo repo.IRepository // Repo is an interface for interacting with the database.
}

// NewWebSocketHandler creates a new instance of WebSocketHandler.
func NewWebSocketHandler(r repo.IRepository) *WebSocketHandler {
	return &WebSocketHandler{Repo: r}
}

// ServeHTTP serves HTTP requests.
func (h *WebSocketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.get(w, r) // If the request method is GET, call the get method.
		return
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

// get handles WebSocket upgrade and authentication.
func (h *WebSocketHandler) get(w http.ResponseWriter, r *http.Request) {
	// Authenticate the user making the WebSocket request.
	c, _ := r.Cookie("SessionID")
	log.Println("[ws/WebsocketHandler], c.Value:", c.Value)
	user, err := auth.AuthenticateRequest(r)
	if err != nil {
		utils.HandleError("Error verifying cookie in WebSocket Handler: ", err)
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	// Upgrade the HTTP connection to a WebSocket connection.
	upgrader := websocket.Upgrader{ReadBufferSize: bufferSize, WriteBufferSize: bufferSize}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		utils.HandleError("Error upgrading conn in WebSocketHandler", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	// Create a new WebSocket client and add it to the appropriate socket group.
	client := NewClient(conn, user, h.Repo)
	client.SocketGroups[0] = socketGroupManager.SocketGroups[0] // Assuming there is a global socketGroupManager

	// Add the client to the default socket group.
	socketGroupManager.SocketGroups[0].Enter <- client

	// Get all groups associated with the user and put the client in those groups.
	groupUsers, err := h.Repo.GetGroupUsersByUserId(user.UserId)
	if err != nil {
		utils.HandleError("Error in GetGroupUsersByUserId in WebSocketHandler", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	for _, item := range groupUsers {
		// Check if the socket group exists, and create it if it doesn't.
		_, exists := socketGroupManager.SocketGroups[item.GroupId]
		if !exists {
			socketGroupManager.SocketGroups[item.GroupId] = NewSocketGroup(item.GroupId)
			go socketGroupManager.SocketGroups[item.GroupId].Run() // Run the socket group in a separate goroutine.
		}
		// Add the client to the socket group.
		client.SocketGroups[item.GroupId] = socketGroupManager.SocketGroups[item.GroupId]
		socketGroupManager.SocketGroups[item.GroupId].Enter <- client
	}

	// Start a goroutine to handle incoming messages from the client.
	go client.Receive()
}
