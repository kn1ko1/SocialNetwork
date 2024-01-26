package models

import (
	"errors"
	"math/rand"
)

type Group struct {
	GroupID     int
	CreatedAt   int64
	CreatorID   int
	Description string
	Title       string
	UpdatedAt   int64
}

func (g *Group) Validate() error {
	// Validate logic here
	if g.CreatedAt <= 0 {
		return errors.New("invalid 'CreatedAt' field")
	}
	if g.CreatorID <= 0 {
		return errors.New("invalid 'CreatorID' field")
	}
	if g.Description == "" {
		return errors.New("invalid 'Description' field")
	}
	if g.Title == "" {
		return errors.New("invalid 'Title' field")
	}
	if g.UpdatedAt <= 0 {
		return errors.New("invalid 'UpdatedAt' field")
	}
	return nil
}

func GenerateValidGroup() *Group {
	ctime := rand.Int63n(1000) + 1
	idxDescriptions := rand.Intn(len(sutDescriptions))
	idxTitle := rand.Intn(len(sutTitle))
	g := &Group{
		CreatedAt:   ctime,
		CreatorID:   rand.Intn(1000) + 1,
		Description: sutDescriptions[idxDescriptions],
		Title:       sutTitle[idxTitle],
		UpdatedAt:   ctime,
	}
	return g
}
