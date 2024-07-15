package models

import (
	"fmt"
	"testing"
)

func TestEventUserValidateInvalidFieldExpectError(t *testing.T) {
	var eventUsers []*EventUser
	for i := 0; i < tableRunCount; i++ {
		eventUsers = append(eventUsers, GenerateInvalidEventUser())
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
		eventUsers = append(eventUsers, GenerateMissingFieldEventUser())
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
