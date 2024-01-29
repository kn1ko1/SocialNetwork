package models

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestEventUserValidateInvalidFieldExpectError(t *testing.T) {
	var eventUsers []*EventUser
	for i := 0; i < tableRunCount; i++ {
		eventUsers = append(eventUsers, generateInvalidEventUser())
	}
	for _, eu := range eventUsers {
		name := fmt.Sprintf("%+v", *eu)
		t.Run(name, func(t *testing.T) {
			err := eu.Validate()
			if err == nil {
				t.Error("expect error for event user field")
			}
		})
	}
}

func TestEventUserValidateMissingFieldExpectError(t *testing.T) {
	var eventUsers []*EventUser
	for i := 0; i < tableRunCount; i++ {
		eventUsers = append(eventUsers, generateMissingFieldEventUser())
	}
	for _, eu := range eventUsers {
		name := fmt.Sprintf("%+v", *eu)
		t.Run(name, func(t *testing.T) {
			err := eu.Validate()
			if err == nil {
				t.Error("expect error for missing event user field")
			}
		})
	}
}

func TestEventUserValidateValidExpectNil(t *testing.T) {
	var eventUsers []*EventUser
	for i := 0; i < tableRunCount; i++ {
		eventUsers = append(eventUsers, GenerateValidEventUser())
	}
	for _, eu := range eventUsers {
		name := fmt.Sprintf("%+v", *eu)
		t.Run(name, func(t *testing.T) {
			err := eu.Validate()
			if err != nil {
				t.Error("expect nil for valid event user")
			}
		})
	}
}

func generateMissingFieldEventUser() *EventUser {
	eu := GenerateValidEventUser()
	missingField := rand.Intn(4)
	switch missingField {
	case 0:
		eu.CreatedAt = 0
	case 1:
		eu.EventId = 0
	case 2:
		eu.UpdatedAt = 0
	case 3:
		eu.UserId = 0
	}
	return eu
}

func generateInvalidEventUser() *EventUser {
	eu := GenerateValidEventUser()
	invalidField := rand.Intn(4)
	switch invalidField {
	case 0:
		eu.CreatedAt = -eu.CreatedAt
	case 1:
		eu.EventId = -eu.EventId
	case 2:
		eu.UpdatedAt = -eu.UpdatedAt
	case 3:
		eu.UserId = -eu.UserId
	}
	return eu
}
