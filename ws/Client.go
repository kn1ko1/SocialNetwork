package ws

import (
	"encoding/json"
	"fmt"
	"log"
	"socialnetwork/models"

	"github.com/gorilla/websocket"
)

type Client struct {
	ClientID   int
	Connection *websocket.Conn
	Groups     map[int]*Group
	Events     map[int]*Event
}

func NewClient(conn *websocket.Conn) *Client {
	return &Client{
		Connection: conn,
		Groups:     make(map[int]*Group),
		Events:     make(map[int]*Event),
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
	fmt.Println(msg)
	switch msg.Code {
	case 1:
		var t TestBody
		err := json.Unmarshal([]byte(msg.Body), &t)
		if err != nil {
			log.Println(err.Error())
		}
		fmt.Printf("%+v\n", t)
		fmt.Println("1 testBody")

	case 2:
		var p Person
		err := json.Unmarshal([]byte(msg.Body), &p)
		if err != nil {
			log.Println(err.Error())
		}
		fmt.Printf("%+v\n", p)
		fmt.Println("2 person")

	case 3:
		var event models.Event
		err := json.Unmarshal([]byte(msg.Body), &event)
		if err != nil {
			log.Println(err.Error())
		}
		fmt.Printf("%+v\n", event)
		fmt.Println("3 event")
		fmt.Println(msg.Body)
		c.Events[0].Broadcast <- msg
	}
	// c.Groups[0].Broadcast <- msg
}
