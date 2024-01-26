package models

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestMessageValidateInvalidFieldExpectError(t *testing.T) {
	var messages []*Message
	for i := 0; i < tableRunCount; i++ {
		messages = append(messages, generateInvalidMessage())
	}
	for _, m := range messages {
		name := fmt.Sprintf("%+v", *m)
		t.Run(name, func(t *testing.T) {
			err := m.Validate()
			if err == nil {
				t.Error("expect error for invalid event field")
			}
		})
	}
}

func TestMessageValidateMissingFieldExpectError(t *testing.T) {
	var messages []*Message
	for i := 0; i < tableRunCount; i++ {
		messages = append(messages, generateMissingFieldMessage())
	}
	for _, m := range messages {
		name := fmt.Sprintf("%+v", *m)
		t.Run(name, func(t *testing.T) {
			err := m.Validate()
			if err == nil {
				t.Error("expect error for missing event field")
			}
		})
	}
}

func TestMessageValidateValidExpectNil(t *testing.T) {
	var messages []*Message
	for i := 0; i < tableRunCount; i++ {
		messages = append(messages, GenerateValidMessage())
	}
	for _, m := range messages {
		name := fmt.Sprintf("%+v", *m)
		t.Run(name, func(t *testing.T) {
			err := m.Validate()
			if err != nil {
				t.Error("expect nil for valid event")
			}
		})
	}
}

func generateMissingFieldMessage() *Message {
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

func generateInvalidMessage() *Message {
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
