package models

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestPostUserValidateInvalidFieldExpectError(t *testing.T) {
	var postUsers []*PostUser
	for i := 0; i < tableRunCount; i++ {
		postUsers = append(postUsers, generateInvalidPostUser())
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
		postUsers = append(postUsers, generateMissingFieldPostUser())
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

func generateMissingFieldPostUser() *PostUser {
	pu := GenerateValidPostUser()
	missingField := rand.Intn(4)
	switch missingField {
	case 0:
		pu.CreatedAt = 0
	case 1:
		pu.PostId = 0
	case 2:
		pu.UpdatedAt = 0
	case 3:
		pu.UserId = 0
	}
	return pu
}

func generateInvalidPostUser() *PostUser {
	pu := GenerateValidPostUser()
	invalidField := rand.Intn(4)
	switch invalidField {
	case 0:
		pu.CreatedAt = -pu.CreatedAt
	case 1:
		pu.PostId = -pu.PostId
	case 2:
		pu.UpdatedAt = -pu.UpdatedAt
	case 3:
		pu.UserId = -pu.UserId
	}
	return pu
}
