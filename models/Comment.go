package models

import "errors"

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
