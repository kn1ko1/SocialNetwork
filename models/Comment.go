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
	// Validate logic here
	if c.Body == "" {
		return errors.New("comment body must not be empty")
	}
	return nil
}

