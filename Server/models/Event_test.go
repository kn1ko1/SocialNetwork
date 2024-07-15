package models

import (
	"fmt"
	"testing"
)

func TestEventValidateInvalidFieldExpectError(t *testing.T) {
	var events []*Event
	for i := 0; i < tableRunCount; i++ {
		events = append(events, GenerateInvalidEvent())
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
		events = append(events, GenerateMissingFieldEvent())
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
