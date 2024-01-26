package models

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestUserValidateInvalidFieldExpectError(t *testing.T) {
	var users []*User
	for i := 0; i < tableRunCount; i++ {
		users = append(users, generateInvalidUser())
	}
	for _, u := range users {
		name := fmt.Sprintf("%+v", *u)
		t.Run(name, func(t *testing.T) {
			err := u.Validate()
			if err == nil {
				t.Error("expect error for invalid event field")
			}
		})
	}
}

func TestUserValidateMissingFieldExpectError(t *testing.T) {
	var users []*User
	for i := 0; i < tableRunCount; i++ {
		users = append(users, generateMissingFieldUser())
	}
	for _, u := range users {
		name := fmt.Sprintf("%+v", *u)
		t.Run(name, func(t *testing.T) {
			err := u.Validate()
			if err == nil {
				t.Error("expect error for missing event field")
			}
		})
	}
}

func TestUserValidateValidExpectNil(t *testing.T) {
	var users []*User
	for i := 0; i < tableRunCount; i++ {
		users = append(users, GenerateValidUser(false))
	}
	for _, u := range users {
		name := fmt.Sprintf("%+v", *u)
		t.Run(name, func(t *testing.T) {
			err := u.Validate()
			if err != nil {
				t.Error("expect nil for valid event")
			}
		})
	}
}

func generateMissingFieldUser() *User {
	u := GenerateValidUser(false)
	missingField := rand.Intn(7)
	switch missingField {
	case 0:
		u.CreatedAt = 0
	case 1:
		u.Email = ""
	case 2:
		u.EncryptedPassword = ""
	case 3:
		u.FirstName = ""
	case 4:
		u.LastName = ""
	case 5:
		u.UpdatedAt = 0
	case 6:
		u.Username = ""
	}
	return u
}

func generateInvalidUser() *User {
	u := GenerateValidUser(false)
	invalidField := rand.Intn(8)
	switch invalidField {
	case 0:
		u.CreatedAt = -u.CreatedAt
	case 1:
		u.DOB = -631238400
	case 2:
		u.Email = ""
	case 3:
		u.EncryptedPassword = ""
	case 4:
		u.FirstName = ""
	case 5:
		u.LastName = ""
	case 6:
		u.UpdatedAt = -u.UpdatedAt
	case 7:
		u.Username = ""
	}
	return u
}
