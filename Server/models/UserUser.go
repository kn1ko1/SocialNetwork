package models

import (
	"errors"
	"math/rand"
)

type UserUser struct {
	UserUserId int   `json:"userUserId"`
	CreatedAt  int64 `json:"createdAt"`
	FollowerId int   `json:"followerId"`
	SubjectId  int   `json:"subjectId"`
	UpdatedAt  int64 `json:"updatedAt"`
}

func (uu *UserUser) Validate() error {
	if uu.CreatedAt <= 0 {
		return errors.New("invalid 'CreatedAt' field")
	}
	if uu.FollowerId <= 0 {
		return errors.New("invalid 'FollowerId' field")
	}
	if uu.SubjectId <= 0 {
		return errors.New("invalid 'SubjectId' field")
	}
	if uu.UpdatedAt < uu.CreatedAt {
		return errors.New("invalid 'UpdatedAt' field. cannot be before 'CreatedAt' field")
	}
	return nil
}

func GenerateValidUserUser() *UserUser {
	ctime := rand.Int63n(1000) + 1
	uu := &UserUser{
		CreatedAt:  ctime,
		FollowerId: rand.Intn(1000) + 1,
		SubjectId:  rand.Intn(1000) + 1,
		UpdatedAt:  ctime,
	}
	return uu
}

func GenerateMissingFieldUserUser() *UserUser {
	uu := GenerateValidUserUser()
	missingField := rand.Intn(4)
	switch missingField {
	case 0:
		uu.CreatedAt = 0
	case 1:
		uu.FollowerId = 0
	case 2:
		uu.SubjectId = 0
	case 3:
		uu.UpdatedAt = 0
	}
	return uu
}

func GenerateInvalidUserUser() *UserUser {
	uu := GenerateValidUserUser()
	invalidField := rand.Intn(4)
	switch invalidField {
	case 0:
		uu.CreatedAt = -uu.CreatedAt
	case 1:
		uu.FollowerId = -uu.FollowerId
	case 2:
		uu.SubjectId = -uu.SubjectId
	case 3:
		uu.UpdatedAt = -uu.UpdatedAt
	}
	return uu
}
