package models

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestEventValidateInvalidFieldExpectError(t *testing.T) {
	var events []*Event
	for i := 0; i < tableRunCount; i++ {
		events = append(events, generateInvalidEvent())
	}
	for _, e := range events {
		name := fmt.Sprintf("%+v", *e)
		t.Run(name, func(t *testing.T) {
			err := e.Validate()
			if err == nil {
				t.Error("expect error for invalid event field")
			}
		})
	}
}

func TestEventValidateMissingFieldExpectError(t *testing.T) {
	var events []*Event
	for i := 0; i < tableRunCount; i++ {
		events = append(events, generateMissingFieldEvent())
	}
	for _, e := range events {
		name := fmt.Sprintf("%+v", *e)
		t.Run(name, func(t *testing.T) {
			err := e.Validate()
			if err == nil {
				t.Error("expect error for missing event field")
			}
		})
	}
}

func TestEventValidateValidExpectNil(t *testing.T) {
	var events []*Event
	for i := 0; i < tableRunCount; i++ {
		events = append(events, GenerateValidEvent())
	}
	for _, e := range events {
		name := fmt.Sprintf("%+v", *e)
		t.Run(name, func(t *testing.T) {
			err := e.Validate()
			if err != nil {
				t.Error("expect nil for valid event")
			}
		})
	}
}

func generateMissingFieldEvent() *Event {
	e := GenerateValidEvent()
	missingField := rand.Intn(7)
	switch missingField {
	case 0:
		e.CreatedAt = 0
	case 1:
		e.DateTime = 0
	case 2:
		e.Description = ""
	case 3:
		e.GroupId = 0
	case 4:
		e.Title = ""
	case 5:
		e.UpdatedAt = 0
	case 6:
		e.UserId = 0
	}
	return e
}

func generateInvalidEvent() *Event {
	e := GenerateValidEvent()
	invalidField := rand.Intn(7)
	switch invalidField {
	case 0:
		e.CreatedAt = -e.CreatedAt
	case 1:
		e.DateTime = -e.DateTime
	case 2:
		e.Description = ""
	case 3:
		e.GroupId = -e.GroupId
	case 4:
		e.Title = ""
	case 5:
		e.UpdatedAt = -e.UpdatedAt
	case 6:
		e.UserId = -e.UserId
	}
	return e
}
