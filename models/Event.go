package models

import (
	"errors"
	"math/rand"
)

type Event struct {
	EventId     int
	CreatedAt   int64
	DateTime    int64
	Description string
	GroupId     int
	Title       string
	UpdatedAt   int64
	UserId      int
}

func (e *Event) Validate() error {
	// Validate logic here
	if e.CreatedAt <= 0 {
		return errors.New("invalid 'CreatedAt' field")
	}
	if e.DateTime <= 0 {
		return errors.New("invalid 'DateTime' field")
	}
	if e.Description == "" {
		return errors.New("invalid 'Description' field")
	}
	if e.GroupId <= 0 {
		return errors.New("invalid 'GroupId' field")
	}
	if e.Title == "" {
		return errors.New("invalid 'Title' field")
	}
	if e.UpdatedAt <= 0 {
		return errors.New("invalid 'UpdatedAt' field")
	}
	if e.UserId <= 0 {
		return errors.New("invalid 'UserId' field")
	}
	return nil
}

func GenerateValidEvent() *Event {
	ctime := rand.Int63n(1000) + 1
	idxDesc := rand.Intn(len(sutDescriptions))
	idxTitle := rand.Intn(len(sutTitles))
	e := &Event{
		CreatedAt:   ctime,
		DateTime:    rand.Int63n(1000) + 1,
		Description: sutDescriptions[idxDesc],
		GroupId:     rand.Intn(1000) + 1,
		Title:       sutTitles[idxTitle],
		UpdatedAt:   ctime,
		UserId:      rand.Intn(1000) + 1,
	}
	return e
}
