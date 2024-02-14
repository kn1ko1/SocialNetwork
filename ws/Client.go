package ws

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	ClientID   int
	Connection *websocket.Conn
}

func NewClient(conn *websocket.Conn) *Client {
	return &Client{Connection: conn}
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
	for {
		err := c.Connection.WriteJSON(v)
		if err != nil {
			log.Println(err.Error())
			return
		}
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
	case 2:
		var p Person
		err := json.Unmarshal([]byte(msg.Body), &p)
		if err != nil {
			log.Println(err.Error())
		}
		fmt.Printf("%+v\n", p)
	}
}
