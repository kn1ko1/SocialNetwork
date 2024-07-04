package models

import (
	"fmt"
	"testing"
)

func TestGroupUserValidateInvalidFieldExpectError(t *testing.T) {
	var groupUsers []*GroupUser
	for i := 0; i < tableRunCount; i++ {
		groupUsers = append(groupUsers, GenerateInvalidGroupUser())
	}
	for _, gu := range groupUsers {
		name := fmt.Sprintf("%+v", *gu)
		t.Run(name, func(t *testing.T) {
			err := gu.Validate()
			if err == nil {
				t.Error("expect error for group user field")
			}
		})
	}
}

func TestGroupUserValidateMissingFieldExpectError(t *testing.T) {
	var groupUsers []*GroupUser
	for i := 0; i < tableRunCount; i++ {
		groupUsers = append(groupUsers, GenerateMissingFieldGroupUser())
	}
	for _, gu := range groupUsers {
		name := fmt.Sprintf("%+v", *gu)
		t.Run(name, func(t *testing.T) {
			err := gu.Validate()
			if err == nil {
				t.Error("expect error for missing group user field")
			}
		})
	}
}

func TestGroupUserValidateValidExpectNil(t *testing.T) {
	var groupUsers []*GroupUser
	for i := 0; i < tableRunCount; i++ {
		groupUsers = append(groupUsers, GenerateValidGroupUser())
	}
	for _, gu := range groupUsers {
		name := fmt.Sprintf("%+v", *gu)
		t.Run(name, func(t *testing.T) {
			err := gu.Validate()
			if err != nil {
				t.Error("expect nil for valid group user")
			}
		})
	}
}
