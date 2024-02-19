package ws

import (
	"fmt"
	"log"
)

type Event struct {
	EventID   int
	Clients   map[int]*Client
	Enter     chan *Client
	Exit      chan *Client
	Broadcast chan WebSocketMessage
}

func NewEvent(id int) *Event {
	ret := new(Event)
	ret.EventID = id
	ret.Clients = make(map[int]*Client)
	ret.Enter = make(chan *Client)
	ret.Exit = make(chan *Client)
	ret.Broadcast = make(chan WebSocketMessage)
	return ret
}

func (e *Event) Run() {
	for {
		select {
		case c := <-e.Enter:
			e.Clients[c.ClientID] = c
			fmt.Printf("User %d entered.\n", c.ClientID)
			// Add to broadcast
		case c := <-e.Exit:
			delete(e.Clients, c.ClientID)
			fmt.Printf("User %d exited.\n", c.ClientID)
			// Add to broadcast
		case msg := <-e.Broadcast:
			for _, c := range e.Clients {
				c.Send(msg)
				log.Println("I'm in run in ws/Event.go")
			}
		}
	}
}
