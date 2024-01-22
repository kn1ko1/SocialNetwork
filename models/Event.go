package models

import "errors"

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
