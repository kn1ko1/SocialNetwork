package models

import (
	"errors"
	"math/rand"
)

type GroupUser struct {
	GroupUserId int
	CreatedAt   int64
	GroupId     int
	UpdatedAt   int64
	UserId      int
}

func (gu *GroupUser) Validate() error {
	if gu.CreatedAt <= 0 {
		return errors.New("invalid 'CreatedAt' field")
	}
	if gu.GroupId <= 0 {
		return errors.New("invalid 'GroupId' field")
	}
	if gu.UpdatedAt < gu.CreatedAt {
		return errors.New("invalid 'UpdatedAt' field. cannot be before 'CreatedAt' field")
	}
	if gu.UserId <= 0 {
		return errors.New("invalid 'UserId' field")
	}
	return nil
}

func GenerateValidGroupUser() *GroupUser {
	ctime := rand.Int63n(1000) + 1
	gu := &GroupUser{
		CreatedAt: ctime,
		GroupId:   rand.Intn(1000) + 1,
		UpdatedAt: ctime,
		UserId:    rand.Intn(1000) + 1,
	}
	return gu
}
