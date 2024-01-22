package models

import "errors"

type Message struct {
	MessageId   int
	Body        string
	CreatedAt   int64
	MessageType string
	SenderId    int
	TargetId    int
	UpdatedAt   int64
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
