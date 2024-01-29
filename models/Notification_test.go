package models

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestNotificationValidateInvalidFieldExpectError(t *testing.T) {
	var notifications []*Notification
	for i := 0; i < tableRunCount; i++ {
		notifications = append(notifications, generateInvalidNotification())
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
		notifications = append(notifications, generateMissingFieldNotification())
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

func generateMissingFieldNotification() *Notification {
	n := GenerateValidNotification()
	missingField := rand.Intn(6)
	switch missingField {
	case 0:
		n.CreatedAt = 0
	case 1:
		n.NotificationType = ""
	case 2:
		n.SenderId = 0
	case 3:
		n.Status = ""
	case 4:
		n.TargetId = 0
	case 5:
		n.UpdatedAt = 0
	}
	return n
}

func generateInvalidNotification() *Notification {
	n := GenerateValidNotification()
	invalidField := rand.Intn(6)
	switch invalidField {
	case 0:
		n.CreatedAt = -n.CreatedAt
	case 1:
		n.NotificationType = ""
	case 2:
		n.SenderId = -n.SenderId
	case 3:
		n.Status = ""
	case 4:
		n.TargetId = -n.TargetId
	case 5:
		n.UpdatedAt = -n.UpdatedAt
	}
	return n
}
