package models

import (
	"errors"
	"math/rand"
)

type EventUser struct {
	EventUserId int
	CreatedAt   int64
	EventId     int
	UpdatedAt   int64
	UserId      int
}

func (eu *EventUser) Validate() error {
	if eu.CreatedAt <= 0 {
		return errors.New("invalid 'CreatedAt' field")
	}
	if eu.EventId <= 0 {
		return errors.New("invalid 'EventId' field")
	}
	if eu.UpdatedAt < eu.CreatedAt {
		return errors.New("invalid 'UpdatedAt' field. cannot be before 'CreatedAt' field")
	}
	if eu.UserId <= 0 {
		return errors.New("invalid 'UserId' field")
	}
	return nil
}

func GenerateValidEventUser() *EventUser {
	ctime := rand.Int63n(1000) + 1
	eu := &EventUser{
		CreatedAt: ctime,
		EventId:   rand.Intn(1000) + 1,
		UpdatedAt: ctime,
		UserId:    rand.Intn(1000) + 1,
	}
	return eu
}
