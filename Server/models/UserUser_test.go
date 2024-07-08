package models

import (
	"fmt"
	"testing"
)

func TestUserUserValidateInvalidFieldExpectError(t *testing.T) {
	var userUsers []*UserUser
	for i := 0; i < tableRunCount; i++ {
		userUsers = append(userUsers, GenerateInvalidUserUser())
	}
	for _, uu := range userUsers {
		name := fmt.Sprintf("%+v", *uu)
		t.Run(name, func(t *testing.T) {
			err := uu.Validate()
			if err == nil {
				t.Error("expect error for invalid user user field")
			}
		})
	}
}

func TestUserUserValidateMissingFieldExpectError(t *testing.T) {
	var userUsers []*UserUser
	for i := 0; i < tableRunCount; i++ {
		userUsers = append(userUsers, GenerateMissingFieldUserUser())
	}
	for _, uu := range userUsers {
		name := fmt.Sprintf("%+v", *uu)
		t.Run(name, func(t *testing.T) {
			err := uu.Validate()
			if err == nil {
				t.Error("expect error for missing user user field")
			}
		})
	}
}

func TestUserUserValidateValidExpectNil(t *testing.T) {
	var userUsers []*UserUser
	for i := 0; i < tableRunCount; i++ {
		userUsers = append(userUsers, GenerateValidUserUser())
	}
	for _, uu := range userUsers {
		name := fmt.Sprintf("%+v", *uu)
		t.Run(name, func(t *testing.T) {
			err := uu.Validate()
			if err != nil {
				t.Error("expect nil for valid user user")
			}
		})
	}
}
