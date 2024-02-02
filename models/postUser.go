package models

import (
	"errors"
	"math/rand"
)

type PostUser struct {
	PostUserId int
	CreatedAt  int64
	PostId     int
	UpdatedAt  int64
	UserId     int
}

func (pu *PostUser) Validate() error {
	if pu.CreatedAt <= 0 {
		return errors.New("invalid 'CreatedAt' field")
	}
	if pu.PostId <= 0 {
		return errors.New("invalid 'PostId' field")
	}
	if pu.UpdatedAt < pu.CreatedAt {
		return errors.New("invalid 'UpdatedAt' field. cannot be before 'CreatedAt' field")
	}
	if pu.UserId <= 0 {
		return errors.New("invalid 'UserId' field")
	}
	return nil
}

func GenerateValidPostUser() *PostUser {
	ctime := rand.Int63n(1000) + 1
	pu := &PostUser{
		CreatedAt: ctime,
		PostId:    rand.Intn(1000) + 1,
		UpdatedAt: ctime,
		UserId:    rand.Intn(1000) + 1,
	}
	return pu
}
