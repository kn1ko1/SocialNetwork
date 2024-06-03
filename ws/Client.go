package ws

import (
	"encoding/json"
	"fmt"
	"log"
	"socialnetwork/models"
	"socialnetwork/repo"

	"github.com/gorilla/websocket"
)

const (
	GROUP_CHAT_MESSAGE = 1
	PRIVATE_MESSAGE    = 2
	CREATE_EVENT       = 3
	GROUP_REQUEST      = 4
	GROUP_INVITE       = 5
	EVENT_INVITE       = 6
)

type Client struct {
	ClientID     int
	Connection   *websocket.Conn
	SocketGroups map[int]*SocketGroup
	User         models.User
	Repo         repo.IRepository // Add a field to hold the repository instance
}

func NewClient(conn *websocket.Conn, user models.User, repo repo.IRepository) *Client {
	return &Client{
		ClientID:     user.UserId,
		Connection:   conn,
		SocketGroups: make(map[int]*SocketGroup),
		User:         user,
		Repo:         repo, // Pass the repository instance to the Client
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
	var message models.Message
	switch msg.Code {
	case GROUP_CHAT_MESSAGE:
		err := json.Unmarshal([]byte(msg.Body), &message)
		if err != nil {
			log.Println(err.Error())
			return
		}
		fmt.Printf("%+v\n", message)
		fmt.Println("1 testBody")
		groupId := message.TargetId
		group, ok := c.SocketGroups[groupId]
		if !ok {
			log.Printf("SocketGroup %d does not exist\n", groupId)
			return
		}
		group.Broadcast <- msg
		c.Repo.CreateMessage(message)
		log.Println("Group Message added to db in Client.go is:", message)
		// store message in DB
		// do stuff
	case PRIVATE_MESSAGE:
		err := json.Unmarshal([]byte(msg.Body), &message)
		if err != nil {
			log.Println(err.Error())
			return
		}
		fmt.Printf("%+v\n", message)
		fmt.Println("2 person")
		group, ok := c.SocketGroups[0]
		if !ok {
			log.Println("Private message group does not exist")
			return
		}
		group.Broadcast <- msg
		c.Repo.CreateMessage(message)
		log.Println("'prvivate' (group 0) Message added to db in Client.go is:", message)

	}
}
