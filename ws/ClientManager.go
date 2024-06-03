package ws

type ClientManager struct {
	clients map[int]*Client
}

func NewClientManager() *ClientManager {
	return &ClientManager{
		clients: make(map[int]*Client),
	}
}

func (cm *ClientManager) RegisterClient(client *Client) {
	cm.clients[client.ClientID] = client
}

func (cm *ClientManager) UnregisterClient(clientID int) {
	delete(cm.clients, clientID)
}

func (cm *ClientManager) GetClient(clientID int) (*Client, bool) {
	client, ok := cm.clients[clientID]
	return client, ok
}
