package ws

import (
	"encoding/json"
	"log"
	"socialnetwork/models"
	"socialnetwork/repo"
	"socialnetwork/transport"
	"socialnetwork/utils"
	"time"

	"github.com/gorilla/websocket"
)

// Define constants for different types of WebSocket messages
const (
	GROUP_CHAT_MESSAGE        = 1
	PRIVATE_MESSAGE           = 2
	FOLLOW_REQUEST            = 3
	GROUP_REQUEST             = 4
	GROUP_INVITE              = 5
	EVENT_INVITE              = 6
	CreateGroupAndSocketGroup = 10
	AcceptGroupAndSocketGroup = 11
)

// Client represents a connected user
type Client struct {
	ClientID     int
	Connection   *websocket.Conn
	SocketGroups map[int]*SocketGroup
	User         models.User
	Repo         repo.IRepository
}

// NewClient creates a new Client instance
func NewClient(conn *websocket.Conn, user models.User, repo repo.IRepository) *Client {
	return &Client{
		ClientID:     user.UserId,
		Connection:   conn,
		SocketGroups: make(map[int]*SocketGroup),
		User:         user,
		Repo:         repo,
	}
}

// Receive listens for incoming messages from the WebSocket connection
func (c *Client) Receive() {
	defer func() {
		// Remove the client from all SocketGroups when the connection is closed
		for _, group := range c.SocketGroups {
			group.Exit <- c
		}
		// Close the WebSocket connection
		c.Connection.Close()
	}()
	for {
		var wsm WebSocketMessage
		// Read a message from the WebSocket connection
		_, p, err := c.Connection.ReadMessage()
		if err != nil {
			log.Println(err.Error())
			return
		}
		// Unmarshal the message into a WebSocketMessage struct
		err = json.Unmarshal(p, &wsm)
		if err != nil {
			log.Println(err.Error())
			return
		}
		switch wsm.Code {
		case 10:
			c.CreateGroupAndSocketGroup(wsm)
		case 11:
			c.AcceptGroupAndSocketGroup(wsm)
		default:
			// Handle the received message
			c.HandleMessage(wsm)
		}

	}
}

// Send sends a message to the WebSocket connection
func (c *Client) Send(v any) {
	err := c.Connection.WriteJSON(v)
	if err != nil {
		log.Println(err.Error())
		return
	}
}

// CreateGroupAndSocketGroup processes messages for creating groups and associated socketGroup
func (c *Client) CreateGroupAndSocketGroup(msg WebSocketMessage) {
	//switch msg.Code {
	//case GROUP_CHAT_MESSAGE:
	var group models.Group
	ctime := time.Now().UTC().UnixMilli()

	err := json.Unmarshal([]byte(msg.Body), &group)
	if err != nil {
		log.Println(err.Error())
		return
	}
	group.CreatedAt = ctime
	group.UpdatedAt = ctime
	// Validate the group
	if validationErr := group.Validate(); validationErr != nil {
		utils.HandleError("Validation failed:", validationErr)
		return
	}
	createdGroup, err := c.Repo.CreateGroup(group)
	if err != nil {
		utils.HandleError("Error in CreateGroup, in ws/Client.go.", err)
		return
	}

	groupUser := models.GroupUser{
		CreatedAt: ctime,
		GroupId:   createdGroup.GroupId,
		UpdatedAt: ctime,
		UserId:    createdGroup.CreatorId,
	}

	_, createGroupUserErr := c.Repo.CreateGroupUser(groupUser)
	if createGroupUserErr != nil {
		utils.HandleError("Failed to add user to groupUser table. ", createGroupUserErr)
		return
	}

	log.Println("CLient.Go.  Trying to create socket group", createdGroup.GroupId)
	// Check if the socket group exists, and create it if it doesn't.
	_, exists := socketGroupManager.SocketGroups[createdGroup.GroupId]
	if !exists {
		socketGroupManager.SocketGroups[createdGroup.GroupId] = NewSocketGroup(createdGroup.GroupId)
		go socketGroupManager.SocketGroups[createdGroup.GroupId].Run() // Run the socket group in a separate goroutine.
		log.Printf("Client.go.  Socket group %d created and Run() called", createdGroup.GroupId)
		log.Printf("SocketGroups map now contains: %v", socketGroupManager.SocketGroups)
	}
	// Add the client to the socket group.

	c.SocketGroups[createdGroup.GroupId] = socketGroupManager.SocketGroups[createdGroup.GroupId]
	socketGroupManager.SocketGroups[createdGroup.GroupId].Enter <- c
	//	}

	//case PRIVATE_MESSAGE:

	//}
}

