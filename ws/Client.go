package ws

import (
	"encoding/json"
	"fmt"
	"log"
	"socialnetwork/models"
	"socialnetwork/repo"
	"time"

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
	defer func() {
		// Remove the client from all SocketGroups when the connection is closed
		for _, group := range c.SocketGroups {
			group.Exit <- c
		}
		c.Connection.Close()
	}()
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
	fmt.Println("[ws/Client.go] message is:", msg)
	var message models.Message
	switch msg.Code {
	case GROUP_CHAT_MESSAGE:
		err := json.Unmarshal([]byte(msg.Body), &message)
		if err != nil {
			log.Println(err.Error())
			return
		}
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
		fmt.Println("2 person")
		group, ok := c.SocketGroups[0]
		if !ok {
			log.Println("Private message group does not exist")
			return
		}
		group.Broadcast <- msg
		c.Repo.CreateMessage(message)
		log.Println("'prvivate' (group 0) Message added to db in Client.go is:", message)
	case EVENT_INVITE:
		var event models.Event
		err := json.Unmarshal([]byte(msg.Body), &event)
		if err != nil {
			log.Println(err.Error())
			return
		}
		ctime := time.Now().UTC().UnixMilli()
		event.CreatedAt = ctime
		event.UpdatedAt = ctime
		returnEvent, err := c.Repo.CreateEvent(event)
		if err != nil {
			log.Println(err.Error())
			return
		}

		fmt.Println("6 EVENT_INVITE")
		fmt.Println("[ws/Client.go]] subjectId", event.UserId)

		notification := models.Notification{
			CreatedAt:        ctime,
			NotificationType: "event",
			ObjectId:         returnEvent.EventId,
			SenderId:         returnEvent.UserId,
			Status:           "pending",
			TargetId:         returnEvent.GroupId,
			UpdatedAt:        ctime,
		}
		c.Repo.CreateNotification(notification)
		log.Println("6 EVENT_INVITE Notification added to db in Client.go is:", notification)
		jsonNotification, err := json.Marshal(notification)
		if err != nil {
			log.Println("[ws/client.go]", err.Error())

		}
		returnMessage := WebSocketMessage{
			Code: 6,
			Body: string(jsonNotification),
		}

		groupId := notification.TargetId
		group, ok := c.SocketGroups[groupId]
		if !ok {
			log.Printf("SocketGroup %d does not exist\n", groupId)
			return
		}
		log.Println("[ws/Client.go] Event GroupId is:", groupId)
		group.Broadcast <- returnMessage

	}
}
