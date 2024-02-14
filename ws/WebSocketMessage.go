package ws

type WebSocketMessage struct {
	Code int    `json:"code"`
	Body string `json:"body"`
}

type TestBody struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}
