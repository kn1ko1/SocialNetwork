package models

import (
	"fmt"
	"testing"
)

func TestGroupValidateInvalidFieldExpectError(t *testing.T) {
	var groups []*Group
	for i := 0; i < tableRunCount; i++ {
		groups = append(groups, GenerateInvalidGroup())
	}
	for _, g := range groups {
		name := fmt.Sprintf("%+v", *g)
		t.Run(name, func(t *testing.T) {
			err := g.Validate()
			if err == nil {
				t.Error("expect error for invalid group field")
			}
		})
	}
}

func TestGroupValidateMissingFieldExpectError(t *testing.T) {
	var groups []*Group
	for i := 0; i < tableRunCount; i++ {
		groups = append(groups, GenerateMissingFieldGroup())
	}
	for _, g := range groups {
		name := fmt.Sprintf("%+v", *g)
		t.Run(name, func(t *testing.T) {
			err := g.Validate()
			if err == nil {
				t.Error("expect error for missing group field")
			}
		})
	}
}

func TestGroupValidateValidExpectNil(t *testing.T) {
	var groups []*Group
	for i := 0; i < tableRunCount; i++ {
		groups = append(groups, GenerateValidGroup())
	}
	for _, g := range groups {
		name := fmt.Sprintf("%+v", *g)
		t.Run(name, func(t *testing.T) {
			err := g.Validate()
			if err != nil {
				t.Error("expect nil for valid group")
			}
		})
	}
}
