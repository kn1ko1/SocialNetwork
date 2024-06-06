package ws

import (
	"encoding/json"
	"fmt"
	"log"
	"socialnetwork/models"
)

// const (
// 	GROUP_CHAT_MESSAGE = 1
// 	PRIVATE_MESSAGE    = 2
// 	CREATE_EVENT       = 3
// 	GROUP_REQUEST      = 4
// 	GROUP_INVITE       = 5
// 	EVENT_INVITE       = 6
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
	var message models.Message
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
			case PRIVATE_MESSAGE:
				err := json.Unmarshal([]byte(msg.Body), &message)
				if err != nil {
					log.Println(err.Error())
				}
				log.Println("Private Message Body in socketGroup is:", message)

				c, ok := g.Clients[message.TargetId]
				if !ok {
					log.Printf("Target client %d not found for private message\n", message.TargetId)
					return
				}
				c.Send(msg)
			case GROUP_CHAT_MESSAGE:
				err := json.Unmarshal([]byte(msg.Body), &message)
				if err != nil {
					log.Println(err.Error())
				}
				log.Println("Group Message Body in socketGroup is:", message)
				// Persist the group message to the database.
				// ctime := time.Now().UTC().UnixMilli()
				// message.CreatedAt = ctime
				// message.UpdatedAt = ctime
				// ret, err := g.Repo.CreateMessage(message)
				// if err != nil {
				// 	log.Println(err.Error())
				// }
				// log.Println("Group Message added to db in socketGroup is:", ret)

				// Broadcast the message to all clients in the group.
				if len(g.Clients) > 1 {
					for _, c := range g.Clients {
						c.Send(msg)
					}
				}
			case EVENT_INVITE:
				var notification models.Notification
				err := json.Unmarshal([]byte(msg.Body), &notification)
				if err != nil {
					log.Println(err.Error())
				}
				log.Println("Event Body in socketGroup is:", notification)

				// Broadcast the message to all clients in the group.
				targetClient, ok := g.Clients[notification.TargetId]
				if ok {
					targetClient.Send(msg)
				}
			}

		}
	}
}
