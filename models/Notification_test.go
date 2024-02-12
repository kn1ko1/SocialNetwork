package models

import (
	"fmt"
	"testing"
)

func TestNotificationValidateInvalidFieldExpectError(t *testing.T) {
	var notifications []*Notification
	for i := 0; i < tableRunCount; i++ {
		notifications = append(notifications, GenerateInvalidNotification())
	}
	for _, n := range notifications {
		name := fmt.Sprintf("%+v", *n)
		t.Run(name, func(t *testing.T) {
			err := n.Validate()
			if err == nil {
				t.Error("expect error for invalid notification field")
			}
		})
	}
}

func TestNotificationValidateMissingFieldExpectError(t *testing.T) {
	var notifications []*Notification
	for i := 0; i < tableRunCount; i++ {
		notifications = append(notifications, GenerateMissingFieldNotification())
	}
	for _, n := range notifications {
		name := fmt.Sprintf("%+v", *n)
		t.Run(name, func(t *testing.T) {
			err := n.Validate()
			if err == nil {
				t.Error("expect error for missing notification field")
			}
		})
	}
}

func TestNotificationValidateValidExpectNil(t *testing.T) {
	var notifications []*Notification
	for i := 0; i < tableRunCount; i++ {
		notifications = append(notifications, GenerateValidNotification())
	}
	for _, n := range notifications {
		name := fmt.Sprintf("%+v", *n)
		t.Run(name, func(t *testing.T) {
			err := n.Validate()
			if err != nil {
				t.Error("expect nil for valid notification")
			}
		})
	}
}
