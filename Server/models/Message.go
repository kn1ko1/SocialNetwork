package models

import (
	"errors"
	"math/rand"
)

type Message struct {
	MessageId   int    `json:"messageId"`
	Body        string `json:"body"`
	CreatedAt   int64  `json:"createdAt"`
	MessageType string `json:"messageType"`
	SenderId    int    `json:"senderId"`
	TargetId    int    `json:"targetId"`
	UpdatedAt   int64  `json:"updatedAt"`
}

func (m *Message) Validate() error {
	// Validate logic here
	if m.Body == "" {
		return errors.New("invalid 'Body' field")
	}
	if m.CreatedAt <= 0 {
		return errors.New("invalid 'CreatedAt' field")
	}
	if !(m.MessageType == "GC" || m.MessageType == "DM") {
		return errors.New("invalid 'MessageType' field")
	}
	if m.SenderId <= 0 {
		return errors.New("invalid 'SenderId' field")
	}
	if m.TargetId <= 0 {
		return errors.New("invalid 'TargetId' field")
	}
	if m.UpdatedAt < m.CreatedAt {
		return errors.New("invalid 'UpdatedAt' field")
	}
	return nil
}

func GenerateValidMessage() *Message {
	ctime := rand.Int63n(1000) + 1
	idxBody := rand.Intn(len(sutBody))
	idxMessageType := rand.Intn(len(sutMessageTypes))

	m := &Message{
		Body:        sutBody[idxBody],
		CreatedAt:   ctime,
		MessageType: sutMessageTypes[idxMessageType],
		SenderId:    rand.Intn(1000) + 1,
		TargetId:    rand.Intn(1000) + 1,
		UpdatedAt:   ctime,
	}
	return m
}

func GenerateMissingFieldMessage() *Message {
	m := GenerateValidMessage()
	missingField := rand.Intn(6)
	switch missingField {
	case 0:
		m.Body = ""
	case 1:
		m.CreatedAt = 0
	case 2:
		m.MessageType = ""
	case 3:
		m.SenderId = 0
	case 4:
		m.TargetId = 0
	case 5:
		m.UpdatedAt = 0
	}
	return m
}

func GenerateInvalidMessage() *Message {
	m := GenerateValidMessage()
	invalidField := rand.Intn(6)
	switch invalidField {
	case 0:
		m.Body = ""
	case 1:
		m.CreatedAt = -m.CreatedAt
	case 2:
		m.MessageType = ""
	case 3:
		m.SenderId = -m.SenderId
	case 4:
		m.TargetId = -m.TargetId
	case 5:
		m.UpdatedAt = -m.UpdatedAt
	}
	return m
}
