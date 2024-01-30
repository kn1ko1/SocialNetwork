package models

import (
	"errors"
	"math/rand"
)

type Comment struct {
	CommentId int
	Body      string
	CreatedAt int64
	ImageURL  string
	PostId    int
	UpdatedAt int64
	UserId    int
}

func (c *Comment) Validate() error {
	if c.Body == "" {
		return errors.New("comment body must not be empty")
	}
	if c.CreatedAt <= 0 {
		return errors.New("invalid 'CreatedAt' field")
	}
	if c.PostId <= 0 {
		return errors.New("invalid 'PostId' field")
	}
	if c.UpdatedAt < c.CreatedAt {
		return errors.New("invalid 'UpdatedAt' field. cannot be before 'CreatedAt' field")
	}
	if c.UserId <= 0 {
		return errors.New("invalid 'UserId' field")
	}

	return nil
}

func GenerateValidComment() *Comment {
	ctime := rand.Int63n(1000) + 1
	idx := rand.Intn(len(sutBody))
	c := &Comment{
		Body:      sutBody[idx],
		CreatedAt: ctime,
		PostId:    rand.Intn(1000) + 1,
		UpdatedAt: ctime,
		UserId:    rand.Intn(1000) + 1,
	}
	return c
}
