package ws

type WebSocketMessage struct {
	Code int    `json:"code"`
	Body string `json:"body"`
}

// 1
type TestBody struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}

// 2
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
