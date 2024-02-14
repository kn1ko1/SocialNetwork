package ws

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	Connection *websocket.Conn
}

func NewClient(conn *websocket.Conn) *Client {
	return &Client{Connection: conn}
}

func (c *Client) Receive() {
	for {
		var wsm WebSocketMessage
		err := c.Connection.ReadJSON(&wsm)
		if err != nil {
			log.Println(err.Error())
			return
		}
		fmt.Println(wsm)
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
