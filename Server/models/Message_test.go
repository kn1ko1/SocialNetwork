package models

import (
	"fmt"
	"testing"
)

func TestMessageValidateInvalidFieldExpectError(t *testing.T) {
	var messages []*Message
	for i := 0; i < tableRunCount; i++ {
		messages = append(messages, GenerateInvalidMessage())
	}
	for _, m := range messages {
		name := fmt.Sprintf("%+v", *m)
		t.Run(name, func(t *testing.T) {
			err := m.Validate()
			if err == nil {
				t.Error("expect error for invalid message field")
			}
		})
	}
}

func TestMessageValidateMissingFieldExpectError(t *testing.T) {
	var messages []*Message
	for i := 0; i < tableRunCount; i++ {
		messages = append(messages, GenerateMissingFieldMessage())
	}
	for _, m := range messages {
		name := fmt.Sprintf("%+v", *m)
		t.Run(name, func(t *testing.T) {
			err := m.Validate()
			if err == nil {
				t.Error("expect error for missing message field")
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
				t.Error("expect nil for valid message")
			}
		})
	}
}
