package ws

import "fmt"

type Group struct {
	Clients   map[int]*Client
	Enter     chan *Client
	Exit      chan *Client
	Broadcast chan WebSocketMessage
}

func NewGroup() *Group {
	ret := new(Group)
	ret.Clients = make(map[int]*Client)
	ret.Enter = make(chan *Client)
	ret.Exit = make(chan *Client)
	ret.Broadcast = make(chan WebSocketMessage)
	return ret
}

func (g *Group) Run() {
	for {
		select {
		case c := <-g.Enter:
			g.Clients[c.ClientID] = c
			fmt.Printf("User %d entered.\n", c.ClientID)
		case c := <-g.Exit:
			delete(g.Clients, c.ClientID)
			fmt.Printf("User %d exited.\n", c.ClientID)
		case msg := <-g.Broadcast:
			for _, c := range g.Clients {
				c.Send(msg)
			}
		}
	}
}