// CreateGroupAndSocketGroup processes messages for accepting group requests/invites and  joining/starting associated socketGroup
func (c *Client) AcceptGroupAndSocketGroup(msg WebSocketMessage) {
	var notification models.Notification

	log.Println("Starting AcceptGroupAndSocketGroup") // Log for function start

	err := json.Unmarshal([]byte(msg.Body), &notification)
	if err != nil {
		log.Println("Unmarshal error:", err.Error())
		return
	}
	log.Println("Unmarshal successful") // Log after successful unmarshal

	log.Println("Trying to create socket group", notification.ObjectId) // Log before creating socket group

	// Check if the socket group exists, and create it if it doesn't.
	_, exists := socketGroupManager.SocketGroups[notification.ObjectId]
	if !exists {
		socketGroupManager.SocketGroups[notification.ObjectId] = NewSocketGroup(notification.ObjectId)
		go socketGroupManager.SocketGroups[notification.ObjectId].Run() // Run the socket group in a separate goroutine.
		log.Printf("Client.go.  Socket group %d created and Run() called", notification.ObjectId)
		log.Printf("SocketGroups map now contains: %v", socketGroupManager.SocketGroups)
	} else {
		log.Println("Socket group", notification.ObjectId, "already exists") // Log if socket group already exists
	}

	// Add the client to the socket group.
	c.SocketGroups[notification.ObjectId] = socketGroupManager.SocketGroups[notification.ObjectId]
	log.Println("Client added to socket group") // Log after adding client to socket group

	log.Println()
	log.Println("Client.go. Print me, print me, I'm here you damned code!!!!", c.SocketGroups) // Log we want to reach
	log.Println()

	socketGroupManager.SocketGroups[notification.ObjectId].Enter <- c
	log.Println("Client added to socket group's Enter channel") // Log after sending client to Enter channel
}

