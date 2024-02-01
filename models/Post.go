package models

import (
	"errors"
	"math/rand"
)

type Post struct {
	PostId    int
	Body      string
	CreatedAt int64
	GroupId   int
	ImageURL  string
	Privacy   string
	UpdatedAt int64
	UserId    int
}

func (p *Post) Validate() error {
	// Validate logic here
	if p.Body == "" {
		return errors.New("invalid 'Body' field")
	}
	if p.CreatedAt <= 0 {
		return errors.New("invalid 'CreatedAt' field")
	}
	// GroupID can be 0 - i.e. not posted to a Group - but cannot be negative
	if p.GroupId < 0 {
		return errors.New("invalid 'GroupId' field")
	}
	if p.UpdatedAt < p.CreatedAt {
		return errors.New("invalid 'UpdatedAt' field")
	}
	if p.UserId <= 0 {
		return errors.New("invalid 'UserId' field")
	}
	return nil
}

func GenerateValidPost() *Post {
	ctime := rand.Int63n(1000) + 1
	idxBody := rand.Intn(len(sutBody))
	idxImageURL := rand.Intn(len(sutImageURL))
	idxPrivacy := rand.Intn(len(sutPrivacy))

	p := &Post{
		Body:      sutBody[idxBody],
		CreatedAt: ctime,
		GroupId:   rand.Intn(1000) + 1,
		ImageURL:  sutImageURL[idxImageURL],
		Privacy:   sutPrivacy[idxPrivacy],
		UpdatedAt: ctime,
		UserId:    rand.Intn(1000) + 1,
	}
	return p
}
