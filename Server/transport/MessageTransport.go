package transport

type MessageTransport struct {
	MessageId      int    `json:"messageId"`
	Body           string `json:"body"`
	CreatedAt      int64  `json:"createdAt"`
	MessageType    string `json:"messageType"`
	SenderUsername string `json:"senderUsername"`
	TargetId       int    `json:"targetId"`
	UpdatedAt      int64  `json:"updatedAt"`
}
