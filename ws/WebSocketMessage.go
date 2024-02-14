package ws

type WebSocketMessage struct {
	Code int    `json:"code"`
	Body string `json:"body"`
}
