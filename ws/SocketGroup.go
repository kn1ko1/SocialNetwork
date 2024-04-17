package ws

import (
	"encoding/json"
	"fmt"
	"log"
)

type SocketGroup struct {
	SocketGroupID int
	Clients       map[int]*Client
	Enter         chan *Client
	Exit          chan *Client
	Broadcast     chan WebSocketMessage
}

func NewSocketGroup(id int) *SocketGroup {
	ret := new(SocketGroup)
	ret.SocketGroupID = id
	ret.Clients = make(map[int]*Client)
	ret.Enter = make(chan *Client)
	ret.Exit = make(chan *Client)
	ret.Broadcast = make(chan WebSocketMessage)
	return ret
}

func (g *SocketGroup) Run() {
	for {
		select {
		case c := <-g.Enter:
			g.Clients[c.ClientID] = c
			fmt.Printf("User %d entered.\n", c.ClientID)
			// Add to broadcast
		case c := <-g.Exit:
			delete(g.Clients, c.ClientID)
			fmt.Printf("User %d exited.\n", c.ClientID)
			// Add to broadcast
		case msg := <-g.Broadcast:
			switch msg.Code {
			case PRIVATE_MESSAGE:
				var body PrivateMessageBody
				err := json.Unmarshal([]byte(msg.Body), &body)
				if err != nil {
					log.Println(err.Error())
				}
				c := g.Clients[body.TargetUserID]
				c.Send(msg)
			case GROUP_CHAT_MESSAGE:
				var body GroupChatBody
				err := json.Unmarshal([]byte(msg.Body), &body)
				if err != nil {
					log.Println(err.Error())
				}
				for _, c := range g.Clients {
					c.Send(msg)
					log.Println("I'm in run in ws/SocketGroup.go")
				}
			}

		}
	}
}
