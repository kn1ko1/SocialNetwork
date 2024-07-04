package models

import (
	"errors"
	"math/rand"
)

type EventUser struct {
	EventUserId int   `json:"eventUserId"`
	CreatedAt   int64 `json:"createdAt"`
	EventId     int   `json:"eventId"`
	IsGoing     bool  `json:"isGoing"`
	UpdatedAt   int64 `json:"updatedAt"`
	UserId      int   `json:"userId"`
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
	idxIsGoing := rand.Intn(len(sutIsGoing))
	eu := &EventUser{
		CreatedAt: ctime,
		EventId:   rand.Intn(1000) + 1,
		IsGoing:   sutIsGoing[idxIsGoing],
		UpdatedAt: ctime,
		UserId:    rand.Intn(1000) + 1,
	}
	return eu
}

func GenerateMissingFieldEventUser() *EventUser {
	eu := GenerateValidEventUser()
	missingField := rand.Intn(4)
	switch missingField {
	case 0:
		eu.CreatedAt = 0
	case 1:
		eu.EventId = 0
	case 2:
		eu.UpdatedAt = 0
	case 3:
		eu.UserId = 0
	}
	return eu
}

func GenerateInvalidEventUser() *EventUser {
	eu := GenerateValidEventUser()
	invalidField := rand.Intn(4)
	switch invalidField {
	case 0:
		eu.CreatedAt = -eu.CreatedAt
	case 1:
		eu.EventId = -eu.EventId
	case 2:
		eu.UpdatedAt = -eu.UpdatedAt
	case 3:
		eu.UserId = -eu.UserId
	}
	return eu
}
