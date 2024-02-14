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
		fmt.Println(wsm)
		switch wsm.Code {
		case 1:
			var t TestBody
			err := json.Unmarshal([]byte(wsm.Body), &t)
			if err != nil {
				log.Println(err.Error())
			}
			fmt.Printf("%+v\n", t)
		}
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