// HandleMessage processes incoming WebSocket messages
func (c *Client) HandleMessage(msg WebSocketMessage) {
	switch msg.Code {
	case GROUP_CHAT_MESSAGE:
		var message models.Message

		// Handle group chat message
		err := json.Unmarshal([]byte(msg.Body), &message)
		if err != nil {
			log.Println(err.Error())
			return
		}

		// Store the message in the database
		returnGroupMessage, err := c.Repo.CreateMessage(message)
		if err != nil {
			utils.HandleError("[ws/client.go] Error adding message to database in CreateMessage", err)
		}
		transportMessage := transport.MessageTransport{
			MessageId:      returnGroupMessage.MessageId,
			Body:           returnGroupMessage.Body,
			CreatedAt:      returnGroupMessage.CreatedAt,
			MessageType:    returnGroupMessage.MessageType,
			SenderUsername: c.User.Username,
			TargetId:       returnGroupMessage.TargetId,
			UpdatedAt:      returnGroupMessage.UpdatedAt,
		}
		jsonGroupMessage, err := json.Marshal(transportMessage)
		if err != nil {
			utils.HandleError("[ws/client.go] Error marshalling returnNotification", err)
		}
		returnMsg := WebSocketMessage{
			Code: 1,
			Body: string(jsonGroupMessage),
		}
		groupId := message.TargetId
		group, ok := c.SocketGroups[groupId]
		if !ok {
			log.Printf("SocketGroup %d does not exist\n", groupId)
			return
		}

		// Broadcast the message to the group
		group.Broadcast <- returnMsg

	case PRIVATE_MESSAGE:

		var message models.Message

		// Handle private message
		err := json.Unmarshal([]byte(msg.Body), &message)
		if err != nil {
			log.Println(err.Error())
			return
		}

		returnPrivateMessage, err := c.Repo.CreateMessage(message)
		if err != nil {
			utils.HandleError("[ws/client.go] Error adding message to database in CreateMessage", err)
		}
		transportMessage := transport.MessageTransport{
			MessageId:      returnPrivateMessage.MessageId,
			Body:           returnPrivateMessage.Body,
			CreatedAt:      returnPrivateMessage.CreatedAt,
			MessageType:    returnPrivateMessage.MessageType,
			SenderUsername: c.User.Username,
			TargetId:       returnPrivateMessage.TargetId,
			UpdatedAt:      returnPrivateMessage.UpdatedAt,
		}

		jsonPrivateMessage, err := json.Marshal(transportMessage)
		if err != nil {
			utils.HandleError("[ws/client.go] Error marshalling returnNotification", err)
		}
		group, ok := c.SocketGroups[0]
		if !ok {
			log.Println("primary group does not exist")
			return
		}
		returnMsg := WebSocketMessage{
			Code: 2,
			Body: string(jsonPrivateMessage),
		}
		// Broadcast the message to the main group (group 0)
		group.Broadcast <- returnMsg
		// Store the message in the database
	case FOLLOW_REQUEST:
		ctime := time.Now().UTC().UnixMilli()

		var notification models.Notification
		// Handle private message
		err := json.Unmarshal([]byte(msg.Body), &notification)
		if err != nil {
			log.Println(err.Error())
			return
		}
		notification.CreatedAt = ctime
		notification.UpdatedAt = ctime
		returnNotification, err := c.Repo.CreateNotification(notification)
		if err != nil {
			utils.HandleError("Error in CreateNotification, in ws/Client.go.", err)
			return
		}

		jsonNotification, err := json.Marshal(returnNotification)
		if err != nil {
			utils.HandleError("[ws/client.go] Error marshalling returnNotification", err)
		}
		returnMsg := WebSocketMessage{
			Code: 3,
			Body: string(jsonNotification),
		}

		group, ok := c.SocketGroups[0]
		if !ok {
			log.Println("primary socket group does not exist")
			return
		}

		// Broadcast the message to the main group (group 0)
		group.Broadcast <- returnMsg
		// Store the message in the database

	case GROUP_INVITE:

		ctime := time.Now().UTC().UnixMilli()

		var notification models.Notification

		// Handle group invite
		err := json.Unmarshal([]byte(msg.Body), &notification)
		if err != nil {
			log.Println(err.Error())
			return
		}

		notification.CreatedAt = ctime
		notification.UpdatedAt = ctime

		returnNotification, err := c.Repo.CreateNotification(notification)
		if err != nil {
			utils.HandleError("Failed to create notification. ", err)
			return
		}

		jsonNotification, err := json.Marshal(returnNotification)
		if err != nil {
			utils.HandleError("[ws/client.go] Error marshalling returnNotification", err)
			return
		}
		returnMsg := WebSocketMessage{
			Code: 5,
			Body: string(jsonNotification),
		}

		group, ok := c.SocketGroups[0]
		if !ok {
			log.Println("primary group does not exist")
			return
		}
		// Broadcast the message to the main group (group 0)
		group.Broadcast <- returnMsg
		// Store the message in the database
	case GROUP_REQUEST:
		ctime := time.Now().UTC().UnixMilli()
		var notification models.Notification

		// Handle group invite
		err := json.Unmarshal([]byte(msg.Body), &notification)
		if err != nil {
			log.Println(err.Error())
			return
		}

		notification.CreatedAt = ctime
		notification.UpdatedAt = ctime

		returnNotification, err := c.Repo.CreateNotification(notification)
		if err != nil {
			utils.HandleError("Failed to create notification. ", err)
			return
		}

		jsonNotification, err := json.Marshal(returnNotification)
		if err != nil {
			utils.HandleError("[ws/client.go] Error marshalling returnNotification", err)
			return
		}
		returnMsg := WebSocketMessage{
			Code: 4,
			Body: string(jsonNotification),
		}

		group, ok := c.SocketGroups[0]
		if !ok {
			log.Println("primary group does not exist")
			return
		}
		// Broadcast the message to the main group (group 0)
		group.Broadcast <- returnMsg
		// Store the message in the database

	case EVENT_INVITE:

		ctime := time.Now().UTC().UnixMilli()

		// Handle event invite
		var event models.Event
		err := json.Unmarshal([]byte(msg.Body), &event)
		if err != nil {
			log.Println(err.Error())
			return
		}
		log.Printf("Handling EVENT_INVITE. GroupId: %d", event.GroupId)
		event.CreatedAt = ctime
		event.UpdatedAt = ctime
		// Adds Event to db
		returnEvent, err := c.Repo.CreateEvent(event)
		if err != nil {
			log.Println(err.Error())
			return
		}

		// Adds user who made event to eventUsers table.  It's their event, they better be going!
		eventUserWhoMadeEvent := models.EventUser{
			CreatedAt: ctime,
			EventId:   returnEvent.EventId,
			IsGoing:   true,
			UpdatedAt: ctime,
			UserId:    event.UserId,
		}
		c.Repo.CreateEventUser(eventUserWhoMadeEvent)
		groupId := event.GroupId
		group, ok := c.SocketGroups[groupId]
		if !ok {
			log.Printf("SocketGroup %d does not exist\n", groupId)
			return
		}
		// retrieves all members of the event's group
		groupUsers, err := c.Repo.GetGroupUsersByGroupId(event.GroupId)
		if err != nil {
			utils.HandleError("Error in GetGroupUsersByGroupId, in ws/Client.go.", err)
			return
		}
		for i := 0; i < len(groupUsers); i++ {
			// so long as the member of the group is not the person who made the event (they're automatically attending the event)
			if groupUsers[i].UserId != event.UserId {
				// Create a notification for each user in the group
				notification := models.Notification{
					CreatedAt:        ctime,
					NotificationType: "eventInvite",
					ObjectId:         returnEvent.EventId,
					SenderId:         returnEvent.UserId,
					Status:           "pending",
					TargetId:         groupUsers[i].UserId,
					UpdatedAt:        ctime,
				}
				// Store the notification in the database
				returnNotification, err := c.Repo.CreateNotification(notification)
				if err != nil {
					utils.HandleError("Error in CreateNotification, in ws/Client.go.", err)
				}

				jsonNotification, err := json.Marshal(returnNotification)
				if err != nil {
					utils.HandleError("[ws/client.go] Error marshalling returnNotification", err)
					continue
				}
				returnMsg := WebSocketMessage{
					Code: 6,
					Body: string(jsonNotification),
				}
				_, ok := group.Clients[groupUsers[i].UserId]
				if ok {
					group.Broadcast <- returnMsg
				}

			}

		}
	}
}
