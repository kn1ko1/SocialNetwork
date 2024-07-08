package models

import (
	"errors"
	"math/rand"
)

type Group struct {
	GroupId     int    `json:"groupId"`
	CreatedAt   int64  `json:"createdAt"`
	CreatorId   int    `json:"creatorId"`
	Description string `json:"description"`
	Title       string `json:"title"`
	UpdatedAt   int64  `json:"updatedAt"`
}

func (g *Group) Validate() error {
	// Validate logic here
	if g.CreatedAt <= 0 {
		return errors.New("invalid 'CreatedAt' field")
	}
	if g.CreatorId <= 0 {
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
		CreatorId:   rand.Intn(1000) + 1,
		Description: sutDescriptions[idxDescriptions],
		Title:       sutTitle[idxTitle],
		UpdatedAt:   ctime,
	}
	return g
}

func GenerateMissingFieldGroup() *Group {
	g := GenerateValidGroup()
	missingField := rand.Intn(5)
	switch missingField {
	case 0:
		g.CreatedAt = 0
	case 1:
		g.CreatorId = 0
	case 2:
		g.Description = ""
	case 3:
		g.Title = ""
	case 4:
		g.UpdatedAt = 0
	}
	return g
}

func GenerateInvalidGroup() *Group {
	g := GenerateValidGroup()
	invalidField := rand.Intn(5)
	switch invalidField {
	case 0:
		g.CreatedAt = -g.CreatedAt
	case 1:
		g.CreatorId = -g.CreatorId
	case 2:
		g.Description = ""
	case 3:
		g.Title = ""
	case 4:
		g.UpdatedAt = -g.UpdatedAt
	}
	return g
}
