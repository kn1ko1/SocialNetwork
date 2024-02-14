package ws

import (
	"log"
	"net/http"
	"socialnetwork/repo"

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
	upgrader := websocket.Upgrader{ReadBufferSize: bufferSize, WriteBufferSize: bufferSize}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err.Error())
	}
	client := NewClient(conn)
	manager.Groups[0].Enter <- client
}
