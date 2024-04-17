package ws

import (
	"log"
	"net/http"
	"socialnetwork/auth"
	"socialnetwork/repo"
	"socialnetwork/utils"

	"github.com/gorilla/websocket"
)

const (
	bufferSize = 8192
)

type WebSocketHandler struct {
	Repo repo.IRepository
}

func NewWebSocketHandler(r repo.IRepository) *WebSocketHandler {
	return &WebSocketHandler{Repo: r}
}

func (h *WebSocketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.get(w, r)
		return
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (h *WebSocketHandler) get(w http.ResponseWriter, r *http.Request) {
	user, err := auth.AuthenticateRequest(r)
	if err != nil {
		utils.HandleError("Error verifying cookie in WebSocket Handler: ", err)
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	upgrader := websocket.Upgrader{ReadBufferSize: bufferSize, WriteBufferSize: bufferSize}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	client := NewClient(conn, user)
	client.SocketGroups[0] = socketGroupManager.SocketGroups[0]
	socketGroupManager.SocketGroups[0].Enter <- client
	// Get all groups and put client in those groups
	groupUsers, err := h.Repo.GetGroupUsersByUserId(user.UserId)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	for _, item := range groupUsers {
		_, exists := socketGroupManager.SocketGroups[item.GroupId]
		if !exists {
			socketGroupManager.SocketGroups[item.GroupId] = NewSocketGroup(item.GroupId)
		}
		client.SocketGroups[item.GroupId] = socketGroupManager.SocketGroups[item.GroupId]
		socketGroupManager.SocketGroups[item.GroupId].Enter <- client
	}
	go client.Receive()
}
