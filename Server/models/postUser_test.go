package models

import (
	"fmt"
	"testing"
)

func TestPostUserValidateInvalidFieldExpectError(t *testing.T) {
	var postUsers []*PostUser
	for i := 0; i < tableRunCount; i++ {
		postUsers = append(postUsers, GenerateInvalidPostUser())
	}
	for _, pu := range postUsers {
		name := fmt.Sprintf("%+v", *pu)
		t.Run(name, func(t *testing.T) {
			err := pu.Validate()
			if err == nil {
				t.Error("expect error for post user field")
			}
		})
	}
}

func TestPostUserValidateMissingFieldExpectError(t *testing.T) {
	var postUsers []*PostUser
	for i := 0; i < tableRunCount; i++ {
		postUsers = append(postUsers, GenerateMissingFieldPostUser())
	}
	for _, pu := range postUsers {
		name := fmt.Sprintf("%+v", *pu)
		t.Run(name, func(t *testing.T) {
			err := pu.Validate()
			if err == nil {
				t.Error("expect error for missing post user field")
			}
		})
	}
}

func TestPostUserValidateValidExpectNil(t *testing.T) {
	var postUsers []*PostUser
	for i := 0; i < tableRunCount; i++ {
		postUsers = append(postUsers, GenerateValidPostUser())
	}
	for _, pu := range postUsers {
		name := fmt.Sprintf("%+v", *pu)
		t.Run(name, func(t *testing.T) {
			err := pu.Validate()
			if err != nil {
				t.Error("expect nil for valid post user")
			}
		})
	}
}
