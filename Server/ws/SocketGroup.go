package ws

import (
	"encoding/json"
	"fmt"
	"log"
	"socialnetwork/Server/models"

	"socialnetwork/Server/transport"
)

// const (
// 	GROUP_CHAT_MESSAGE = 1
// 	PRIVATE_MESSAGE    = 2
// 	FOLLOW_REQUEST       = 3
// 	GROUP_REQUEST      = 4
// 	GROUP_INVITE       = 5
// 	CREATE_EVENT       = 6
// )

// SocketGroup represents a group of WebSocket clients.
type SocketGroup struct {
	SocketGroupID int                   // ID of the socket group.
	Clients       map[int]*Client       // Map of client IDs to Client instances.
	Enter         chan *Client          // Channel for entering clients.
	Exit          chan *Client          // Channel for exiting clients.
	Broadcast     chan WebSocketMessage // Channel for broadcasting messages to clients.
}

// NewSocketGroup creates a new SocketGroup instance
func NewSocketGroup(id int) *SocketGroup {
	ret := new(SocketGroup)
	ret.SocketGroupID = id
	ret.Clients = make(map[int]*Client)
	ret.Enter = make(chan *Client)
	ret.Exit = make(chan *Client)
	ret.Broadcast = make(chan WebSocketMessage)
	return ret
}

// Run starts the SocketGroup's main loop for handling client events.
func (g *SocketGroup) Run() {
	for {
		select {
		case c := <-g.Enter:
			g.Clients[c.ClientID] = c
			fmt.Printf("User %d entered.\n", c.ClientID)
			// TODO: Add logic for broadcasting the "user entered" event.
		case c := <-g.Exit:
			delete(g.Clients, c.ClientID)
			fmt.Printf("User %d exited.\n", c.ClientID)
			// TODO: Add logic for broadcasting the "user exited" event.
		case msg := <-g.Broadcast:
			// Handle different types of messages.
			switch msg.Code {

			case GROUP_CHAT_MESSAGE:
				var message transport.MessageTransport

				err := json.Unmarshal([]byte(msg.Body), &message)
				if err != nil {
					log.Println(err.Error())
				}
				log.Println("Group Message Body in socketGroup is:", message)

				// Broadcast the message to all clients in the group.
				if len(g.Clients) > 1 {
					for _, c := range g.Clients {
						if c.User.Username != message.SenderUsername {
							c.Send(msg)
						}

					}
				}
			case PRIVATE_MESSAGE:
				var message models.Message

				err := json.Unmarshal([]byte(msg.Body), &message)
				if err != nil {
					log.Println(err.Error())
				}
				log.Println("Private Message Body in socketGroup is:", message)

				c, ok := g.Clients[message.TargetId]
				if !ok {
					log.Printf("Target client %d not found for private message\n", message.TargetId)
				} else {
					c.Send(msg)
				}
			case FOLLOW_REQUEST:
				var notification models.Notification
				err := json.Unmarshal([]byte(msg.Body), &notification)
				if err != nil {
					log.Println(err.Error())
				}
				log.Println("FOLLOW_REQUEST Body in socketGroup is:", notification)

				c, ok := g.Clients[notification.TargetId]
				if !ok {
					log.Printf("Target client %d not found for FOLLOW_REQUEST\n", notification.TargetId)
				} else {
					c.Send(msg)
				}
			case CREATE_EVENT:
				var notification models.Notification
				err := json.Unmarshal([]byte(msg.Body), &notification)
				if err != nil {
					log.Println(err.Error())
				}

				// Broadcast the message to target within the group.
				targetClient, ok := g.Clients[notification.TargetId]
				if ok {
					targetClient.Send(msg)
				}
			case GROUP_INVITE:
				var notification models.Notification
				err := json.Unmarshal([]byte(msg.Body), &notification)
				if err != nil {
					log.Println(err.Error())
				}

				c, ok := g.Clients[notification.TargetId]
				if !ok {
					log.Printf("Target client %d not found for group invite\n", notification.TargetId)
				} else {
					c.Send(msg)
				}

			case GROUP_REQUEST:
				var notification models.Notification
				err := json.Unmarshal([]byte(msg.Body), &notification)
				if err != nil {
					log.Println(err.Error())
				}

				c, ok := g.Clients[notification.TargetId]
				if !ok {
					log.Printf("Target client %d not found for group invite\n", notification.TargetId)
				} else {
					c.Send(msg)
				}

			}

		}
	}
}
