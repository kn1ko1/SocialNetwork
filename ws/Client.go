package ws

import (
	"encoding/json"
	"fmt"
	"log"
	"socialnetwork/models"

	"github.com/gorilla/websocket"
)

const (
	GROUP_CHAT_MESSAGE = 1
	PRIVATE_MESSAGE    = 2
	CREATE_EVENT       = 3
)

type Client struct {
	ClientID     int
	Connection   *websocket.Conn
	SocketGroups map[int]*SocketGroup
	User         models.User
}

func NewClient(conn *websocket.Conn, user models.User) *Client {
	return &Client{
		ClientID:     user.UserId,
		Connection:   conn,
		SocketGroups: make(map[int]*SocketGroup),
		User:         user,
	}
}

func (c *Client) Receive() {
	for {
		var wsm WebSocketMessage
		_, p, err := c.Connection.ReadMessage()
		if err != nil {
			log.Println(err.Error())
			return
		}
		err = json.Unmarshal(p, &wsm)
		if err != nil {
			log.Println(err.Error())
			return
		}
		c.HandleMessage(wsm)
	}
}

func (c *Client) Send(v any) {
	err := c.Connection.WriteJSON(v)
	if err != nil {
		log.Println(err.Error())
		return
	}
}

func (c *Client) HandleMessage(msg WebSocketMessage) {
	fmt.Println("message is:", msg)
	switch msg.Code {
	case GROUP_CHAT_MESSAGE:
		var body GroupChatBody
		err := json.Unmarshal([]byte(msg.Body), &body)
		if err != nil {
			log.Println(err.Error())
		}
		fmt.Printf("%+v\n", body)
		fmt.Println("1 testBody")
		groupId := body.GroupID
		c.SocketGroups[groupId].Broadcast <- msg
		// store message in DB
		// do stuff
	case PRIVATE_MESSAGE:
		var body PrivateMessageBody
		err := json.Unmarshal([]byte(msg.Body), &body)
		if err != nil {
			log.Println(err.Error())
		}
		fmt.Printf("%+v\n", body)
		fmt.Println("2 person")
		c.SocketGroups[0].Broadcast <- msg
		// case CREATE_EVENT:
		// 	// User action on front-end triggers this
		// 	var body models.Event
		// 	err := json.Unmarshal([]byte(msg.Body), &body)
		// 	if err != nil {
		// 		log.Println(err.Error())
		// 	}
		// 	body.Validate()
		// 	// Create Event in DB
		// 	event, err = c.Repo.CreateEvent(body)
		// 	if err != nil {
		// 		log.Println(err.Error())
		// 	}
		// 	n := models.Notification{
		// 		NotificationType: "Create Event",
		// 		ObjectId:         body.GroupID,
		// 		SenderId:         body.SenderID,
		// 		TargetId:         body.GroupID,
		// 	}
		// case 3:
		// 	var event models.Event
		// 	err := json.Unmarshal([]byte(msg.Body), &event)
		// 	if err != nil {
		// 		log.Println(err.Error())
		// 	}
		// 	fmt.Printf("%+v\n", event)
		// 	fmt.Println("3 event")
		// 	fmt.Println(msg.Body)
	}
}
