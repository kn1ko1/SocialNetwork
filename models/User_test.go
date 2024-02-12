package models

import (
	"fmt"
	"testing"
)

func TestUserValidateInvalidFieldExpectError(t *testing.T) {
	var users []*User
	for i := 0; i < tableRunCount; i++ {
		users = append(users, GenerateInvalidUser())
	}
	for _, u := range users {
		name := fmt.Sprintf("%+v", *u)
		t.Run(name, func(t *testing.T) {
			err := u.Validate()
			if err == nil {
				t.Error("expect error for invalid user field")
			}
		})
	}
}

func TestUserValidateMissingFieldExpectError(t *testing.T) {
	var users []*User
	for i := 0; i < tableRunCount; i++ {
		users = append(users, GenerateMissingFieldUser())
	}
	for _, u := range users {
		name := fmt.Sprintf("%+v", *u)
		t.Run(name, func(t *testing.T) {
			err := u.Validate()
			if err == nil {
				t.Error("expect error for missing user field")
			}
		})
	}
}

func TestUserValidateValidExpectNil(t *testing.T) {
	var users []*User
	for i := 0; i < tableRunCount; i++ {
		users = append(users, GenerateValidUser())
	}
	for _, u := range users {
		name := fmt.Sprintf("%+v", *u)
		t.Run(name, func(t *testing.T) {
			err := u.Validate()
			if err != nil {
				t.Error("expect nil for valid user")
			}
		})
	}
}
