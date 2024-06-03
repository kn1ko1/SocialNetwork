package ws

import (
	"encoding/json"
	"fmt"
	"log"
	"socialnetwork/models"
	"socialnetwork/repo"
	"time"
)

// SocketGroup represents a group of WebSocket clients.
type SocketGroup struct {
	SocketGroupID int                   // ID of the socket group.
	Clients       map[int]*Client       // Map of client IDs to Client instances.
	Enter         chan *Client          // Channel for entering clients.
	Exit          chan *Client          // Channel for exiting clients.
	Broadcast     chan WebSocketMessage // Channel for broadcasting messages to clients.
	Repo          repo.IRepository
}

// NewSocketGroup creates a new instance of SocketGroup.
func NewSocketGroup(id int) *SocketGroup {
	ret := new(SocketGroup)
	ret.SocketGroupID = id
	ret.Clients = make(map[int]*Client)
	ret.Enter = make(chan *Client)
	ret.Exit = make(chan *Client)
	ret.Broadcast = make(chan WebSocketMessage)
	ret.Repo = repo.NewSQLiteRepository()
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
				// Find the target user and send the message.
				ctime := time.Now().UTC().UnixMilli()
				message.CreatedAt = ctime
				message.UpdatedAt = ctime
				ret, err := g.Repo.CreateMessage(message)
				if err != nil {
					log.Println(err.Error())
				}
				log.Println("Group Message added to db in socketGroup is:", ret)

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
				ctime := time.Now().UTC().UnixMilli()
				message.CreatedAt = ctime
				message.UpdatedAt = ctime
				ret, err := g.Repo.CreateMessage(message)
				if err != nil {
					log.Println(err.Error())
				}
				log.Println("Group Message added to db in socketGroup is:", ret)

				// Broadcast the message to all clients in the group.
				if len(g.Clients) > 1 {
					for _, c := range g.Clients {
						c.Send(msg)
					}
				}

			}

		}
	}
}
