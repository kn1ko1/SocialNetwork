package ws

import (
	"encoding/json"
	"log"
	"socialnetwork/models"
	"socialnetwork/repo"
	"socialnetwork/utils"
	"time"

	"github.com/gorilla/websocket"
)

// Define constants for different types of WebSocket messages
const (
	GROUP_CHAT_MESSAGE = 1
	PRIVATE_MESSAGE    = 2
	FOLLOW_REQUEST     = 3
	GROUP_REQUEST      = 4
	GROUP_INVITE       = 5
	EVENT_INVITE       = 6
)

// Client represents a connected user
type Client struct {
	ClientID     int
	Connection   *websocket.Conn
	SocketGroups map[int]*SocketGroup
	User         models.User
	Repo         repo.IRepository
}

// NewClient creates a new Client instance
func NewClient(conn *websocket.Conn, user models.User, repo repo.IRepository) *Client {
	return &Client{
		ClientID:     user.UserId,
		Connection:   conn,
		SocketGroups: make(map[int]*SocketGroup),
		User:         user,
		Repo:         repo,
	}
}

// Receive listens for incoming messages from the WebSocket connection
func (c *Client) Receive() {
	defer func() {
		// Remove the client from all SocketGroups when the connection is closed
		for _, group := range c.SocketGroups {
			group.Exit <- c
		}
		// Close the WebSocket connection
		c.Connection.Close()
	}()
	for {
		var wsm WebSocketMessage
		// Read a message from the WebSocket connection
		_, p, err := c.Connection.ReadMessage()
		if err != nil {
			log.Println(err.Error())
			return
		}
		// Unmarshal the message into a WebSocketMessage struct
		err = json.Unmarshal(p, &wsm)
		if err != nil {
			log.Println(err.Error())
			return
		}
		// Handle the received message
		c.HandleMessage(wsm)
	}
}

// Send sends a message to the WebSocket connection
func (c *Client) Send(v any) {
	err := c.Connection.WriteJSON(v)
	if err != nil {
		log.Println(err.Error())
		return
	}
}

// HandleMessage processes incoming WebSocket messages
func (c *Client) HandleMessage(msg WebSocketMessage) {
	switch msg.Code {
	case GROUP_CHAT_MESSAGE:
		var message models.Message

		// Handle group chat message
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
		// Broadcast the message to the group
		group.Broadcast <- msg
		// Store the message in the database
		c.Repo.CreateMessage(message)
		log.Println("Group Message added to db in Client.go is:", message)

	case PRIVATE_MESSAGE:
		var message models.Message

		// Handle private message
		err := json.Unmarshal([]byte(msg.Body), &message)
		if err != nil {
			log.Println(err.Error())
			return
		}
		c.Repo.CreateMessage(message)

		group, ok := c.SocketGroups[0]
		if !ok {
			log.Println("primary group does not exist")
			return
		}
		// Broadcast the message to the main group (group 0)
		group.Broadcast <- msg
		// Store the message in the database
	case FOLLOW_REQUEST:
		var notification models.Notification
		// Handle private message
		err := json.Unmarshal([]byte(msg.Body), &notification)
		if err != nil {
			log.Println(err.Error())
			return
		}
		group, ok := c.SocketGroups[0]
		if !ok {
			log.Println("Private message group does not exist")
			return
		}

		// Broadcast the message to the main group (group 0)
		group.Broadcast <- msg
		// Store the message in the database
		c.Repo.CreateNotification(notification)
	case EVENT_INVITE:
		ctime := time.Now().UTC().UnixMilli()

		// Handle event invite
		var event models.Event
		err := json.Unmarshal([]byte(msg.Body), &event)
		if err != nil {
			log.Println(err.Error())
			return
		}
		event.CreatedAt = ctime
		event.UpdatedAt = ctime
		// Adds Event to db
		returnEvent, err := c.Repo.CreateEvent(event)
		if err != nil {
			log.Println(err.Error())
			return
		}

		// Adds user who made event to eventUsers table.  It's their event, they better be going!
		eventUserWhoMadeEvent := models.EventUser{
			CreatedAt: ctime,
			EventId:   returnEvent.EventId,
			IsGoing:   true,
			UpdatedAt: ctime,
			UserId:    event.UserId,
		}
		c.Repo.CreateEventUser(eventUserWhoMadeEvent)
		groupId := event.GroupId
		group, ok := c.SocketGroups[groupId]
		if !ok {
			log.Printf("SocketGroup %d does not exist\n", groupId)
			return
		}
		// retrieves all members of the event's group
		groupUsers, err := c.Repo.GetGroupUsersByGroupId(event.GroupId)
		if err != nil {
			utils.HandleError("Error in GetGroupUsersByGroupId, in ws/Client.go.", err)
			return
		}
		for i := 0; i < len(groupUsers); i++ {
			// so long as the member of the group is not the person who made the event (they're automatically attending the event)
			if groupUsers[i].UserId != event.UserId {
				// Create a notification for each user in the group
				notification := models.Notification{
					CreatedAt:        ctime,
					NotificationType: "eventInvite",
					ObjectId:         returnEvent.EventId,
					SenderId:         returnEvent.UserId,
					Status:           "pending",
					TargetId:         groupUsers[i].UserId,
					UpdatedAt:        ctime,
				}
				// Store the notification in the database
				returnNotification, err := c.Repo.CreateNotification(notification)
				if err != nil {
					utils.HandleError("Error in CreateNotification, in ws/Client.go.", err)
				}

				jsonNotification, err := json.Marshal(returnNotification)
				if err != nil {
					utils.HandleError("[ws/client.go] Error marshalling returnNotification", err)
					continue
				}
				returnMessage := WebSocketMessage{
					Code: 6,
					Body: string(jsonNotification),
				}
				_, ok := group.Clients[groupUsers[i].UserId]
				if ok {
					group.Broadcast <- returnMessage
				}
			}

		}
	}
}
