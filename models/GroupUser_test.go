package models

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestGroupUserValidateInvalidFieldExpectError(t *testing.T) {
	var groupUsers []*GroupUser
	for i := 0; i < tableRunCount; i++ {
		groupUsers = append(groupUsers, generateInvalidGroupUser())
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
		groupUsers = append(groupUsers, generateMissingFieldGroupUser())
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

func generateMissingFieldGroupUser() *GroupUser {
	gu := GenerateValidGroupUser()
	missingField := rand.Intn(4)
	switch missingField {
	case 0:
		gu.CreatedAt = 0
	case 1:
		gu.GroupId = 0
	case 2:
		gu.UpdatedAt = 0
	case 3:
		gu.UserId = 0
	}
	return gu
}

func generateInvalidGroupUser() *GroupUser {
	gu := GenerateValidGroupUser()
	invalidField := rand.Intn(4)
	switch invalidField {
	case 0:
		gu.CreatedAt = -gu.CreatedAt
	case 1:
		gu.GroupId = -gu.GroupId
	case 2:
		gu.UpdatedAt = -gu.UpdatedAt
	case 3:
		gu.UserId = -gu.UserId
	}
	return gu
}
