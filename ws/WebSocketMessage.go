package ws

type WebSocketMessage struct {
	Code int    `json:"code"`
	Body string `json:"body"`
}

// 1
type GroupChatBody struct {
	Message string `json:"message"`
	GroupID int    `json:"groupID"`
}

type PrivateMessageBody struct {
	Message      string `json:"message"`
	SenderUserID int    `json:"senderUserID"`
	TargetUserID int    `json:"targetUserID"`
}
